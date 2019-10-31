// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package main_test

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testidentity"
	"storj.io/storj/internal/version"
	"storj.io/storj/pkg/identity"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/versioncontrol"
)

const (
	oldVersion = "v0.19.0"
	newVersion = "v0.19.5"
)

func TestAutoUpdater_unix(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("requires storagenode and storagenode-updater to be installed as windows services")
	}

	// TODO cleanup `.exe` extension for different OS

	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	oldSemVer, err := version.NewSemVer(oldVersion)
	require.NoError(t, err)

	newSemVer, err := version.NewSemVer(newVersion)
	require.NoError(t, err)

	oldInfo := version.Info{
		Timestamp:  time.Now(),
		CommitHash: "",
		Version:    oldSemVer,
		Release:    false,
	}

	// build real storagenode and updater with old version
	oldStoragenodeBin := ctx.CompileWithVersion("storj.io/storj/cmd/storagenode/", oldInfo)
	storagenodePath := ctx.File("fake", "storagenode.exe")
	move(t, storagenodePath, oldStoragenodeBin)

	oldUpdaterBin := ctx.CompileWithVersion("", oldInfo)
	updaterPath := ctx.File("fake", "storagenode-updater.exe")
	move(t, updaterPath, oldUpdaterBin)

	// build real storagenode and updater with new version
	newInfo := version.Info{
		Timestamp:  time.Now(),
		CommitHash: "",
		Version:    newSemVer,
		Release:    false,
	}
	newStoragenodeBin := ctx.CompileWithVersion("storj.io/storj/cmd/storagenode/", newInfo)
	newUpdaterBin := ctx.CompileWithVersion("", newInfo)
	updateBins := map[string]string{
		"storagenode":         newStoragenodeBin,
		"storagenode-updater": newUpdaterBin,
	}

	// run versioncontrol and update zips http servers
	versionControlPeer, cleanupVersionControl := testVersionControlWithUpdates(ctx, t, updateBins)
	defer cleanupVersionControl()

	logPath := ctx.File("storagenode-updater.log")

	// write identity files to disk for use in rollout calculation
	identConfig := testIdentityFiles(ctx, t)

	// run updater (update)
	args := []string{"run"}
	args = append(args, "--config-dir", ctx.Dir())
	args = append(args, "--server-address", "http://"+versionControlPeer.Addr())
	args = append(args, "--binary-location", storagenodePath)
	args = append(args, "--check-interval", "0s")
	args = append(args, "--identity.cert-path", identConfig.CertPath)
	args = append(args, "--identity.key-path", identConfig.KeyPath)
	args = append(args, "--log", logPath)

	// NB: updater currently uses `log.SetOutput` so all output after that call
	// only goes to the log file.
	out, err := exec.Command(updaterPath, args...).CombinedOutput()
	logData, logErr := ioutil.ReadFile(logPath)
	if assert.NoError(t, logErr) {
		logStr := string(logData)
		t.Log(logStr)
		if !assert.Contains(t, logStr, "storagenode restarted successfully") {
			t.Log(logStr)
		}
		if !assert.Contains(t, logStr, "storagenode-updater restarted successfully") {
			t.Log(logStr)
		}
	} else {
		t.Log(string(out))
	}
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	oldStoragenode := ctx.File("fake", "storagenode"+".old."+oldVersion+".exe")
	oldStoragenodeInfo, err := os.Stat(oldStoragenode)
	require.NoError(t, err)
	require.NotNil(t, oldStoragenodeInfo)
	require.NotZero(t, oldStoragenodeInfo.Size())

	backupUpdater := ctx.File("fake", "storagenode-updater.old.exe")
	backupUpdaterInfo, err := os.Stat(backupUpdater)
	require.NoError(t, err)
	require.NotNil(t, backupUpdaterInfo)
	require.NotZero(t, backupUpdaterInfo.Size())
}

func move(t *testing.T, dst, src string) {
	err := os.Rename(src, dst)
	require.NoError(t, err)
}

func testIdentityFiles(ctx *testcontext.Context, t *testing.T) identity.Config {
	t.Helper()

	ident, err := testidentity.PregeneratedIdentity(0, storj.LatestIDVersion())
	require.NoError(t, err)

	identConfig := identity.Config{
		CertPath: ctx.File("identity", "identity.cert"),
		KeyPath:  ctx.File("identity", "identity.Key"),
	}
	err = identConfig.Save(ident)
	require.NoError(t, err)

	configData := fmt.Sprintf(
		"identity.cert-path: %s\nidentity.key-path: %s",
		identConfig.CertPath,
		identConfig.KeyPath,
	)
	err = ioutil.WriteFile(ctx.File("config.yaml"), []byte(configData), 0644)
	require.NoError(t, err)

	return identConfig
}

func testVersionControlWithUpdates(ctx *testcontext.Context, t *testing.T, updateBins map[string]string) (peer *versioncontrol.Peer, cleanup func()) {
	t.Helper()

	var mux http.ServeMux
	for name, src := range updateBins {
		dst := ctx.File("updates", name+".zip")
		zipBin(ctx, t, dst, src)
		zipData, err := ioutil.ReadFile(dst)
		require.NoError(t, err)

		mux.HandleFunc("/"+name, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write(zipData)
			require.NoError(t, err)
		}))
	}

	ts := httptest.NewServer(&mux)

	config := &versioncontrol.Config{
		// TODO: add STORJ_VERSION_SERVER_ADDR property to Product.wxs for testing
		// TODO: set address back to `127.0.0.1:0`
		Address: "127.0.0.1:10000",
		// NB: this config field is required for versioncontrol to run.
		Versions: versioncontrol.OldVersionConfig{
			Satellite:   "v0.0.1",
			Storagenode: "v0.0.1",
			Uplink:      "v0.0.1",
			Gateway:     "v0.0.1",
			Identity:    "v0.0.1",
		},
		Binary: versioncontrol.ProcessesConfig{
			Storagenode: versioncontrol.ProcessConfig{
				Suggested: versioncontrol.VersionConfig{
					Version: newVersion,
					URL:     ts.URL + "/storagenode",
				},
				Rollout: versioncontrol.RolloutConfig{
					Seed:   "0000000000000000000000000000000000000000000000000000000000000001",
					Cursor: 100,
				},
			},
			StoragenodeUpdater: versioncontrol.ProcessConfig{
				Suggested: versioncontrol.VersionConfig{
					Version: newVersion,
					URL:     ts.URL + "/storagenode-updater",
				},
				Rollout: versioncontrol.RolloutConfig{
					Seed:   "0000000000000000000000000000000000000000000000000000000000000001",
					Cursor: 100,
				},
			},
		},
	}
	peer, err := versioncontrol.New(zaptest.NewLogger(t), config)
	require.NoError(t, err)
	ctx.Go(func() error {
		return peer.Run(ctx)
	})
	return peer, func() {
		ts.Close()
		ctx.Check(peer.Close)
	}
}

func zipBin(ctx *testcontext.Context, t *testing.T, dst, src string) {
	t.Helper()

	zipFile, err := os.Create(dst)
	require.NoError(t, err)

	base := filepath.Base(dst)
	base = base[:len(base)-len(".zip")]

	writer := zip.NewWriter(zipFile)
	contents, err := writer.Create(base)
	require.NoError(t, err)

	data, err := ioutil.ReadFile(src)
	require.NoError(t, err)

	_, err = contents.Write(data)
	require.NoError(t, err)

	ctx.Check(writer.Close)
}

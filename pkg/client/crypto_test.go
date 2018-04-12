// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package client

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const mockMnemonic = "uncle obtain april oxygen cover patient layer abuse off text royal normal"

func TestDecryptBucketName(t *testing.T) {
	for i, tt := range []struct {
		encryptedName string
		mnemonic      string
		decryptedName string
		errString     string
	}{
		{"", "", "", "Invalid mnemonic"},
		{"", mockMnemonic, "", "Invalid encrypted name"},
		{mockEncryptedBucketName, "", "", "Invalid mnemonic"},
		{mockEncryptedBucketName, mockMnemonic, mockDecryptedBucketName, ""},
		{mockEncryptedBucketNameDiffMnemonic, mockMnemonic, "", "cipher: message authentication failed"},
	} {
		decryptedName, err := decryptBucketName(tt.encryptedName, tt.mnemonic)
		errTag := fmt.Sprintf("Test case #%d", i)
		if tt.errString != "" {
			assert.EqualError(t, err, tt.errString, errTag)
			continue
		}
		if assert.NoError(t, err, errTag) {
			assert.Equal(t, tt.decryptedName, decryptedName, errTag)
		}
	}
}

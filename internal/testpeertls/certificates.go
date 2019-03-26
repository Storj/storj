// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package testpeertls

import (
	"crypto"
	"crypto/x509"
	"storj.io/storj/pkg/peertls"
	"storj.io/storj/pkg/peertls/extensions"
	"storj.io/storj/pkg/pkcrypto"
	"storj.io/storj/pkg/storj"
)

// NewCertChain creates a valid peertls certificate chain (and respective keys) of the desired length.
// NB: keys are in the reverse order compared to certs (i.e. first key belongs to last cert)!
func NewCertChain(length int, versionNumber storj.IDVersionNumber) (keys []crypto.PrivateKey, certs []*x509.Certificate, _ error) {
	version, err := storj.GetIDVersion(versionNumber)
	if err != nil {
		return nil, nil, err
	}

	for i := 0; i < length; i++ {
		key, err := pkcrypto.GeneratePrivateKey()
		if err != nil {
			return nil, nil, err
		}
		keys = append(keys, key)

		var template *x509.Certificate
		if i == length-1 {
			template, err = peertls.CATemplate()
			if err = extensions.AddExtraExtension(template, storj.NewVersionExt(version)); err != nil {
				return nil, nil, err
			}
		} else {
			template, err = peertls.LeafTemplate()
		}
		if err != nil {
			return nil, nil, err
		}

		var cert *x509.Certificate
		if i == 0 {
			cert, err = peertls.CreateSelfSignedCertificate(key, template)
		} else {
			cert, err = peertls.CreateCertificate(pkcrypto.PublicKeyFromPrivate(key), keys[i-1], template, certs[i-1:][0])
		}
		if err != nil {
			return nil, nil, err
		}

		certs = append([]*x509.Certificate{cert}, certs...)
	}
	return keys, certs, nil
}

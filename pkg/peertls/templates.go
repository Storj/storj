// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package peertls

import (
	"crypto/x509"
)

// CATemplate returns x509.Certificate template for certificate authority
func CATemplate() (*x509.Certificate, error) {
	serialNumber, err := newSerialNumber()
	if err != nil {
		return nil, ErrTLSTemplate.Wrap(err)
	}

	template := &x509.Certificate{
		SerialNumber:          serialNumber,
		KeyUsage:              x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA: true,
	}

	return template, nil
}

// LeafTemplate returns x509.Certificate template for signing and encrypting
func LeafTemplate() (*x509.Certificate, error) {
	serialNumber, err := newSerialNumber()
	if err != nil {
		return nil, ErrTLSTemplate.Wrap(err)
	}

	template := &x509.Certificate{
		SerialNumber:          serialNumber,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IsCA: false,
	}

	return template, nil
}

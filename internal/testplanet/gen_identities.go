// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information

// +build ignore

// gen_identities generates random identities table for testing
package main

import (
	"bytes"
	"context"
	"crypto/x509"
	"flag"
	"fmt"
	"go/format"
	"os"

	"storj.io/storj/pkg/identity"
	"storj.io/storj/pkg/pkcrypto"
)

func main() {
	signed := flag.Bool("signed", false, "if true, generate a signer and sign all identities")
	count := flag.Int("count", 5, "number of identities to create")
	out := flag.String("out", "identities_table.go", "generated file")
	flag.Parse()

	var buf bytes.Buffer
	buf.WriteString(`
		// Copyright (C) 2019 Storj Labs, Inc.
		// See LICENSE for copying information
		
		// Code generated by gen_identities. DO NOT EDIT.

		package testplanet
		
		var (
	`)

	var (
		signer *identity.FullCertificateAuthority
		restChain []*x509.Certificate
		err error
	)
	if *signed {
		signer, err = identity.NewCA(context.Background(), identity.NewCAOptions{
			Difficulty:  12,
			Concurrency: 4,
		})
		if err != nil {
			panic(err)
		}
		restChain = []*x509.Certificate{signer.Cert}

		var chain bytes.Buffer
		err = pkcrypto.WriteCertPEM(&chain, signer.Cert)
		if err != nil {
			panic(err)
		}

		var keys bytes.Buffer
		err = pkcrypto.WritePrivateKeyPEM(&keys, signer.Key)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(&buf, "pregeneratedSigner = mustParseCertificateAuthorityPEM(%q, %q)\n", chain.Bytes(), keys.Bytes())
	}

	if *signed {
		buf.WriteString(`
			pregeneratedSignedIdentities = NewIdentities(
		`)
	} else {
		buf.WriteString(`
			pregeneratedIdentities = NewIdentities(
		`)
	}
	for k := 0; k < *count; k++ {
		fmt.Println("Creating", k)
		ca, err := identity.NewCA(context.Background(), identity.NewCAOptions{
			Difficulty:  12,
			Concurrency: 4,
		})
		if err != nil {
			panic(err)
		}

		if *signed {
			ca.Cert, err = signer.Sign(ca.Cert)
			if err != nil {
				panic(err)
			}
			ca.RestChain = restChain
		}

		ident, err := ca.NewIdentity()
		if err != nil {
			panic(err)
		}

		var chain bytes.Buffer
		certs := append([]*x509.Certificate{ident.Leaf, ca.Cert}, ca.RestChain...)
		err = pkcrypto.WriteCertPEM(&chain, certs...)
		if err != nil {
			panic(err)
		}

		var keys bytes.Buffer
		err = pkcrypto.WritePrivateKeyPEM(&keys, ident.Key)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(&buf, "mustParseIdentityPEM(%q, %q),\n", chain.Bytes(), keys.Bytes())
	}

	buf.WriteString(`))`)

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	file, err := os.Create(*out)
	if err != nil {
		panic(err)
	}

	if _, err := file.Write(formatted); err != nil {
		panic(err)
	}

	if err := file.Close(); err != nil {
		panic(err)
	}
}

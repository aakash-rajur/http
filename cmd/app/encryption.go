package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func generateRSAKeyPair(hash crypto.Hash, modulusLength int) (*RSAKeyPair, error) {
	key, err := rsa.GenerateKey(rand.Reader, modulusLength)

	if err != nil {
		return nil, err
	}

	publicBuffer, err := x509.MarshalPKIXPublicKey(key.Public())

	if err != nil {
		return nil, err
	}

	public := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicBuffer,
		},
	)

	decrypt := func(ciphertext []byte) ([]byte, error) {
		return key.Decrypt(nil, ciphertext, &rsa.OAEPOptions{Hash: hash})
	}

	keyPair := &RSAKeyPair{
		PrivateKey:     key,
		PublicKey:      key.Public().(*rsa.PublicKey),
		PublicKeyBytes: public,
		Decrypt:        decrypt,
	}

	return keyPair, nil
}

type RSAKeyPair struct {
	PrivateKey     *rsa.PrivateKey
	PublicKey      *rsa.PublicKey
	PublicKeyBytes []byte
	Decrypt        Decrypt
}

type Decrypt func(ciphertext []byte) ([]byte, error)

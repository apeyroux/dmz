package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"code.google.com/p/go.crypto/ssh"
)

func GenerateKey() (pkPem []byte, pubkPem []byte, pubSSHAK []byte, err error) {
	pk, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}
	pkDer := x509.MarshalPKCS1PrivateKey(pk)
	pkBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   pkDer,
	}
	pkPem = pem.EncodeToMemory(&pkBlock)

	pubk := pk.PublicKey
	pubkDer, err := x509.MarshalPKIXPublicKey(&pubk)
	if err != nil {
		return
	}

	pubkBlock := pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   pubkDer,
	}
	pubkPem = pem.EncodeToMemory(&pubkBlock)

	pubSSH, err := ssh.NewPublicKey(&pubk)
	if err != nil {
		return
	}
	pubSSHAK = ssh.MarshalAuthorizedKey(pubSSH)

	return
}

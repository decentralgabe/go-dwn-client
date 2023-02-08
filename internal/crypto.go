package internal

import (
	"crypto"

	ssicrypto "github.com/TBD54566975/ssi-sdk/crypto"
	"github.com/TBD54566975/ssi-sdk/did"
)

type KeyPair struct {
	ID         string            `json:"id"`
	KeyType    ssicrypto.KeyType `json:"keyType"`
	PrivateKey crypto.PrivateKey `json:"privateKey,omitempty"`
	PublicKey  crypto.PublicKey  `json:"publicKey"`
}

func GenerateDIDKeyPair(keyType ssicrypto.KeyType) (*KeyPair, error) {
	privKey, didKey, err := did.GenerateDIDKey(ssicrypto.Ed25519)
	if err != nil {
		return nil, err
	}
	pubKey, _, _, err := didKey.Decode()
	if err != nil {
		return nil, err
	}
	return &KeyPair{
		ID:         didKey.String(),
		KeyType:    keyType,
		PrivateKey: privKey,
		PublicKey:  pubKey,
	}, nil
}

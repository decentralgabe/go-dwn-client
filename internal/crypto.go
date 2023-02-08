package internal

import (
	ssicrypto "github.com/TBD54566975/ssi-sdk/crypto"
	"github.com/TBD54566975/ssi-sdk/did"
	"github.com/mr-tron/base58"
)

type KeyPair struct {
	ID               string            `json:"id"`
	KeyType          ssicrypto.KeyType `json:"keyType"`
	PublicKeyBase58  string            `json:"publicKey"`
	PrivateKeyBase58 string            `json:"privateKey,omitempty"`
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
		ID:               didKey.String(),
		KeyType:          keyType,
		PublicKeyBase58:  base58.Encode(pubKey),
		PrivateKeyBase58: base58.Encode(privKey.([]byte)),
	}, nil
}

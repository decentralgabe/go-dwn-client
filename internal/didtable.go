package internal

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

type DIDTable struct {
	dids map[string]KeyPair
}

func NewDIDTableFromConfig(config interface{}) *DIDTable {
	dids := make(map[string]KeyPair)
	if config != nil {
		for did, keyPairJSON := range config.(map[string]interface{}) {
			keyPairBytes, err := json.Marshal(keyPairJSON)
			if err != nil {
				logrus.WithError(err).Error("failed to marshal keyPairJSON from config")
				continue
			}
			var keyPair KeyPair
			if err = json.Unmarshal(keyPairBytes, &keyPair); err != nil {
				logrus.WithError(err).Error("failed to unmarshal keyPairJSON from config")
				continue
			}
			dids[did] = keyPair
		}
	}
	return &DIDTable{dids: dids}
}

func (dt *DIDTable) AddDID(kp KeyPair) error {
	if dt.dids == nil {
		return errors.New("DIDTable is not initialized")
	}
	did := kp.ID
	if gotDID, ok := dt.dids[did]; ok {
		return fmt.Errorf("did<%s> already exists: %s", did, gotDID)
	}
	dt.dids[did] = kp
	if err := SaveDIDTable(dt); err != nil {
		return err
	}
	dt.PrintDIDs()
	return nil
}

func (dt *DIDTable) GetDID(did string) (*KeyPair, error) {
	if dt.dids == nil {
		return nil, errors.New("DIDTable is not initialized")
	}
	if gotDID, ok := dt.dids[did]; ok {
		return &gotDID, nil
	}
	return nil, fmt.Errorf("did<%s> not found", did)
}

func (dt *DIDTable) PrintDIDs() {
	if dt.dids == nil {
		return
	}
	fmt.Println("All DIDs:")
	for did, keyPair := range dt.dids {
		fmt.Printf("%s -> %s:%s\n", did, keyPair.KeyType, keyPair.PublicKeyBase58)
	}
	fmt.Println()
}

func (dt *DIDTable) PrintDID(did string) {
	if dt.dids == nil {
		return
	}
	if keyPair, ok := dt.dids[did]; ok {
		fmt.Printf("did<%s>: \nKey Type: %s\nPublic Key: %s\nPrivate Key: %s\n", did, keyPair.KeyType, keyPair.PublicKeyBase58, keyPair.PrivateKeyBase58)
		return
	}
	fmt.Printf("No keyPair found for did<%s>, try adding one with 'rdr did add <did> <keyType> <pubKeyBase58> <privKeyBase58>'\n", did)
}

func (dt *DIDTable) RemoveDID(did string) error {
	if dt.dids == nil {
		return errors.New("DIDTable is not initialized")
	}
	if _, ok := dt.dids[did]; ok {
		delete(dt.dids, did)
		fmt.Printf("Removed did<%s>\n", did)
		if err := SaveDIDTable(dt); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("did<%s> not found", did)
}

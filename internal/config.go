package internal

import (
	"encoding/json"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Routes map[string]string  `json:"routes"`
	DIDs   map[string]KeyPair `json:"dids"`
}

func ToStringMapString(i map[string]interface{}) map[string]string {
	m := make(map[string]string)
	for k, v := range i {
		m[k] = v.(string)
	}
	return m
}

func ReadConfig() (*Config, error) {
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func SaveRouteTable(rt *RouteTable) error {
	cfg, err := ReadConfig()
	if err != nil {
		return err
	}
	cfg.Routes = rt.routes
	configJSON, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(viper.ConfigFileUsed(), configJSON, 0755)
}

func SaveDIDTable(dt *DIDTable) error {
	cfg, err := ReadConfig()
	if err != nil {
		return err
	}
	cfg.DIDs = dt.dids
	configJSON, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(viper.ConfigFileUsed(), configJSON, 0755)
}

package internal

import (
	"encoding/json"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Routes map[string]string `json:"routes"`
}

func ToStringMapString(i map[string]interface{}) map[string]string {
	m := make(map[string]string)
	for k, v := range i {
		m[k] = v.(string)
	}
	return m
}

func ReadRouteTable() (*RouteTable, error) {
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return NewRouteTableFromConfig(c.Routes), nil
}

func SaveRouteTable(rt *RouteTable) error {
	updateConfig := Config{Routes: rt.routes}
	configJSON, err := json.Marshal(updateConfig)
	if err != nil {
		return err
	}
	return os.WriteFile(viper.ConfigFileUsed(), configJSON, 0755)
}

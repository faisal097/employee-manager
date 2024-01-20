package config

import (
	"encoding/json"
	"log"
	"os"
)

type HttpServerConfig struct {
	Port string `json:"port"`
}

func ParseConfig(filepath string) (*HttpServerConfig, error) {
	h := &HttpServerConfig{}
	b, err := os.ReadFile(filepath)
	if err != nil {
		log.Println("Error in config file read: ", err.Error())
		return nil, err
	}
	err = json.Unmarshal(b, h)
	if err != nil {
		log.Println("Error in config unmarshal: ", err.Error())
		return nil, err
	}
	return h, nil
}

func (h *HttpServerConfig) SetPort(port string) {
	h.Port = port
}

func (h *HttpServerConfig) GetPort() string {
	return h.Port
}

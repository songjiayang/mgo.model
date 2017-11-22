package model

import "encoding/json"

type Config struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Passwd   string `json:"password"`
	Database string `json:"database"`
	Mode     string `json:"mode"`
	Pool     int    `json:"pool"`
	Timeout  int    `json:"timeout"`
}

func NewConfig(data []byte) (*Config, error) {
	var mgoCfg Config

	err := json.Unmarshal(data, &mgoCfg)
	if err != nil {
		return nil, err
	}

	return &mgoCfg, err
}

func (c *Config) Copy() *Config {
	config := *c

	return &config
}

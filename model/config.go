package model

type Config struct {
	Service string              `json:"service"`
	Data    []map[string]string `json:"data"`
}

func NewConfig() *Config {
	return &Config{Data: []map[string]string{make(map[string]string)}}
}

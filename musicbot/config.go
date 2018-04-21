package musicbot

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	DefaultConfigFileLocation = "config.json"
	DefaultWhiteListFile      = "whitelist.txt"
	DefauultMaster            = "swiltink"
)

type Config struct {
	WhiteListFile string    `json:"whitelistFile"`
	Master        string    `json:"master"`
	Irc           IRCConfig `json:"irc"`
}

type IRCConfig struct {
	Server   string `json:"server"`
	Channel  string `json:"channel"`
	Nick     string `json:"nick"`
	RealName string `json:"realname"`
	Pass     string `json:"pass"`
	Ssl      bool   `json:"ssl"`
}

func (config *Config) applyDefaults() {
	config.WhiteListFile = DefaultWhiteListFile
	config.Master = DefauultMaster
}

func LoadConfig(fileLocation string) (*Config, error) {
	file, err := os.Open(fileLocation)

	if err != nil {
		return nil, fmt.Errorf("unable to load config: %v", err)
	}

	config := &Config{}
	config.applyDefaults()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)

	if err != nil {
		return nil, fmt.Errorf("unable to decode config file: %v", err)
	}

	return config, nil
}

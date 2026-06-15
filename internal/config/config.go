package config 

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct{
	DBURL string `json:"db_url"`
	USERNAME string `json:"current_user_name"`
}

func getConfigPath() (string, error){
	homeDir, err := os.UserHomeDir()
	if err != nil{
		return "", err
	}
	jsonDir := filepath.Join(homeDir, ".gatorconfig.json")
	return jsonDir, nil
}

func Read() (Config, error) {
	jsonDir, err := getConfigPath()
	if err != nil{
		return Config{}, err
	}

	var config Config
	file, err := os.Open(jsonDir)
	if err != nil{
		return Config{}, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil{
		return Config{}, err
	}
	return config, nil
}

func (c *Config) SetUser(username string) error {
	c.USERNAME = username
	return write(*c)
}


func write(cfg Config) error{
	jsonDir, err := getConfigPath()
	if err != nil{
		return err
	}
	file, err := os.Create(jsonDir)	
	if err != nil{
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil{
		return err
	}
	return nil
}




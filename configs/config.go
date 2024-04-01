package Config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Email EmailConfig `json:"email"`
	Mysql MysqlConfig `json:"mysql"`
	Redis RedisConfig `json:"redis"`
}

type EmailConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type MysqlConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"dbname"`
}
type RedisConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

func ReadConfig() Config {
	dataBytes, err := os.ReadFile("conf.yaml")
	if err != nil {
		log.Println(err)
		return Config{}
	}
	config := Config{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		log.Println(err)
		return Config{}
	}
	return config
}

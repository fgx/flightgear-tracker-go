
package config

import(
	"log"
	"os"
	"io/ioutil"
	"encoding/json"
)

type Config struct{
	DbServer string `json:"db_server"`
	DbUser string `json:"db_user"`
	DbPassword string `json:"db_password"`
	Database string `json:"database"`
	HttpPort int `json:"http_port"`
}

func LoadConfig() *Config {
	file, err := ioutil.ReadFile("config.json")
  	if err != nil {
    	log.Println("open config: ", err)
    	os.Exit(1) 
  	}

	conf := new(Config)
	if err = json.Unmarshal(file, conf); err != nil {
    	log.Println("parse config: ", err)
    	os.Exit(1) 
  	}
  	return conf
}
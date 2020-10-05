package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//Config struct
type Config struct {
	IPAddress string `json:"server_address"`
	Port      string `json:"server_port"`
	User      string `json:"username"`
	Password  string `json:"password"`
}

//Conf - Global con la configuracion
var Conf Config

//LoadConfig carga configuracion
func LoadConfig(path string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	Conf = Config{}
	err = json.Unmarshal(file, &Conf)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

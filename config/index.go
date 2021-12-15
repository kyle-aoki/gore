package config

import (
	"encoding/json"
	"fmt"
	"gore/util"

	"log"

	"os"

	_ "github.com/go-sql-driver/mysql"
)

type IGoreConfig struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Subdatabase string `json:"subdatabase"`

	S3_REGION string `json:"S3_REGION"`
	S3_BUCKET string `json:"S3_BUCKET"`
}

var Gore IGoreConfig

func LoadGoreConfig() {
	envPath := util.HomeDir() + "/env/gore.json"

	f, err := os.ReadFile(envPath)
	if err != nil {
		fmt.Println("Must have gore.json at ~/env/gore.json.")
		log.Printf("Error reading %s: %+v\n", envPath, err)
		log.Fatalf(IGoreConfigStructText)
	}

	Gore = IGoreConfig{}
	err = json.Unmarshal(f, &Gore)
	if err != nil {
		log.Printf("Error loading json into GoreConfig struct: %+v", err)
		log.Fatal(IGoreConfigStructText)
	}
}

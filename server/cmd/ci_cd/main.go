package main

import (
	"log"
	"ci_cd/configs"
)

func main() {
	cfg := configs.GetConfig()
	dbUrl := configs.GetDBUrl()
	log.Println(cfg)
	log.Println(dbUrl)
}
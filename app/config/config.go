/*
	Project: hps
	Author: Sarvesh Chitko (chitkosarvesh@gmail.com) (sarvesh)
	Date: 22-March-2018 4:47 PM

	Package -> config

	This package provides all the configuration required to run HPS
*/
package config

import (
	"hps/app/gonfig"
	"os"
	"log"
)
var config gonfig.Gonfig
//	Reads the configuration file stored in ${ROOT}/config/hps.json
func Read()(jsonGonfig gonfig.Gonfig)	{
	f,err := os.Open("config/hps.json")
	defer f.Close()
	if err != nil {
		log.Fatal("ERROR ",err)
	} else {
		config,err := gonfig.FromJson(f)
		if err != nil {
			log.Fatal("Error ", err)

		} else {
			return config
		}
	}
	return config
}
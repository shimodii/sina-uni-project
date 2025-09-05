package config

import (
  "log"
  "os"
)

func UploadsPath(){
  if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		err := os.Mkdir("./uploads", os.ModePerm)
		if err != nil {
			log.Fatalf("failed to create uploads folder: %v", err)
		}
	}
}

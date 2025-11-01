package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	loadConfig()
	fmt.Print("windows tools")
}

func loadConfig() {
	confFile := "conf.env"

	if _, err := os.Stat(confFile); os.IsNotExist(err) {
		log.Fatalf(`
====================================================
Config file not found: %s
====================================================
`, confFile)
	}

	fmt.Printf(`
====================================================
Loading config...
====================================================
`)

	if err := godotenv.Load(confFile); err != nil {
		log.Fatalf(`
====================================================
Error loading %s: %v
====================================================
`, confFile, err)
	}

	fmt.Printf(`
====================================================
Config loaded: %s
====================================================
`, confFile)
}

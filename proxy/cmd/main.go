package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/jrjaro18/elastic-raft-go/proxy/internal/proxy"
)

func main() {
	if err := godotenv.Load(); err != nil {
        log.Fatalln("No .env file found")
    }

	apikey, exists := os.LookupEnv("API_KEY")
	
	if (!exists) {
		log.Fatalln("Api key not found in the .env")
	}	

	p1 := proxy.NewProxy(":6000", apikey)
	// p2 := proxy.NewProxy(":6001", apikey)
	// p3 := proxy.NewProxy(":6003", apikey)
	
	p1.Start()
	// p2.Start()
	// p3.Start()
}
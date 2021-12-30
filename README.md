# consul
consul config

### Example

`
package main

import (
	"log"
	"os"

	c "github.com/alexsukhrin/consul"
)

func main() {
	consulConfig := new(c.ConsulConfig)
	consulConfig.HOST = os.Getenv("CONSUL_HOST")
	consulConfig.PORT = os.Getenv("CONSUL_PORT")
	consulConfig.STAGE = os.Getenv("STAGE")
	consulConfig.SERVICE = os.Getenv("CONSUL_SERVICE")
	consulConfig.PATH_TOKEN = os.Getenv("CONSUL_SECRET_TOKEN")

	consul := new(c.Consul)
	consul.Config = consulConfig

	config := consul.GetConfig()

	log.Printf("Config %s", config)
}

`
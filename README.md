# consul
consul config

### Example

```
package main

import (
	c "github.com/alexsukhrin/consul"
	"log"
	"os"
)

var ConsulHost, ConsulPort, ConsulService, PathToken, Stage string

func init() {
	ConsulHost = os.Getenv("CONSUL_HOST")
	ConsulPort = os.Getenv("CONSUL_PORT")
	ConsulService = os.Getenv("CONSUL_SERVICE")
	PathToken = os.Getenv("CONSUL_SECRET_TOKEN")
	Stage = os.Getenv("STAGE")
}

func main() {
	consulConfig := new(c.ConsulConfig)
	consulConfig.HOST = ConsulHost
	consulConfig.PORT = ConsulPort
	consulConfig.STAGE = Stage
	consulConfig.SERVICE = ConsulService
	consulConfig.PATH_TOKEN = PathToken

	consul := new(c.Consul)
	consul.Config = consulConfig

	config := consul.GetConfig()
	log.Printf("Config %s", config)
}

```
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
	consul := new(c.Consul)
	consul.HOST = ConsulHost
	consul.PORT = ConsulPort
	consul.STAGE = Stage
	consul.SERVICE = ConsulService
	consul.PATH_TOKEN = PathToken
	consul.Address = consul.BuildAddressConsul()
	consul.ConfigPath = consul.BuildPathConfig()
	consul.Token = consul.GetToken()
	
	config := consul.GetConfig()
	log.Printf("Config %s", config)
}

```
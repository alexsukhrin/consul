package consul

import (
	"fmt"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	consul := new(Consul)
	consul.Host = os.Getenv("CONSUL_HOST")
	consul.Port = os.Getenv("CONSUL_PORT")
	consul.Stage = os.Getenv("STAGE")
	consul.ServiceName = os.Getenv("CONSUL_SERVICE")
	consul.TokenPath = os.Getenv("CONSUL_SECRET_TOKEN")
	consul.ConfigAddress = consul.Address()
	consul.ConfigPath = consul.Path()
	consul.TokenKey = consul.Token()

	config := consul.Config()
	t.Log(fmt.Sprintf("Config %s", config))
}

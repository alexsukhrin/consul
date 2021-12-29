package consul

import (
	"fmt"
	"github.com/goccy/go-yaml"
	consul "github.com/hashicorp/consul/api"
	"io/ioutil"
	"log"
)

type Token struct {
	Value string `yaml:"token"`
}

type ConsulConfig struct {
	HOST, PORT, STAGE, SERVICE, PATH_CONFIG, TOKEN string
}

func BuildAddresConsul(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}

func BuildPathConfig(stage, service string) string {
	return fmt.Sprintf("DigitalCore/%s/%s/config", stage, service)
}

func GetConfig(c *ConsulConfig) []byte {
	newConfig := consul.DefaultConfig()
	newConfig.Token = c.TOKEN

	newConfig.Address = BuildAddresConsul(c.HOST, c.PORT)
	client, err := consul.NewClient(newConfig)

	if err != nil {
		log.Fatal("Error not connect to consul")
	}

	log.Println("Connect to consul successfully")

	// Get a handle to the KV API
	kv := client.KV()

	// Lookup the pair
	params := consul.QueryOptions{}
	params.Token = c.TOKEN

	pair, _, err := kv.Get(c.PATH_CONFIG, &params)

	if err != nil {
		log.Fatal(fmt.Sprintf("Error not valid path %s error %s", c.PATH_CONFIG, err))
	}

	return pair.Value
}

func getToken(path string) Token {
	envToken, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	token := Token{}
	err = yaml.Unmarshal([]byte(envToken), &token)

	if err != nil {
		log.Fatal(err)
	}

	return token
}

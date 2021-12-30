package consul

import (
	"fmt"
	"github.com/goccy/go-yaml"
	apiConsul "github.com/hashicorp/consul/api"
	"io/ioutil"
	"log"
)

type ConsulConfig struct {
	HOST, PORT, STAGE, SERVICE, PATH_TOKEN string
}

type Consul struct {
	Config *ConsulConfig
}

type Token struct {
	Token string `yaml:"token"`
}

func (c *Consul) buildAddresConsul() string {
	return fmt.Sprintf("%s:%s", c.Config.HOST, c.Config.PORT)
}

func (c *Consul) buildPathConfig() string {
	return fmt.Sprintf("DigitalCore/%s/%s/config", c.Config.STAGE, c.Config.STAGE)
}

func (c *Consul) GetConfig(conf *ConsulConfig) []byte {
	consul := new(Consul)
	consul.Config = conf

	token := consul.getToken().Token

	newConfig := apiConsul.DefaultConfig()
	newConfig.Token = token
	newConfig.Address = consul.buildAddresConsul()

	client, err := apiConsul.NewClient(newConfig)
	if err != nil {
		log.Fatal("Error not connect to consul")
	}

	log.Println("Connect to consul successfully")

	// Get a handle to the KV API
	kv := client.KV()

	// Lookup the pair
	params := apiConsul.QueryOptions{}
	params.Token = token
	
	path := consul.buildPathConfig()
	pair, _, err := kv.Get(path, &params)

	if err != nil {
		log.Fatal(fmt.Sprintf("Error not valid path %s error %s", path, err))
	}

	return pair.Value
}

func (c *Consul) getToken() Token {
	envToken, err := ioutil.ReadFile(c.Config.PATH_TOKEN)
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

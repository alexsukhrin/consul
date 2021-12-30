package consul

import (
	"fmt"
	"github.com/goccy/go-yaml"
	apiConsul "github.com/hashicorp/consul/api"
	"io/ioutil"
	"log"
)

type Consul struct {
	HOST, PORT, STAGE, SERVICE, PATH_TOKEN, Address, ConfigPath, Token string
}

type Token struct {
	Token string `yaml:"token"`
}

func (consul *Consul) BuildAddressConsul() string {
	return fmt.Sprintf("%s:%s", consul.HOST, consul.PORT)
}

func (consul *Consul) BuildPathConfig() string {
	return fmt.Sprintf("DigitalCore/%s/%s/config", consul.STAGE, consul.SERVICE)
}

func (consul *Consul) GetConfig() []byte {
	newConfig := apiConsul.DefaultConfig()
	newConfig.Token = consul.Token
	newConfig.Address = consul.Address

	client, err := apiConsul.NewClient(newConfig)
	if err != nil {
		log.Fatal("Error not connect to consul")
	}

	log.Println("Connect to consul successfully")

	// Get a handle to the KV API
	kv := client.KV()

	// Lookup the pair
	params := apiConsul.QueryOptions{}
	params.Token = consul.Token

	pair, _, err := kv.Get(consul.ConfigPath, &params)

	if err != nil {
		log.Fatal(fmt.Sprintf("Error not valid path %s error %s", consul.ConfigPath, err))
	}

	return pair.Value
}

func (consul *Consul) GetToken() string {
	envToken, err := ioutil.ReadFile(consul.PATH_TOKEN)
	if err != nil {
		log.Fatal(err)
	}

	token := Token{}
	err = yaml.Unmarshal([]byte(envToken), &token)

	if err != nil {
		log.Fatal(err)
	}

	return token.Token
}

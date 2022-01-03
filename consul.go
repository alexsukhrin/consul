package consul

import (
	"fmt"
	"github.com/goccy/go-yaml"
	apiConsul "github.com/hashicorp/consul/api"
	"io/ioutil"
	"log"
)

type Consul struct {
	Host, Port, Stage, ServiceName, TokenKey, ConfigAddress, ConfigPath, TokenPath string
}

type TokenKey struct {
	Token string `yaml:"token"`
}

func (consul *Consul) Address() string {
	return fmt.Sprintf("%s:%s", consul.Host, consul.Port)
}

func (consul *Consul) Path() string {
	return fmt.Sprintf("DigitalCore/%s/%s/config", consul.Stage, consul.ServiceName)
}

func (consul *Consul) Config() []byte {
	newConfig := apiConsul.DefaultConfig()
	newConfig.Token = consul.TokenKey
	newConfig.Address = consul.ConfigAddress

	client, err := apiConsul.NewClient(newConfig)
	if err != nil {
		log.Fatal("Error not connect to consul")
	}

	log.Println("Connect to consul successfully")
	kv := client.KV()

	params := apiConsul.QueryOptions{}
	params.Token = consul.TokenKey

	pair, _, err := kv.Get(consul.ConfigPath, &params)

	if err != nil {
		log.Fatal(fmt.Sprintf("Error not valid path %s error %s", consul.ConfigPath, err))
	}

	return pair.Value
}

func (consul *Consul) Token() string {
	envToken, err := ioutil.ReadFile(consul.TokenPath)
	if err != nil {
		log.Fatal(err)
	}

	token := new(TokenKey)
	err = yaml.Unmarshal([]byte(envToken), token)

	if err != nil {
		log.Fatal(err)
	}

	return token.Token
}

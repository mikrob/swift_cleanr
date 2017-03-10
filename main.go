package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/ncw/swift"
	"github.com/spf13/viper"
	"io/ioutil"
)

var (
	config    = flag.String("f", "./config.yml", "absolute path to the yaml config file for openstack")
	container = flag.String("c", "screencapture-staging", "name of container to clean")
)

func main() {

	viper.SetConfigType("yaml")

	data, err := ioutil.ReadFile(*config)
	if err != nil {
		panic("Cannot read config")
	}

	viper.ReadConfig(bytes.NewBuffer(data))

	if viper.Get("openstack") == nil {
		panic("Cannot found openstack in config")
	}

	username := viper.GetString("openstack.username")
	api_key := viper.GetString("openstack.api_key")
	auth_url := viper.GetString("openstack.auth_url")
	authtenant_name := viper.GetString("openstack.authtenant_name")
	region := viper.GetString("openstack.region")

	// Create a connection
	c := swift.Connection{
		UserName: username,
		ApiKey:   api_key,
		AuthUrl:  auth_url,
		Tenant:   authtenant_name, // Name of the tenant (v2 auth only)
		Region:   region,
	}
	// Authenticate
	errAuth := c.Authenticate()
	if errAuth != nil {
		panic(errAuth)
	}
	// List all the containers
	containers, err := c.ContainerNames(nil)
	fmt.Println(containers)
	swiftContainer, _, errContainer := c.Container(*container)

	fmt.Println("Count : ")
	fmt.Println(swiftContainer.Count)

	if errContainer != nil {
		fmt.Println("Cannont retrieve container ", *container)
	}
	files, errNames := c.ObjectNames(*container, nil)
	if errNames != nil {
		panic("Cannot list container " + *container)
	}

	fmt.Println("Starting delete")
	c.BulkDelete(*container, files)
	fmt.Println("Delete done")
}

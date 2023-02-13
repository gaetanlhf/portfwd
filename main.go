package main

import (
	"io/ioutil"
	"os"
	"os/signal"

	"log"

	"gopkg.in/yaml.v3"
)

var (
	version   string
	buildTime string
	forward   ForwardStruct
	config    ConfigStruct
)

type ForwardStruct struct {
	Protocol string `yaml:"protocol"`
	From     string `yaml:"from"`
	To       string `yaml:"to"`
}

type ConfigStruct struct {
	Forward []ForwardStruct `yaml:"forward"`
}

func main() {
	log.Printf("Starting portfwd %s build on %s", version, buildTime)
	configFilePath := os.Getenv("PORTFWD_CONFIG_FILE_PATH")
	log.Printf("Loading configuration file located at %s", configFilePath)
	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	if len(config.Forward) == 0 {
		log.Fatal("Nothing to forward! Please check your configuration.")
	}

	for _, forward = range config.Forward {
		switch forward.Protocol {
		case "tcp", "tcp4", "tcp6":
			go tcpForward(forward)
			log.Printf("Forwarding TCP traffic between %s and %s.", forward.From, forward.To)
		case "udp", "udp4", "udp6":
			go udpForward(forward)
			log.Printf("Forwarding UDP traffic between %s and %s.", forward.From, forward.To)
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
}

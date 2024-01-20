package main

import (
	"flag"
	"log"

	"github.com/faisal097/employee-manager/config"
	"github.com/faisal097/employee-manager/httpserver"
)

var (
	confPath *string
	conf     *config.HttpServerConfig
)

func init() {
	confPath = flag.String("conf", "../config.json", "config file path")
	flag.Parse()

	//var err error
	conf, err = config.ParseConfig(*confPath)
	if err != nil {
		log.Fatalf("Error in ParseConfig: %s", err.Error())
	}
}

func main() {
	defer func() {
		log.Println("Exiting main!")
	}()

	done := make(chan struct{})
	go func() {
		err := httpserver.Serve(conf)
		log.Printf("Http Server stopped, Err: %v", err.Error())
	}()
	<-done
}

package main

import (
	"flag"
	"fmt"
	"github.com/corentone/golang-boilerplate/pkg/boilerplatebin"
	"github.com/corentone/golang-boilerplate/pkg/boilerplatebin/config"
)

func main() {
	fmt.Println("Starting boilerplatebin...")

	err := flag.CommandLine.Parse([]string{}) // This line is necessary to make flag think its run.
	if err != nil {
		panic(err)
	}
	err = flag.Set("logtostderr", "true")
	if err != nil {
		panic(err)
	}

	c := &config.Config{}
	err = c.Parse()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Config used is: %+v\n", c)

	err = flag.Set("v", c.LogLevel)
	if err != nil {
		panic(err)
	}

	boilerplatebin.Run(c)
}

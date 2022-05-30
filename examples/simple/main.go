package main

import (
	"fmt"
	"time"

	"github.com/crgimenes/goconfig"
)

/*
step 1: Declare your configuration struct,
it may or may not contain substructures.
*/

type mongoDB struct {
	Host string `cfgDefault:"example.com" cfgHelper:"mongoDB host URL"`
	Port int    `cfgDefault:"999" cfgHelper:"mongoDB http port"`
}

type configTest struct {
	Timeout   time.Duration `cfg:"timeout" cfgDefault:"5m"`
	Domain    string
	DebugMode bool `json:"db" cfg:"db" cfgDefault:"false"`
	MongoDB   mongoDB
}

func main() {

	// step 2: Instantiate your structure.
	config := configTest{}

	// step 3: Pass the instance pointer to the parser
	err := goconfig.Parse(&config)
	if err != nil {
		fmt.Println(err)
		return
	}

	/*
	   The parser populated your struct with the data
	   it took from environment variables and command
	   line and now you can use it.
	*/

	println("config.Timeout.....:", config.Timeout)
	println("config.Domain......:", config.Domain)
	println("config.DebugMode...:", config.DebugMode)
	println("config.MongoDB.Host:", config.MongoDB.Host)
	println("config.MongoDB.Port:", config.MongoDB.Port)
}

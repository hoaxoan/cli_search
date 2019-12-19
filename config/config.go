package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var Config ConfigSchema

func init() {
	viper.SetConfigName("default")   // name of config file (without extension)
	viper.AddConfigPath("./config") // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

type ConfigSchema struct {
	Data struct {
		Organization      string `mapstructure:"Organization"`
		Ticket     string `mapstructure:"Ticket"`
		User string `mapstructure:"User"`
	} `mapstructure:"Data"`
}

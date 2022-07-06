package config

import (
	"app/cmd"
	"fmt"
	"github.com/spf13/viper"
)

type Configurations struct {
	Storage   Storage
	WebServer WebServer
}

type Storage struct {
	Redis RedisConfiguration
}

type RedisConfiguration struct {
	Hostname string
	Port     int
	Password string
	Database int
}

type WebServer struct {
	Port         string
	Tpl          string
	TplPattern   string `mapstructure:"tpl_pattern"`
	StaticPath   string `mapstructure:"static_path"`
	StaticPrefix string `mapstructure:"static_prefix"`
}

func Load(args *cmd.Args) Configurations {

	// Set the file name of the configurations file
	viper.SetConfigName("config." + args.Env)

	// Set the path to look for the configurations file
	viper.AddConfigPath("./config")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration Configurations

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Error reading config file, %s", err))
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		panic(fmt.Sprintf("Unable to decode into struct, %v", err))
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())

	return configuration
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/KawashiroNitori/MoeManager/cmd"
	"github.com/kardianos/service"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"os"
)

func main() {
	viper.SetConfigName("config.toml")
	fmt.Println(os.Executable())
	if service.Interactive() {
		viper.AddConfigPath(".")
	}
	viper.AddConfigPath(lo.Must(os.Executable()))
	lo.Must0(viper.ReadInConfig())
	viper.WatchConfig()

	cmd.Execute()
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/KawashiroNitori/MoeManager/cmd"
	"github.com/kardianos/service"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	if service.Interactive() {
		viper.AddConfigPath(".")
	}
	viper.AddConfigPath(filepath.Dir(lo.Must(os.Executable())))
	lo.Must0(viper.ReadInConfig())
	viper.WatchConfig()

	cmd.Execute()
}

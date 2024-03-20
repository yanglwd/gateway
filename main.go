/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/spf13/viper"
	"github.com/yanglwd/gateway/cmd"
)

func main() {
	cmd.Execute()

	println(viper.Get("addr").(string))
}

// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	Password string
	Devicegroup string
	Shared bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "pan-cli",
	Short: "Paloalto firewall (PANOS) command line. ",
	Long: `Copyright 2017 @IrekRomaniuk. Package using github.com/scottdware/go-panos by Scott Ware to interact with Palo Alto and Panorama devices using the XML API.

`,
// Uncomment the following line if your bare application
// has an action associated with it:
//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Yaml config file ")
	//i.e. --config "C:\Users\irekromaniuk\Vagrant\trusty64\src\github.com\IrekRomaniuk\pan-cli\pan-cli.yaml"
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//RootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	RootCmd.PersistentFlags().StringP("device", "d", "", "Device to connect")
	RootCmd.PersistentFlags().StringVarP(&Devicegroup, "devicegroup","g", "","Panorama devicegroup")
	RootCmd.PersistentFlags().BoolVarP(&Shared, "shared","s", false,"True for shared objects")
	RootCmd.PersistentFlags().StringP("login", "u", "admin", "Login name")
	RootCmd.PersistentFlags().StringVarP(&Password, "password","p", "","Password")
	viper.BindPFlag("device", RootCmd.PersistentFlags().Lookup("device"))
	viper.BindPFlag("login", RootCmd.PersistentFlags().Lookup("login"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	//viper.SetConfigType("yaml")
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}
	//viper.SetConfigName(".pan-cli") // name of config file (without extension)
	//viper.AddConfigPath("$HOME")  // adding home directory as first search path
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

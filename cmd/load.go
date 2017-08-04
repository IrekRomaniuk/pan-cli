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
	"github.com/spf13/cobra"
	. "github.com/scottdware/go-panos"
	//. "github.com/IrekRomaniuk/go-panos"
	"log"
	"github.com/spf13/viper"
)

// loadCmd represents the load command

var (
	file string
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Loads address object from csv file",
	Long: `Command takes a .csv file with the following format: name,type,address,description,address-group
'name' is what you want the address object to be called.
'type' is one of: ip, range, or fqdn.
'address' is the address of the object.
'description' is optional, just leave the field blank if you do not want one.
'address-group' will assign the object to the given address-group if you wish (leave blank if you do not want to add it to a group).
If creating shared address objects on a Panorama device, then specify "true" for the shared parameter, and omit the device-group.
If not creating a shared object, then just specify "false."
	`,
	Run: func(cmd *cobra.Command, args []string) {
		host := viper.GetString("device") // case-insensitive Setting & Getting
		login := viper.GetString("login")
		pan1session, err := NewSession(host, login, Password)
		if err != nil {
			log.Fatal(err)
		}
		//csv file format: name,type,address,description,address-group
		//fmt.Println(file, Shared, Devicegroup, host, login, Password)
		err = pan1session.CreateAddressFromCsv(file, Shared, Devicegroup)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(loadCmd)
	loadCmd.Flags().StringVarP(&file, "file","f", "","Csv file to load from (full path)")
	//i.e."C:\Users\irekromaniuk\Vagrant\trusty64\src\github.com\IrekRomaniuk\pan-cli\cmd\address.csv"
}

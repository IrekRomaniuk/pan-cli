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
	//"fmt"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	. "github.com/scottdware/go-panos"
	"log"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit to a Panorama or firewall",
	Long: `Commit issues a commit to a Panorama device, with the given 'devicegroup.'  
or fiewall if 'devicegroup.' is empty
	`,
	Run: func(cmd *cobra.Command, args []string) {
		host := viper.GetString("device") // case-insensitive Setting & Getting
		login := viper.GetString("login")
		//fmt.Println(Devicegroup, host, login, Password)
		pan1session, err := NewSession(host, login, Password)
		if err != nil {
			log.Fatal(err)
		}
		switch Devicegroup {
			case "":
				err = pan1session.Commit()				
			default: 
				err = pan1session.CommitAll(Devicegroup)
							
		}		
		if err != nil {
					log.Fatal(err)
				}
	},
}

func init() {
	RootCmd.AddCommand(commitCmd)
}

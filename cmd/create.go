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
	"log"
	//. "github.com/IrekRomaniuk/go-panos"
	. "github.com/scottdware/go-panos"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	object string
	value string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create objects on the device",
	Long: `CreateAddress will add a new address object to the device.
addrtype should be one of: ip, range, or fqdn.
If creating a shared address object on a Panorama device, then specify "true" for the shared parameter, and omit the device-group.
If not creating a shared object, then just specify "false."`,
	Run: func(cmd *cobra.Command, args []string) {
		/*pan1 := PaloAlto{
			Host: "192.168.3.101",
			Key: "LUFRPT1jRVRDTmo1VVpCZ2wwa3hCU1Roc1pWUVh0VTA9QU5jREpOWVFCaFBXbW5xZ214UU9zQT09",
		}*/
		host := viper.GetString("device") // case-insensitive Setting & Getting
		login := viper.GetString("login")
		pan1session, err := NewSession(host, login, Password)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(Shared, Devicegroup, host, login, Password)
		//fmt.Printf("%+v\n", pan1session)
		err = pan1session.CreateAddress(object,"ip",value,"", Devicegroup)  //Shared depreciated ??
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&object, "object","o", "","Object name")
	createCmd.Flags().StringVarP(&value, "value","v", "","Object value")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

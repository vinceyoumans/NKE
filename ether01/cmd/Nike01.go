/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/spf13/cobra"
)

// Nike01Cmd represents the Nike01 command
var Nike01Cmd = &cobra.Command{
	Use:   "Nike01",
	Short: "An experiment to query the Ethereum BC",
	Long:  `Another Ether experiment`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Nike01 called")

		do01()
	},
}

func init() {
	rootCmd.AddCommand(Nike01Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// Nike01Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// Nike01Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func do01() {

	// Create an account
	key, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println("----   crypto.Generate... something wrong")
	}

	// Get the address
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	fmt.Println("========   address ====")
	fmt.Println(address)
	// 0x8ee3333cDE801ceE9471ADf23370c48b011f82a6

	// Get the private key
	privateKey := hex.EncodeToString(key.D.Bytes())
	// 05b14254a1d0c77a49eae3bdf080f926a2df17d8e2ebdf7af941ea001481e57f
	fmt.Println("========   private key ====")
	fmt.Println(privateKey)

}

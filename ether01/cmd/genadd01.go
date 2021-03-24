/*
Copyright © 2021 Vincent Youmans me@vyoumans.com
*/

package cmd

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// 	"github.com/ethereum/go-ethereum/crypto/sha3"

// genadd01Cmd represents the genadd01 command
var genadd01Cmd = &cobra.Command{
	Use:   "genadd01",
	Short: "Will Generate REAL Ethereum Addresses",
	Long:  `This will generate a Real Ethereum address on the Ethereum Block Chain`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("genadd01 This will create a REAL Ether Address that you could send a token")

		printHivy()

		genkey()

		genkey02()

	},
}

func init() {
	rootCmd.AddCommand(genadd01Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genadd01Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genadd01Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printHivy() {
	fmt.Println("------   hi vy ----")
}

func genkey() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	// hash := sha3.NewKeccak256()
	// hash.Write(publicKeyBytes[1:])
	// fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e

}

func genkey02() {
	data := []byte("hello")
	hash := sha256.Sum256(data)
	fmt.Printf("%x", hash[:]) // 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
}

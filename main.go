package main

import (
	"fmt"
	"github.com/parvizi/bms/sdkconfig"
	"os"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Network parameters 
//		OrdererID: "orderer.hf.chainhero.io",

		// Channel parameters
//		ChannelID:     "chainhero",
//		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/chainHero/heroes-service/fixtures/artifacts/chainhero.channel.tx",

		// Chaincode parameters
//		ChainCodeID:     "heroes-service",
//		ChaincodeGoPath: os.Getenv("GOPATH"),
//		ChaincodePath:   "github.com/chainHero/heroes-service/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "org1",
		ConfigFile:      "config_sdk.yaml",

		// User parameters
		UserName: "Admin",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}
	// Close SDK
	defer fSetup.CloseSDK()	
}

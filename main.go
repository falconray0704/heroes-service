package main

import (
	"bufio"
	"fmt"
	"github.com/chainHero/heroes-service/blockchain"
	"io"
	"log"
	"os"
	"strconv"
	"syscall"
	"time"
)

func main() {
	var err error
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Network parameters
		OrdererID: "orderer.hf.chainhero.io",

		// Channel parameters
		ChannelID:     "chainhero",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/chainHero/heroes-service/fixtures/artifacts/chainhero.channel.tx",

		// Chaincode parameters
		//ChainCodeID:     "heroes-service",
		ChainCodeID:     "secc-service",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/chainHero/heroes-service/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "org1",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err = fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}
	// Close SDK
	defer fSetup.CloseSDK()

	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
		return
	}

	/*
	// Launch the web application listening
	app := &controllers.Application{
		Fabric: &fSetup,
	}
	web.Serve(app)
	*/

	////////////////////////////////

	var pipeFile = "pipLogs"
	//var file *os.File
	os.Remove(pipeFile)
	e := syscall.Mkfifo(pipeFile, 0666)
	if e != nil {
		log.Fatal("Make named pipe file error:", e)
	}
	//go scheduleWrite()
	fmt.Println("open a named pipe file for read.")
	file, e := os.OpenFile(pipeFile, os.O_CREATE, os.ModeNamedPipe)
	if e != nil {
		log.Fatal("Open named pipe file error:", e)
	}

	reader := bufio.NewReader(file)
	var line []byte

	/*
	for {
		line, e := reader.ReadBytes('\n')
		if e == nil {
			fmt.Printf("load string:" + string(line))
		} else {
			//fmt.Printf("--- err:%s \n", err.Error())
			time.Sleep(time.Millisecond * 100)
		}
	}
	*/

	for {
		line, err = reader.ReadBytes('\n')
		if err == nil {
			key := strconv.FormatInt(time.Now().Unix(), 10)
			//key := "hello"
			fmt.Printf("\n\n### Pushing key:%s log:%s ", key, string(line))
			_, err = fSetup.InvokeKeySet(key, string(line))
			if err != nil {
				fmt.Printf("--- Published key:%s log:%s fail! ---\n", key, string(line))
				continue
			} else {
				fmt.Printf("+++ Published key:%s log:%s success! ---\n", key, string(line))
			}
			fmt.Printf("=== Querying privious log with key:%s \n", key)
			var v string
			v, err = fSetup.QueryKeyGet(key)
			if err != nil {
				fmt.Printf("--- Queried with key:%s fail! ---\n", key)
			} else {
				fmt.Printf("+++ Queried with key:%s get log:%s success! ---\n", key, v)
			}
		} else if err == io.EOF {
			time.Sleep(time.Millisecond * 100)
		} else {
			fmt.Printf("------------------------------\n")
		}
	}
}

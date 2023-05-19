package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

var (
	ipfsPath string
	Port     string
	Muladd ="/ip4/127.0.0.1/tcp/" 	  
)

func InitialiseIpfsDirectory() error {
	fmt.Println("Enter the ipfs file number:")
	var num string
	fmt.Scan(&num)
	Usr, err := user.Current()
	if err != nil {
		return err
	}

	// Construct the IPFS path using the home directory and user input
	ipfsPath = fmt.Sprintf("%s/ipfs%s", Usr.HomeDir, num)
	fmt.Println("Ipfs Path:", ipfsPath)
	initCmd := exec.Command("ipfs", "init")
	initCmd.Env = append(initCmd.Env, "IPFS_PATH="+ipfsPath)
	if err := initCmd.Run(); err != nil {
		return err
	}
	return nil
}

func ReadAndCheckPorts() error {
	fmt.Println("Enter the ipfs API port Number:")
	fmt.Scan(&Port)
	ApiMaddr,err := ma.NewMultiaddr(Muladd+Port)
	ApiLis, err := manet.Listen(ApiMaddr)
	if err != nil {
		os.RemoveAll(ipfsPath)// deletes the ipfs directory that was created 
		fmt.Println("The Initialized folder has been deleted")
		log.Fatalf("serveHTTPApi: manet.Listen(%s) failed: %s", ApiMaddr, err)
		os.Exit(1)
		return err
			}

	apiPortCmd := exec.Command("ipfs", "config", "Addresses.API", Muladd+Port)
	apiPortCmd.Env = append(apiPortCmd.Env, "IPFS_PATH="+ipfsPath)
	if err := apiPortCmd.Run(); err != nil {
		return err
			}

	fmt.Println("Enter the ipfs Gateway Number:")
	var Gateway string
	fmt.Scan(&Gateway)
	gateMaddr,err := ma.NewMultiaddr(Muladd+Gateway)
	gatelis, err := manet.Listen(gateMaddr)
	if err != nil {
		os.RemoveAll(ipfsPath)
		fmt.Println("The Initialized folder has been deleted")
		log.Fatalf("serveHTTPGate: manet.Listen(%s) failed: %s", gateMaddr, err)
		fmt.Sprintf("the gateLis %s and apiListening %s : ",gatelis,ApiLis)
		os.Exit(1)
			}

	gatewayPortCmd := exec.Command("ipfs", "config", "Addresses.Gateway", Muladd+Gateway)
	gatewayPortCmd.Env = append(gatewayPortCmd.Env, "IPFS_PATH="+ipfsPath)
	if err := gatewayPortCmd.Run(); err != nil {
		log.Fatalf("failed to set the Gateway:%v",err)
			}			 
	return nil	
}
func StartDaemon() error {
	DaemonCmd := exec.Command("ipfs", "daemon")
	DaemonCmd.Env = append(DaemonCmd.Env, "IPFS_PATH="+ipfsPath)
	if err := DaemonCmd.Start(); err != nil {
		return err
		}
		fmt.Println("Daemon is running on Port:", Port)
	return nil
}


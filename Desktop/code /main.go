package main

import "log"

func main(){
	Init_err:= InitialiseIpfsDirectory() 
	if Init_err!=nil{
		log.Fatalf("failed to initialise the IPFS node: %v", Init_err)
	}

	Read_err := ReadAndCheckPorts()
	if Read_err!=nil{
		log.Fatalf("serveHTTPApi: failed: %s",Read_err)
	}
	 Start_err:=StartDaemon(); 
	 if Start_err!=nil{
		log.Fatalf("failed to start the  daemon: %v", Start_err)
	}
}
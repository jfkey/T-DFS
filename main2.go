package main

import (
	"golang-distributed-filesystem/upload"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"golang-distributed-filesystem/datanode"
	"golang-distributed-filesystem/metadatanode"
)

func main() {
	rand.Seed(time.Now().UnixNano())

//D:\cookies\windbg\1.zip
//D:\test
		debug := false;
		listener := "5050";
		dataDir := "D:/test"
		leaderAddress := "127.0.0.1:5050"
		heartbeatInterval := 3*time.Second


		conf := datanode.Config{
			DataDir:           dataDir,
			Debug:             debug,
			Listener:          getNet(listener),
			HeartbeatInterval: heartbeatInterval,
			LeaderAddress:     leaderAddress}
		datanode.Create(conf)
		// Wait on goroutines
		<-make(chan bool)


		clientListener := "5050"
		clusterListener := "5051"
		replicationFactor := 3



		conf2 := metadatanode.Config{
			getNet(clientListener),
			getNet(clusterListener),
			replicationFactor,
			"metadata.db"}
		metadatanode.Create(conf2)
		// Wait on goroutines
		<-make(chan bool)




		leaderAddress2 :=  "127.0.0.1:5050"
		filename := "D:/cookies/windbg/1.zip"

		upload.Upload(getFile(filename), debug, leaderAddress2)


}
func  getNet(ip string ) net.Listener {
	listener, err := net.Listen("tcp", ":"+ ip)
	if err != nil {
		log.Fatalln(err)
	}
	return listener
}

func getFile(filename string) *os.File {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

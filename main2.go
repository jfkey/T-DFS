package main

import (
	"golang-distributed-filesystem/datanode"
	"golang-distributed-filesystem/metadatanode"
	"golang-distributed-filesystem/upload"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	dataDir := "D:/test"
	_, _ = metadatanode.Create(metadatanode.Config{
		ClientListener:    getNet("127.0.0.1:5050"),
		ClusterListener:   getNet("127.0.0.1:5051"),
		ReplicationFactor: 2,
		DatabaseFile:      "metadata.test.db",
	})

	_, _ = datanode.Create(datanode.Config{
		Listener:         getNet("127.0.0.1:5052"),
		LeaderAddress:     "127.0.0.1:5051",
		DataDir:           dataDir + "/node1",
		HeartbeatInterval: 1 * time.Second,
	})

	_, _ = datanode.Create(datanode.Config{
		Listener:         getNet("127.0.0.1:5053"),
		LeaderAddress:     "127.0.0.1:5051",
		DataDir:           dataDir + "/node2",
		HeartbeatInterval: 1 * time.Second,
	})

	_, _ = datanode.Create(datanode.Config{
		Listener:         getNet("127.0.0.1:5054"),
		LeaderAddress:     "127.0.0.1:5051",
		DataDir:           dataDir + "/node3",
		HeartbeatInterval: 1 * time.Second,
	})

	_, _ = datanode.Create(datanode.Config{
		Listener:         getNet("127.0.0.1:5055"),
		LeaderAddress:     "127.0.0.1:5051",
		DataDir:           dataDir + "/node4",
		HeartbeatInterval: 1 * time.Second,
	})



	//D:\cookies\windbg\1.zip
	//D:\test
	debug := false;
	//listener := "127.0.0.1:5052";
	//dataDir := "D:/test"
	//leaderAddress := "127.0.0.1:5051"
	//heartbeatInterval := 3*time.Second
	//
	//
	//conf := datanode.Config{
	//	DataDir:           dataDir,
	//	Debug:             debug,
	//	Listener:          getNet(listener),
	//	HeartbeatInterval: heartbeatInterval,
	//	LeaderAddress:     leaderAddress}
	//datanode.Create(conf)
	//// Wait on goroutines
	////<-make(chan bool)
	//
	//
	//clientListener := "127.0.0.1:5050"
	//clusterListener := "127.0.0.1:5051"
	//replicationFactor := 3
	//
	//
	//
	//conf2 := metadatanode.Config{
	//	getNet(clientListener),
	//	getNet(clusterListener),
	//	replicationFactor,
	//	"metadata.db"}
	//metadatanode.Create(conf2)
	// Wait on goroutines
	//<-make(chan bool)
	//
	//
	//
	//
	leaderAddress2 :=  "127.0.0.1:5050"
	filename := "D:/res/add.csv"

	upload.Upload(getFile(filename), debug, leaderAddress2)


}
func  getNet(ip string ) net.Listener {
	listener, err := net.Listen("tcp", ip)
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
package main

import ("fmt"
	"log"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"time"
)

var (
	device		string = "ens32"
	snapshot_len	int32 = 1024
	promiscuous	bool = false
	err		error
	timeout		time.Duration = 3 * time.Second
	handle		*pcap.Handle
)

func main() {
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {log.Fatal(err)}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
		break
	}

}

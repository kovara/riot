package main

import ("flag"
	"fmt"
	"log"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"time"
)

var (
	device		string = "ens32"
	snapshot_len	int32 = 1024
	promiscuous	bool = true
	err		error
	timeout		time.Duration = -1 * time.Second
	handle		*pcap.Handle
)

func printPacketInfo(packet gopacket.Packet) {
}

func main() {
	clientPtr := flag.String("client", "", "IP of client")
	serverPtr := flag.String("server", "", "IP of server")
	arpPtr := flag.Bool("noarp",false,"Do not Arpspoof")
	flag.Parse()
	if len(*clientPtr) == 0 {log.Fatal("no client IP address entered")}
	if len(*serverPtr) == 0 {log.Fatal("no server IP address entered")}
	if *arpPtr == true {log.Fatal("arp flag")}

	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {log.Fatal(err)}
	err = handle.SetBPFFilter("host "+*clientPtr+" and host "+*serverPtr)
	if err != nil {log.Fatal(err)}
	defer handle.Close()

	fmt.Println("capturing...")
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}

}

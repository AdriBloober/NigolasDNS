package main

import (
	"NigolasDNS/database"
	"NigolasDNS/dns"
	"fmt"
	"log"
	"net"
)

const (
	PORT = "1053"
	MAXIMAL_DNS_BYTES = 512
)

func main() {
	database.Connect()
	serverAddr, err := net.ResolveUDPAddr("udp", ":" + PORT)
	if err != nil {log.Fatal(err)}
	serverConn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {log.Fatal(err)}
	fmt.Println("UDP DNS Server listening on Port", PORT)
	defer serverConn.Close()
	request := make([]byte, MAXIMAL_DNS_BYTES)
	for {
		_, clientAddr, err := serverConn.ReadFromUDP(request)
		if err != nil {
			fmt.Println("An error on client", clientAddr.IP.String(), "is occured:", err)
		} else {
			go dns.HandleDNSRequest(request, serverConn, clientAddr)
		}
	}
}
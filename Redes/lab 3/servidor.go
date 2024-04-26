package main

import (
	"fmt"
	"net"
	"os"
	// "strings"
	"encoding/json"
)

type DNSRecord struct {
	Option string `json:"option"`
	Name   string `json:"domainName"`
	IP     string `json:"ip"`
	TTL    string `json:"ttl"`
	Type   string `json:"type"`
}

var dnsRecords []DNSRecord

func handleDNSRequest(conn *net.UDPConn) {
	buffer := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error reading from UDP connection:", err)
		return
	}

	fmt.Println("Received DNS request from", addr)

	var record DNSRecord

	err = json.Unmarshal(buffer[:n], &record)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	nombre := ""
	fmt.Println("Received DNS record:", record)
	if record.Option == "1" {
		dnsRecords = append(dnsRecords, record)
		fmt.Println("Append realizado")
	} else if record.Option == "2" {
		for _, record2 := range dnsRecords {
			if record2.Name == record.Name {
				nombre = record.Name
				break
			} else {
				nombre = "no existe"
			}
		}
		response := []byte(nombre)
		_, err = conn.WriteToUDP(response, addr)
		if err != nil {
			fmt.Println("error")
		}
	}
	// fmt.Println(nombre)
}

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":63420")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error listening on UDP port 53:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("DNS server listening on port 63420...")

	for {
		handleDNSRequest(conn)
	}
}

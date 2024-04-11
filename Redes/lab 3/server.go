package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"encoding/json"
)

type DNSRecord struct {
    Name  string `json:"name"`
    IP    string `json:"ip"`
    TTL   int    `json:"ttl"`
    Type  string `json:"type"`
}


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

    // Aquí puedes guardar la información del DNS en una base de datos,
    // en memoria o en cualquier otro lugar según tus necesidades.

    fmt.Println("Received DNS record:", record)
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

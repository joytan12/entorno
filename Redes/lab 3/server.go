package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func handleDNSRequest(conn *net.UDPConn) {
	buffer := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error reading from UDP connection:", err)
		return
	}

	fmt.Println("Received DNS request from", addr)

	domain := strings.TrimSpace(string(buffer[:n]))

	ip, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println("Error looking up IP for", domain, ":", err)
		return
	}

	fmt.Println("Resolved", domain, "to", ip[0])

	response := []byte(ip[0].String())
	_, err = conn.WriteToUDP(response, addr)
	if err != nil {
		fmt.Println("Error writing response to UDP connection:", err)
		return
	}
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

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type GameBoard struct {
	Board [][]int `json:"board"`
}

type ServerResponse struct {
	Available bool   `json:"Available"`
	TCP_IP    string `json:"TCP_IP"`
	TCP_Port  string `json:"TCP_Port"`
	Board     [][]int
}

type ServerState struct {
	Continue bool    `json:"Continue"`
	Board    [][]int `json:"Board"`
	Winner   string  `json:"Winner"`
}

func main() {
	// Dirección para conexión con el servidor UDP
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("Error con la dirección del servidor:", err)
		return
	}

	// Establecer conexión UDP con el Servidor
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error al crear la conexión:", err)
		return
	}
	defer conn.Close()

	// Ya con la conexión, debemos enviar la orden de inicio del juego
	message := []byte("Empezar juego")

	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error al enviar datos:", err)
		return
	}

	fmt.Println("Datos enviados correctamente al servidor UDP.")

	// Se espera la disponibilidad del servidor, junto a su ip, puerto y el tablero
	buffer := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error al recibir respuesta:", err)
		return
	}

	// Decodificar JSON
	var serverResponse ServerResponse
	if err := json.Unmarshal(buffer[:n], &serverResponse); err != nil {
		fmt.Println("Error al decodificar la respuesta del servidor:", err)
		return
	}

	fmt.Printf("Respuesta del servidor (%s): %s\n", addr, string(buffer[:n]))

	// Utilizar las variables de la respuesta para conectarse al servidor TCP
	fmt.Println("Disponibilidad del servidor:", serverResponse.Available)
	fmt.Println("Dirección IP del servidor TCP:", serverResponse.TCP_IP)
	fmt.Println("Puerto del servidor TCP:", serverResponse.TCP_Port)
	printBoard(serverResponse.Board)

	// Ciclo de juego
	for {
		// Conectar al servidor TCP
		connTCP, err := net.Dial("tcp", serverResponse.TCP_IP+":"+serverResponse.TCP_Port)
		if err != nil {
			fmt.Println("Error al conectar al servidor TCP:", err)
			return
		}
		defer connTCP.Close()

		fmt.Println("Conexión establecida con el servidor TCP.")

		// Generar jugada y luego enviarla
		fmt.Print("Ingresa un dígito: ")
		reader := bufio.NewReader(os.Stdin)
		digitStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer el dígito:", err)
			return
		}

		// Enviar el dígito al servidor
		_, err = fmt.Fprintf(connTCP, digitStr)
		if err != nil {
			fmt.Println("Error al enviar el dígito al servidor:", err)
			return
		}

		//Esperar respuesta del servidor TCP
		var serverState ServerState
		decoder := json.NewDecoder(connTCP)
		if err := decoder.Decode(&serverState); err != nil {
			fmt.Println("Error al decodificar los datos JSON del servidor:", err)
			return
		}
		fmt.Println("Estado del servidor:", serverState)
		serverResponse.Board = serverState.Board
		fmt.Println("Tablero cliente actual:")
		printBoard(serverResponse.Board)

		if serverState.Continue == false { //Juego finaliza
			fmt.Printf("Juego finalizado. ¡Ganador %s!\n", serverState.Winner)
			break
		}
		//  else { //Juego continúa
		// 	if err := decoder.Decode(&serverState); err != nil {
		// 		fmt.Println("Error al decodificar los datos JSON del servidor:", err)
		// 		return
		// 	}
		// 	fmt.Println("Estado del servidorr:", serverState)
		// }

	}
}

// Función para imprimir el tablero
func printBoard(board [][]int) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}
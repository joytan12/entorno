package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type GameBoard struct {
	Board [][]int `json:"board"`
}

func main() {
	// Establecer la dirección del servidor
	serverAddr := "localhost:8080"

	// Conectar al servidor
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Error al conectar con el servidor: %v", err)
	}
	defer conn.Close()

	// Enviar un mensaje al servidor para iniciar el juego (puedes enviar cualquier cosa que el servidor espere)
	message := "start"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatalf("Error al enviar mensaje al servidor: %v", err)
	}

	// Recibir el tablero del servidor
	var gameBoard GameBoard
	decoder := json.NewDecoder(conn)
	if err := decoder.Decode(&gameBoard); err != nil {
		log.Fatalf("Error al decodificar el tablero: %v", err)
	}

	// Imprimir el tablero recibido
	fmt.Println("Tablero recibido del servidor:")
	printBoard(gameBoard.Board)
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

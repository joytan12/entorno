package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

type GameBoard struct {
	Board [][]int `json:"board"`
}

func main() {
	// Establecer la semilla del generador de números aleatorios
	rand.Seed(time.Now().UnixNano())

	// Generar los tableros
	board1 := generateBoard()
	board2 := generateBoard()

	// Colocar dos unos al azar en cada tablero
	placeRandomOnes(board1)
	placeRandomOnes(board2)

	// Iniciar el servidor en el puerto 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
	defer listener.Close()
	fmt.Println("Servidor en ejecución. Esperando conexión...")

	// Aceptar la conexión del cliente
	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("Error al aceptar la conexión del cliente: %v", err)
	}
	defer conn.Close()
	fmt.Println("Cliente conectado.")

	// Esperar el mensaje del cliente para comenzar el juego
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalf("Error al leer del cliente: %v", err)
	}
	fmt.Printf("Mensaje del cliente: %s\n", buf[:n])

	// Decodificar los datos recibidos del cliente
	var attackCoordinates []int
	if err := json.Unmarshal(buf[:n], &attackCoordinates); err != nil {
		log.Fatalf("Error al decodificar los datos del cliente: %v", err)
	}

	// Aplicar el ataque en el tablero del oponente
	attack(board2, attackCoordinates)

	// Enviar el tablero al cliente
	sendBoard(conn, board2)

	fmt.Println("Juego comenzado. Tablero actualizado y enviado al cliente.")
}

// Generar un tablero 2x2 con ceros
func generateBoard() [][]int {
	board := make([][]int, 2)
	for i := range board {
		board[i] = make([]int, 2)
	}
	return board
}

// Colocar dos unos al azar en el tablero
func placeRandomOnes(board [][]int) {
	for i := 0; i < 2; i++ {
		row := rand.Intn(2)
		col := rand.Intn(2)
		board[row][col] = 1
	}
}

// Aplicar el ataque en el tablero del oponente
func attack(board [][]int, coordinates []int) {
	row, col := coordinates[0], coordinates[1]
	board[row][col] = 0
}

// Enviar el tablero al cliente
func sendBoard(conn net.Conn, board [][]int) {
	gameBoard := GameBoard{Board: board}
	jsonData, err := json.Marshal(gameBoard)
	if err != nil {
		log.Fatalf("Error al serializar el tablero: %v", err)
	}
	conn.Write(jsonData)
}

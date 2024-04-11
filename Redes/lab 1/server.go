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

type ServerInfo struct {
	Available bool    // Indica si el servidor está disponible
	TCP_IP    string  // Dirección IP del servidor TCP
	TCP_Port  string  // Dirección Port del servidor TCP
	Board     [][]int // Tablero 1
}

type State struct {
	Continue bool    // Indica si se sigue jugando o no
	Board    [][]int // Tablero
	Winner   string  // Indica el ganador
}

func main() {
	// Establecer la semilla del generador de números aleatorios
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generar los tableros
	board1 := generateBoard()
	board2 := generateBoard()
	fmt.Println("Tableros generados")

	// Colocar dos unos al azar en cada tablero
	placeRandomOnes(board1)
	placeRandomOnes(board2)
	fmt.Println("Barcos colocados")

	fmt.Println("Tablero servidor:")
	printBoard(board1)
	fmt.Println("\nTablero cliente:")
	printBoard(board2)

	// Creación de conexión UDP
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("Error en la creación UDP:", err)
		return
	}

	// UDP en modo "escuchar"
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error al escuchar en la dirección UDP:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Servidor UDP iniciado.")

	// buffer para leer los datos que recibiremos desde el cliente
	buffer := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error al leer datos:", err)
	}

	respuesta := string(buffer[:n])
	// Imprimir los datos recibidos
	fmt.Printf("Recibido %d bytes de %s: %s\n", n, addr, respuesta)
	if respuesta == "Empezar juego" {
		fmt.Println("¡El cliente dio la orden de iniciar el juego!")
	}

	// Enviar al cliente los datos para iniciar el juego
	// (Disponibilidad, IP, Puerto [TCP] y tablero)
	serverInfo := ServerInfo{
		Available: true,        // Disponibilidad del servidor
		TCP_IP:    "127.0.0.1", // Dirección IP del servidor TCP
		TCP_Port:  "9000",      //Puerto del servidor TCP
		Board:     board2,      // Enviar el tablero 2 (cliente)
	}

	// Pasar esta información a JSON para facilidad de envío y recibimiento
	serverInfoJSON, err := json.Marshal(serverInfo)
	if err != nil {
		fmt.Println("Error al serializar ServerInfo:", err)
		return
	}

	// Abrir servidor TCP utilizando la información del JSON
	tcpAddr := serverInfo.TCP_IP + ":" + serverInfo.TCP_Port
	listener, err := net.Listen("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Error al abrir servidor TCP:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor TCP iniciado en", tcpAddr)

	// Enviar ServerInfo serializado al cliente
	_, err = conn.WriteToUDP(serverInfoJSON, addr)
	if err != nil {
		fmt.Println("Error al enviar datos al cliente:", err)
		return
	}

	fmt.Println("Juego comenzado. Tablero enviado al cliente.")

	// Bucle principal del juego
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error al aceptar conexión entrante:", err)
			continue
		}
		defer conn.Close()

		// Aquí puedes manejar la conexión entrante
		fmt.Println("Conexión entrante aceptada desde:", conn.RemoteAddr())

		// Leer los datos enviados por el cliente
		buffer := make([]byte, 1024)
		_, err = conn.Read(buffer)
		if err != nil {
			fmt.Println("Error al leer datos del cliente:", err)
			conn.Close() // Cerramos la conexión si hay un error al leer los datos
			continue
		}

		// Estado del juego
		serverState := State{
			Continue: true,   // Seguir juego o no
			Board:    board2, // Enviar el tablero2 (cliente)
			Winner:   "-1",   // Indicar ganador
		}

		// Pasar esta información a JSON para facilidad de envío y recibimiento
		serverStateJSON, err := json.Marshal(serverState)
		if err != nil {
			fmt.Println("Error al serializar ServerInfo:", err)
			return
		}

		// Imprimir la jugada recibida (Puede ser A, B, C o D)
		jugada := string(buffer[0])
		jugada_cords := letterToCoordinates(jugada)
		fmt.Printf("Jugada recibida del cliente (%s): %s, con coordenadas: %v\n", conn.RemoteAddr(), jugada, jugada_cords)

		//Realizar jugada Cliente
		attack(board1, jugada_cords)
		// Realizar la jugada y verificar si alguien ganó.
		fmt.Println("Tablero servidor:")
		printBoard(board1)
		fmt.Println("\nTablero cliente:")
		printBoard(board2)

		//Verificar si se continúa o no
		if checkForShips(board1) == true { //Si no hay barcos en el board Servidor
			// fmt.Println("-----------------------------")
			// printBoard(board1)
			// fmt.Println("-----------------------------")
			serverState.Continue = false
			serverState.Winner = "Cliente"
			serverStateJSON, err := json.Marshal(serverState)
			if err != nil {
				fmt.Println("Error al serializar ServerInfo:", err)
				return
			}
			_, err = conn.Write(serverStateJSON)
			break
		}

		// Si quedan barcos en servidor el juego sigue, y continua con coordenadas random
		fmt.Println("Ahora juega el servidor")
		coordenadas_random := generarCoordsRandom()
		fmt.Println("coords random", coordenadas_random)
		//Ataque del servidor
		attack(board2, coordenadas_random) // Jugada servidor
		fmt.Println("Tablero servidor:")
		printBoard(board1)
		fmt.Println("\nTablero cliente:")
		printBoard(board2)

		//Verificar si se continua o no
		if checkForShips(board2) == true { //Si no hay barcos en el board Cliente
			// fmt.Println("-----------------------------")
			// printBoard(board2)
			// fmt.Println("-----------------------------")
			serverState.Continue = false
			serverState.Winner = "Servidor"
			// _, err = conn.Write(serverStateJSON)
		}

		serverStateJSON, err = json.Marshal(serverState)
		if err != nil {
			fmt.Println("Error al serializar ServerInfo:", err)
			return
		}

		_, err = conn.Write(serverStateJSON)

		// fmt.Println("Tablero servidor:")
		// printBoard(board1)
		// fmt.Println("\nTablero cliente:")
		// printBoard(board2)

		// serverState.Continue = (checkForShips(board1) || checkForShips(board2))
		// if checkForShips(board1) {
		// 	serverState.Winner = "Servidor"
		// }
		// if checkForShips(board2) {
		// 	serverState.Winner = "Cliente"
		// }

		// _, err = conn.Write(serverStateJSON)

		if serverState.Continue == false {
			// Si le atinó el servidor, el juego termina con previa señal al cliente.
			fmt.Printf("Juego finalizado. ¡Ganador: %s!\n", serverState.Winner)
			break
		}
		fmt.Println("estado final", serverState)
	}
	// Juego finalizado debido al continue, ya sea de los errores o de que terminó el game
	// defer listener.Close()
	fmt.Println("Conexión TCP finalizada")
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
	// Generar dos pares de coordenadas únicas
	for i := 0; i < 2; i++ {
		row := rand.Intn(2)
		col := rand.Intn(2)
		// Verificar si la posición ya está ocupada por un uno
		for board[row][col] == 1 {
			row = rand.Intn(2)
			col = rand.Intn(2)
		}
		board[row][col] = 1
	}
}

// Aplicar el ataque en el tablero del oponente
func attack(board [][]int, coordinates []int) bool {
	row, col := coordinates[0], coordinates[1]

	if board[row][col] == 1 {
		board[row][col] = -1
		return true
	}
	board[row][col] = -1
	return false
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

func letterToCoordinates(letter string) []int {
	var row, col int

	switch letter {
	case "A":
		row, col = 0, 0
	case "B":
		row, col = 0, 1
	case "C":
		row, col = 1, 0
	case "D":
		row, col = 1, 1
	default:
		fmt.Println("Letra no válida")
		return nil
	}

	// fmt.Println(row, col)
	return []int{row, col}
}

func printBoard(board [][]int) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}

func generarCoordsRandom() []int {
	// rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano())) // Inicializar la semilla del generador de números aleatorios

	// Generar coordenadas aleatorias
	coordenadas := make([]int, 2)
	for i := range coordenadas {
		coordenadas[i] = rand.Intn(2) // Generar un número aleatorio entre 0 y 1
	}
	return coordenadas
}

func checkForShips(board [][]int) bool {
	for _, row := range board {
		for _, cell := range row {
			if cell == 1 {
				return false // Hay al menos un barco en el tablero
			}
		}
	}
	return true // No hay barcos en el tablero
}
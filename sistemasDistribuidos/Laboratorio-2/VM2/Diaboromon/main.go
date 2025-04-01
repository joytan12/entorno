package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	pb "Diaboromon/proto" // Asegúrate de que sea el path correcto generado a partir del .proto

	"google.golang.org/grpc"
)

const (
	port    = ":8087" // Puerto donde escucha
	nodetai = "nodetai:8086"
	daño    = "10"
)

type Input struct {
	// TODO: Describir en el readme.txt que se tendrá en cuenta que solo ps puede tener decimales
	PS float64
	TE int
	TD int
	CD int
	VI int
}

var wg sync.WaitGroup // Definir un WaitGroup
var cd int

// Implementación del servidor gRPC
type server struct {
	pb.UnimplementedDigimonServiceServer
}

// Implementación del método GetDigimonStatus
func (s *server) GetDigimonStatus(ctx context.Context, req *pb.DigimonRequest) (*pb.DigimonResponse, error) {
	if req.Opt == "1" {
		fmt.Println("Me contactó Tai")
		// Devolverle la cantidad de datos que necesita para evolucionar y matarme noOoOo
		cantDatos := strconv.Itoa(cd)
		return &pb.DigimonResponse{Status: cantDatos}, nil
	} else if req.Opt == "2" {
		fmt.Println("[Diaboromon] Así que perdí, pues nada, GGWP! 💋")
		os.Exit(0)
	} else {
		fmt.Println("[Diaboromon] GANÉÉÉ GG!! 😈😈😈")
		os.Exit(0)
	}
	// Enviar confirmación al cliente
	return &pb.DigimonResponse{Status: "respuesta default:o"}, nil
}

func leerInput(archivo string) (Input, error) {
	var config Input
	file, err := os.Open(archivo)
	if err != nil {
		return config, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		// Leer la línea y separar por comas
		valores := strings.Split(scanner.Text(), ",")
		if len(valores) != 5 {
			return config, fmt.Errorf("formato de input incorrecto")
		}

		fmt.Sscanf(valores[0], "%f", &config.PS)
		fmt.Sscanf(valores[1], "%d", &config.TE)
		fmt.Sscanf(valores[2], "%d", &config.TD)
		fmt.Sscanf(valores[3], "%d", &config.CD)
		fmt.Sscanf(valores[4], "%d", &config.VI)
	}

	if err := scanner.Err(); err != nil {
		return config, err
	}
	return config, nil
}

func ataquePeriodico(TD int) {
	defer wg.Done()
	intervalo := time.Duration(TD) * time.Second
	// fmt.Printf("Esperando %.0f segundos\n", intervalo.Seconds())
	fmt.Println("Iniciando goroutine para atacar al hdp de tai muajajaj 😎")
	for {
		// Conectarse al servidor gRPC
		fmt.Printf("Me intentaré conectar a %s, de forma malvada xq soy diaboromon 😈\n", nodetai)
		conn, err := grpc.Dial(nodetai, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Error al conectar al servidor: %v", err)
		}
		defer conn.Close()
		fmt.Printf("Esperando %.0f segundos para atacar al grupo de Tai.\n", intervalo.Seconds())
		time.Sleep(intervalo)
		opt := "1"
		status, err := EnviarConsulta(conn, opt, daño)
		if err != nil {
			log.Fatalf("Error al enviar la consulta: %v\n", err)
		}
		fmt.Printf("Me respondió esto nodetai: %s\n", status)

		if status == "0" {
			// Significa que triunfó el mal GG
			fmt.Println("Ganó Diaboromon, yiyi papa")
			os.Exit(0)
		}
	}
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano())) // Seed "aleatoria"

	// Leer archivo input.txt
	// Valores quedarán en las variables input.PS, input.TA, etc. EJ: fmt.Printf("Probabilidad de sacrificio: %f", input.PS)
	input, err := leerInput("input.txt")
	if err != nil {
		fmt.Println("Error al leer input:", err)
		return
	}

	cd = input.CD

	// // Conectarse al servidor gRPC
	// fmt.Printf("Me intentaré conectar a %s, de forma malvada xq soy diaboromon 😈\n", nodetai)
	// conn, err := grpc.Dial(nodetai, grpc.WithInsecure(), grpc.WithBlock())
	// if err != nil {
	// 	log.Fatalf("Error al conectar al servidor: %v", err)
	// }
	// defer conn.Close()

	// Crear el listener en el puerto definido
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error al iniciar el listener: %v", err)
	}

	// Crear un servidor gRPC
	grpcServer := grpc.NewServer()
	pb.RegisterDigimonServiceServer(grpcServer, &server{})

	log.Printf("Diaboromon escuchando en el puerto %s", port)

	// Inicia la goroutine para atacar a nodetai.
	wg.Add(1)

	// input.TD = 1 // TODO: Quitar esa linea jdskjds

	// time.Sleep(1 * time.Second)
	go ataquePeriodico(input.TD)

	// Iniciar el servidor gRPC
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	wg.Wait()

	fmt.Println("Programa finalizado gg")
}

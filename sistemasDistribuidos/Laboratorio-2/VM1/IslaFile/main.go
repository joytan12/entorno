package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	pb "IslaFile/proto" // Asegúrate de que sea el path correcto generado a partir del .proto

	"google.golang.org/grpc"
)

const (
	port        = ":8084"            // Puerto donde escucha el Data Node
	primarynode = "primarynode:8080" // IP que se conectará por gRPC
	n           = 6                  // Cantidad de Digimons iniciales
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

// Implementación del servidor gRPC
type server struct {
	pb.UnimplementedDigimonServiceServer
}

// Implementación del método GetDigimonStatus
func (s *server) GetDigimonStatus(ctx context.Context, req *pb.DigimonRequest) (*pb.DigimonResponse, error) {
	if req.Opt == "-1" {
		fmt.Println("Apagando IslaFile")
		os.Exit(0)
	}
	// Enviar confirmación al cliente
	respuesta := fmt.Sprintf("Información de %s almacenada correctamente", req.Todo)
	return &pb.DigimonResponse{Status: respuesta}, nil
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

func leerDigimons(archivo string) ([]string, error) {
	var digimons []string
	file, err := os.Open(archivo)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		digimons = append(digimons, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return digimons, nil
}

func seleccionarDigimonAleatorio(digimons []string, ps float64) string {
	// Seleccionar un Digimon al azar
	numero := rand.Intn(len(digimons))
	// fmt.Printf("Número aleatorio: %d, digimon seleccionado: %s\n", numero, digimons[numero])
	digimonAleatorio := digimons[numero]
	atributos := strings.Split(digimonAleatorio, ",")

	// Generar un número aleatorio entre 0 y 1
	salidaRandom := rand.Float64()
	var estado string

	// Determinar si será sacrificado, si es menor
	// fmt.Printf("Porcentaje sacrificio random: %f, ps: %f\n", salidaRandom, ps)
	if salidaRandom < ps {
		estado = "Sacrificado"
	} else {
		estado = "No-Sacrificado"
	}

	// Crear el string en el formato solicitado
	resultado := fmt.Sprintf("%s,%s,%s", atributos[0], atributos[1], estado)

	return resultado
}

func enviarIniciales(conn *grpc.ClientConn, opt string, digimons []string, ps float64, n int) {
	for i := 0; i < n; i++ {
		resultado := seleccionarDigimonAleatorio(digimons, ps)
		status, err := EnviarConsulta(conn, opt, resultado)
		if err != nil {
			log.Fatalf("Error al hacer la consulta: %v", err)
		}
		fmt.Printf("Me respondio esto el server: %s\n", status)
	}
}

func enviarPeriodicamente(conn *grpc.ClientConn, opt string, digimons []string, ps float64, n int, te int) {
	defer wg.Done()
	intervalo := time.Duration(te) * time.Second
	fmt.Println("Iniciando goroutine encargada de enviar digimon periódicamente.")
	for {
		fmt.Printf("Esperando %.0f segundos para enviar un digimon al Primary Node.\n", intervalo.Seconds())
		time.Sleep(intervalo) // Ejecutar cada input.TE [unidad de tiempo].
		resultado := seleccionarDigimonAleatorio(digimons, ps)
		status, err := EnviarConsulta(conn, opt, resultado)
		if err != nil {
			log.Fatalf("Error al enviar la consulta: %v\n", err)
		}
		fmt.Printf("Me respondio esto el server: %s\n", status)

		if status == "-1" {
			fmt.Println("Respuesta '-1' recibida, terminando la goroutine y el programa. 10-4")
			return
		}
	}
}

func main() {
	disponible(primarynode, 5)
	rand.New(rand.NewSource(time.Now().UnixNano())) // Seed "aleatoria"

	// Leer archivo input.txt
	// Valores quedarán en las variables input.PS, input.TA, etc. EJ: fmt.Printf("Probabilidad de sacrificio: %f", input.PS)
	input, err := leerInput("input.txt")
	if err != nil {
		fmt.Println("Error al leer input:", err)
		return
	}

	// Leer archivo digimons.txt
	digimons, err := leerDigimons("digimons.txt")
	if err != nil {
		fmt.Println("Error al leer archivo digimons.txt:", err)
		return
	}

	// Digimon de ejemplo
	// resultado := seleccionarDigimonAleatorio(digimons, input.PS)
	// fmt.Println(resultado)

	// Conectarse al servidor gRPC
	fmt.Println("Me intentaré conectar sdjksdkj")
	conn, err := grpc.Dial(primarynode, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error al conectar al servidor: %v", err)
	}
	defer conn.Close()

	// Enviar los 6 digimons iniciales
	opt := "1"
	enviarIniciales(conn, opt, digimons, input.PS, n)

	// Crear el listener en el puerto definido
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error al iniciar el listener: %v", err)
	}

	// Crear un servidor gRPC
	grpcServer := grpc.NewServer()
	pb.RegisterDigimonServiceServer(grpcServer, &server{})

	log.Printf("Continente Server escuchando en el puerto %s", port)

	// Inicia la goroutine para enviar periódicamente los digimons.
	wg.Add(1)

	go enviarPeriodicamente(conn, opt, digimons, input.PS, 6, input.TE)

	// Iniciar el servidor gRPC
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	wg.Wait()
	fmt.Println("Programa finalizado gg")
}

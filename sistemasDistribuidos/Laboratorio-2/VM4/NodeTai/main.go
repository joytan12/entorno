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

	pb "NodeTai/proto" // Asegúrate de que sea el path correcto generado a partir del .proto

	"google.golang.org/grpc"
)

const (
	port        = ":8086" // Puerto donde escucha el Data Node
	primarynode = "primarynode:8080"
	diaboromon  = "diaboromon:8087"
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
var vida int
var estado int = -1

// Implementación del servidor gRPC
type server struct {
	pb.UnimplementedDigimonServiceServer
}

// Implementación del método GetDigimonStatus
func (s *server) GetDigimonStatus(ctx context.Context, req *pb.DigimonRequest) (*pb.DigimonResponse, error) {
	// Si recibe un 1 es que le atacó diaboromon, si recibe un -1 se acaba todo D:
	if req.Opt == "1" {
		fmt.Printf("Recibí un ataque de Diaboromon con valor de: %s, disminuyendo vida noOoOoOo\n", req.Todo)

		// Disminuir vida
		valorAtaque, err := strconv.Atoi(req.Todo)
		if err != nil {
			fmt.Println("Error al convertir el valor de ataque:", err)
		}

		vida -= valorAtaque
		// Verificar si la vida es 0 o menor
		if vida <= 0 {
			fmt.Println("Finalizó el juego porque nos quedamos sin vida.")
			verificarYAtacar()
			// TODO: Hacer que en el goroutine se verifique xq aquí no c como se hace
			// Terminar el programa
			// os.Exit(0)
		} else {
			fmt.Printf("Vida actual: %d\n", vida)
		}
	} else if req.Opt == "-1" {
		fmt.Printf("Aquí acaba el programa\n")
	}

	return &pb.DigimonResponse{Status: strconv.Itoa(vida)}, nil
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

func verificarYAtacar() {
	defer wg.Done()
	// fmt.Println("Iniciando Goroutine para verificar primarynode owo")
	// cantDatos := strconv.Atoi(cd)
	// Función periódica encargada de preguntarle al primarynode sobre cuántos sacrificados hay 😨.
	for {
		if vida <= 0 {
			fmt.Println("Perdimos:(, ganó Diaboromon, cerrando el programa.")
			// Contactar a diaboromon informando que ganó ff
			// Verificar conectándose a Diaboromon cuántos datos necesito:
			conn2, err2 := grpc.Dial(diaboromon, grpc.WithInsecure(), grpc.WithBlock())
			if err2 != nil {
				log.Fatalf("Error al conectar al servidor: %v", err2)
			}
			defer conn2.Close()

			opt := "3"
			EnviarConsulta(conn2, opt, "Ganaste wey")
			// if err3 != nil {
			// 	log.Fatalf("Error al enviar la consulta a Diaboromon: %v\n", err3)
			// }

			// Avisar a primarynode que gano diaboromon
			conn, err := grpc.Dial(primarynode, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Fatalf("Error al conectar al servidor: %v", err)
			}
			EnviarConsulta(conn, "3", "Ganó Tai yiyi")
			os.Exit(0)
		}
		if estado == 1 {
			// Contactar a primarynode para que finalice todo acto
			fmt.Println("GANAMOOOOS 😜🎉, cerrando el programa.")
			conn, err := grpc.Dial(primarynode, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Fatalf("Error al conectar al servidor: %v", err)
			}
			defer conn.Close()
			if vida <= 0 {
				EnviarConsulta(conn, "3", "Perdió Tai ff")
			} else {
				EnviarConsulta(conn, "3", "Ganó Tai yiyi")
			}
			os.Exit(0)
		}
		// Solicitar opción al usuario
		fmt.Println("Selecciona una opción:")
		fmt.Println("1 - Lanzar ataque a Diaboromon")
		fmt.Println("2 - Salir")

		// Leer opción desde la entrada estándar
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Ingresa tu opción: ")
		opcion, _ := reader.ReadString('\n')
		opcion = strings.TrimSpace(opcion) // Eliminar saltos de línea y espacios

		if vida <= 0 {
			continue
		}

		if opcion == "1" {
			// Conectarse al servidor gRPC del primarynode
			conn, err := grpc.Dial(primarynode, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Fatalf("Error al conectar al servidor: %v", err)
			}
			defer conn.Close()

			// Enviar consulta con opt = "2"
			opt := "2"
			status, err := EnviarConsulta(conn, opt, "Solicitud")
			if err != nil {
				log.Fatalf("Error al enviar la consulta: %v\n", err)
			}

			// Mostrar respuesta del primarynode
			fmt.Printf("Me respondió esto el primarynode: %s\n", status) // TODO QUITAR ESTO

			// Verificar conectándose a Diaboromon cuántos datos necesito:
			conn2, err2 := grpc.Dial(diaboromon, grpc.WithInsecure(), grpc.WithBlock())
			if err2 != nil {
				log.Fatalf("Error al conectar al servidor: %v", err2)
			}
			defer conn2.Close()

			opt = "1"
			rest, err2 := EnviarConsulta(conn2, opt, "Ad")
			if err2 != nil {
				log.Fatalf("Error al enviar la consulta a Diaboromon: %v\n", err2)
			}

			fmt.Printf("Me respondió esto diaboromon: %s\n", rest) // TODO QUITAR ESTO

			// Convertir status a float para comprobar si se puede evolucionar
			cd, _ := strconv.Atoi(rest)
			statusFloat, _ := strconv.ParseFloat(status, 64)
			if statusFloat >= float64(cd) {
				fmt.Printf("Ya puedo evolucionar pa pitearme al Diaboromon, porque tengo %.2f datos y necesitaba %d para evolucionar yiyi\n", statusFloat, cd)
				fmt.Printf("Contactando a Diaboromon para avisarle que perdió 🗣️🗣️🗣️\n")
				opt = "2"
				EnviarConsulta(conn2, opt, "TE GANAMOS, y la keso 💋")
				estado = 1
				// if err2 != nil {
				// 	log.Fatalf("Error al enviar la consulta a Diaboromon: %v\n", err2)
				// }
				// fmt.Printf("Me respondió esto diaboromon: %s\n", rest2) // TODO QUITAR ESTO
			} else {
				vida -= 10
				fmt.Printf("Aún no puedo evolucionar D:, Diaboromon me repele el ataque, quitándome 10 puntos de vida, vida actual: %d\n", vida)
			}
		} else if opcion == "2" {
			// Salir del bucle y terminar el programa
			fmt.Println("Saliendo...")
			break
		} else {
			fmt.Println("Opción no válida, por favor ingresa 1 o 2.")
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

	fmt.Printf("Vida inicial Greymon y Garurumon: %d, Cantidad para evolucionar: %d\n", input.VI, input.CD)
	vida = input.VI

	// Crear el listener en el puerto definido
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error al iniciar el listener: %v", err)
	}

	// Crear un servidor gRPC
	grpcServer := grpc.NewServer()
	pb.RegisterDigimonServiceServer(grpcServer, &server{})

	log.Printf("Node Tai escuchando en el puerto %s", port)

	wg.Add(1)

	go verificarYAtacar()

	// Iniciar el servidor gRPC
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	wg.Wait()

	fmt.Println("Programa finalizado gg")

}

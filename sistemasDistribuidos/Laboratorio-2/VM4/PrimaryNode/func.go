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
	"time"

	pb "PrimaryNode/proto"

	"google.golang.org/grpc"
)

// ---------------------------------------tratado de datos, formatos y logica----------------------------------

// GenerarID genera un ID único utilizando un timestamp y un valor aleatorio.
func GenerarID() string {
	// Establecer la semilla aleatoria basada en el tiempo actual
	rand.Seed(time.Now().UnixNano())

	// Obtener el timestamp actual en nanosegundos
	timestamp := time.Now().UnixNano()

	// Generar un número aleatorio adicional para asegurar la unicidad
	randomPart := rand.Intn(10000)

	// Combinar el timestamp con el número aleatorio para crear un ID único
	id := fmt.Sprintf("%d%04d", timestamp, randomPart)

	return id
}

// SeleccionarDataNode selecciona un Data Node según la inicial del nombre del Digimon.
func SeleccionarDataNode(nombreDigimon string) int {
	// Asegurarse de que el nombre no esté vacío.
	if nombreDigimon == "" {
		fmt.Println("El nombre del Digimon no puede estar vacío")
		return -1
	}

	// Convertir la inicial a mayúsculas para facilitar la comparación.
	inicial := strings.ToUpper(string(nombreDigimon[0]))

	// Comprobar si la inicial está entre A-M o N-Z.
	if inicial >= "A" && inicial <= "M" {
		return 0 // Data Node 1
	} else if inicial >= "N" && inicial <= "Z" {
		return 1 // Data Node 2
	} else {
		fmt.Println("Inicial no válida")
		return -1
	}
}

// FormatearInformaciónDigimon toma la información del Digimon, ID y número del Data Node
// y devuelve dos cadenas con los formatos solicitados.
func FormatearInformaciónDigimon(digimonInfo string, id string, numDataNode int) (string, string) {
	// Dividimos la información del Digimon en partes: Nombre, Atributo, Estado
	partes := strings.Split(digimonInfo, ",")
	if len(partes) != 3 {
		return "", "Error: Formato de entrada incorrecto"
	}
	nombre := partes[0]
	atributo := partes[1]
	estado := partes[2]

	// Convertir numDataNode (int) a string
	numDataNodeStr := strconv.Itoa(numDataNode)

	// Formato para el archivo INFO.txt
	infoTxt := fmt.Sprintf("%s,%s,%s,%s", id, numDataNodeStr, nombre, estado)

	// Formato para el mensaje a enviar al Data Node
	mensajeDataNode := fmt.Sprintf("%s,%s", id, atributo)

	return infoTxt, mensajeDataNode
}

// Guardar una línea en el archivo especificado
func GuardarLineaEnArchivo(linea string, nombreArchivo string) error {
	file, err := os.OpenFile(nombreArchivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Error al abrir el archivo: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(linea + "\n")
	if err != nil {
		return fmt.Errorf("Error al escribir en el archivo: %v", err)
	}

	return nil
}

// LeerDigimonsSacrificados busca los Digimons sacrificados y retorna su nombre y ID
func LeerDigimonsSacrificados(nombreArchivo string) ([]string, error) {
	file, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer file.Close()

	var resultado []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Leer cada línea del archivo
		linea := scanner.Text()
		// Separar los campos por coma
		partes := strings.Split(linea, ",")

		if len(partes) != 4 {
			continue // Si la línea no tiene el formato esperado, la omitimos
		}

		id := partes[0]
		nombre := partes[2]
		estado := partes[3]

		// Si el estado es "Sacrificado", agregamos "nombre,id" al resultado
		if estado == "Sacrificado" {
			resultado = append(resultado, fmt.Sprintf("%s,%s", nombre, id))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error al leer el archivo: %v", err)
	}

	return resultado, nil
}

func ForSimple(array []string, datanodes []string) []string {
	var resultado []string

	for _, linea := range array {
		partes := strings.Split(linea, ",")
		aux := SeleccionarDataNode(partes[0])
		// fmt.Println("aux: ", aux)
		conn, err := Conectar(datanodes[aux])
		if err != nil {
			// Manejar el error, por ejemplo, registrarlo
			continue
		}
		// fmt.Println("partes[1]: ", partes[1])
		res, err := EnviarConsulta(conn, "2", partes[1])
		// fmt.Println("res: ", res)
		if err != nil {
			// Manejar el error, por ejemplo, registrarlo
			fmt.Println(err)
			continue
		}
		// fmt.Println("res", res)
		atributo := strings.Split(res, ",")

		resultado = append(resultado, atributo[1])
	}

	return resultado
}

// SumaDatosTransf recibe un array de atributos de los Digimon y devuelve un string con la suma total de los datos a transferir
func SumaDatosTransf(atributos []string) string {
	var suma float64
	for _, atributo := range atributos {
		// Dependiendo del atributo, sumamos los valores correspondientes
		switch atributo {
		case "Vaccine":
			suma += 3.0
		case "Data":
			suma += 1.5
		case "Virus":
			suma += 0.8
		}
	}

	// Convertimos la suma total a string y la retornamos
	return strconv.FormatFloat(suma, 'f', 2, 64)
}

// CalcularPorcentajes recibe el nombre de un archivo y calcula el porcentaje de Digimon sacrificados y no sacrificados
func CalcularPorcentajes(nombreArchivo string) (string, string, error) {
	file, err := os.Open(nombreArchivo)
	if err != nil {
		return "", "", fmt.Errorf("Error al abrir el archivo: %v", err)
	}
	defer file.Close()

	var sacrificados, noSacrificados int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()
		partes := strings.Split(linea, ",")
		if len(partes) < 4 {
			continue // Saltar líneas que no cumplan el formato
		}

		estado := strings.TrimSpace(partes[3])
		if estado == "Sacrificado" {
			sacrificados++
		} else {
			noSacrificados++
		}
	}

	if err := scanner.Err(); err != nil {
		return "", "", fmt.Errorf("Error al leer el archivo: %v", err)
	}

	total := sacrificados + noSacrificados
	if total == 0 {
		return "", "", fmt.Errorf("No se encontraron registros válidos en el archivo")
	}

	// Convertir a porcentaje y luego a string
	porcentajeSacrificados := fmt.Sprintf("%.2f%%", (float64(sacrificados)/float64(total))*100)
	porcentajeNoSacrificados := fmt.Sprintf("%.2f%%", (float64(noSacrificados)/float64(total))*100)

	return porcentajeSacrificados, porcentajeNoSacrificados, nil
}

// -------------------------------------------Conectar con server----------------------------------------------------

// Función que crea la conexión y escucha en el puerto 8080
func startServer(port string) {
	// Escuchar en el puerto 8080
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	// Crear un nuevo servidor gRPC
	grpcServer := grpc.NewServer()

	// Registrar el servicio en el servidor gRPC
	pb.RegisterDigimonServiceServer(grpcServer, &server{})

	log.Printf("gRPC server running on port %s...\n", port)

	// Iniciar el servidor y escuchar las conexiones
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

// Conectar es una función que recibe un puerto y devuelve la conexión gRPC
func Conectar(port string) (*grpc.ClientConn, error) {
	// Crear la conexión con el servidor gRPC
	conn, err := grpc.Dial(port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("Error al conectar al servidor en el puerto %s: %v", port, err)
	}

	return conn, nil
}

// EnviarConsulta es una función que envía la solicitud gRPC al Data Node
func EnviarConsulta(conn *grpc.ClientConn, opt, todo string) (string, error) {
	// Crear un cliente para el servicio DigimonService
	client := pb.NewDigimonServiceClient(conn)

	// Contexto con timeout para la solicitud
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Crear la solicitud con los parámetros proporcionados
	request := &pb.DigimonRequest{Opt: opt, Todo: todo}

	// Llamar al método GetDigimonStatus del servidor
	response, err := client.GetDigimonStatus(ctx, request)
	if err != nil {
		return "", fmt.Errorf("error al llamar al método GetDigimonStatus: %v", err)
	}

	// Devolver la respuesta del servidor
	return response.Status, nil
}

// ConectarAServidores conecta a una lista de direcciones de servidores gRPC, envía una solicitud y procesa la respuesta
func ConectarAServidores(servidores []string) {
	for _, servidor := range servidores {
		// fmt.Printf("%s\n", servidor)
		// Conectar al servidor
		conn, err := grpc.Dial(servidor, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("No se pudo conectar al servidor %s: %v", servidor, err)
		}
		defer conn.Close()

		// Crear el cliente del servicio Digimon
		client := pb.NewDigimonServiceClient(conn)

		// Establecer un contexto con timeout
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		// Hacer la solicitud al servidor
		client.GetDigimonStatus(ctx, &pb.DigimonRequest{Opt: "-1", Todo: "Terminada todo maldito"})
	}
}

func disponible(servidor string, delay time.Duration) {
	// Intentar conectar a RabbitMQ con reintentos
	for i := 0; i < 8; i++ {
		_, err := grpc.Dial(servidor, grpc.WithInsecure(), grpc.WithBlock())
		if err == nil {
			fmt.Println("Connected to server")
			return
		}

		log.Printf("Failed to connect to server: %v", err)
		log.Printf("Retrying in %v seconds...", delay.Seconds())
		time.Sleep(delay * time.Second)
	}

	log.Fatalf("Could not connect to server after %d retries", 8)
}

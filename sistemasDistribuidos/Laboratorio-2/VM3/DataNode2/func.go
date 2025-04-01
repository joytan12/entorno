package main

import (
	pb "DataNode2/proto"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"
)

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

	log.Println("gRPC server running on port ", port)

	// Iniciar el servidor y escuchar las conexiones
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

// Busca un ID en un archivo y retorna la línea completa
func buscarID(id string, filePath string) (string, error) {
	// Abrir el archivo
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer file.Close()

	// Leer el archivo línea por línea
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()

		// Separar los campos por coma
		datos := strings.Split(linea, ",")

		// Verificar si el ID coincide con el proporcionado
		if datos[0] == id {
			// Retornar la línea completa tal cual
			return linea, nil
		}
	}

	// Si no se encuentra el ID
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error al leer el archivo: %v", err)
	}

	return "", fmt.Errorf("ID %s no encontrado", id)
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

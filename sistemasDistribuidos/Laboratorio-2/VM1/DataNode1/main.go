package main

import (
	"context"
	"fmt"
	"os"

	pb "DataNode1/proto" // Asegúrate de que sea el path correcto generado a partir del .proto
)

const (
	port = ":8081" // Puerto donde escucha el Data Node
)

// Implementación del servidor gRPC
type server struct {
	pb.UnimplementedDigimonServiceServer
}

// Implementación del método GetDigimonStatus
func (s *server) GetDigimonStatus(ctx context.Context, req *pb.DigimonRequest) (*pb.DigimonResponse, error) {
	// Guardar la información del Digimon en un archivo local (INFO_1.txt o INFO_2.txt)
	if req.Opt == "1" {
		err := GuardarLineaEnArchivo(req.Todo, "INFO_1.txt")
		if err != nil {
			return nil, fmt.Errorf("error al guardar la información del Digimon: %v", err)
		}

		// Enviar confirmación al cliente
		respuesta := fmt.Sprintf("Información de %s almacenada correctamente", req.Todo)
		return &pb.DigimonResponse{Status: respuesta}, nil
	}
	if req.Opt == "-1" {
		fmt.Println("Apagando DataNode1")
		os.Exit(0)
	}

	// Buscar el ID del Digimon en los archivos
	respuesta, err := buscarID(req.Todo, "INFO_1.txt") // Cambié el nombre de error a err y añadí el segundo argumento
	if err != nil {
		return nil, fmt.Errorf("error al buscar el ID del Digimon: %v", err)
	}

	return &pb.DigimonResponse{Status: respuesta}, nil
}

func main() {
	startServer(port)
}

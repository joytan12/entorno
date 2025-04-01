package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "Supervisor2/proto" // Asegúrate de que sea el path correcto generado a partir del .proto

	"google.golang.org/grpc"
)

const (
	port = ":8086" // Puerto donde escucha el Supervisor 2
)

// Implementación del servidor gRPC
type server struct {
	pb.UnimplementedDigimonServiceServer
}

// Implementación del método GetDigimonStatus
func (s *server) GetDigimonStatus(ctx context.Context, req *pb.DigimonRequest) (*pb.DigimonResponse, error) {
	// Enviar confirmación al cliente
	respuesta := fmt.Sprintf("Información de %s almacenada correctamente", req.Todo)
	return &pb.DigimonResponse{Status: respuesta}, nil
}

func main() {
	// Crear el listener en el puerto definido
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error al iniciar el listener: %v", err)
	}

	// Crear un servidor gRPC
	grpcServer := grpc.NewServer()
	pb.RegisterDigimonServiceServer(grpcServer, &server{})

	log.Printf("Supervisor 2 escuchando en el puerto %s", port[1:])

	// Iniciar el servidor gRPC
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

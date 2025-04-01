package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "ContinenteServer/proto"

	"google.golang.org/grpc"
)

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

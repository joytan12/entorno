package main

import (
	"context"
	"fmt"
	"time"

	pb "NodeTai/proto"

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

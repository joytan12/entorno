package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"client/proto"

	"google.golang.org/grpc"
)

func processFile(filename string, client proto.PedidoServiceClient) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Error al abrir el archivo: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields, err := parseLine(line)
		if err != nil {
			return err
		}

		// Enviar la solicitud al servidor
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		req := &proto.GenerarOrden{
			Id:      fields[0],
			Tipo:    fields[1],
			Nombre:  fields[2],
			Valor:   fields[3],
			Escolta: fields[4],
			Destino: fields[5],
		}

		res, err := client.GetPedidoStatus(ctx, req)
		if err != nil {
			return fmt.Errorf("Error al llamar al servidor de Logística: %v", err)
		}

		// Mostrar la respuesta del servidor
		fmt.Printf("Número de seguimiento recibido para ID %s: %s\n", fields[0], res.GetNumeroSeguimiento())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Error al leer el archivo: %v", err)
	}

	return nil
}

func parseLine(line string) ([]string, error) {
	data := strings.TrimSpace(line)
	fields := strings.Split(data, ",")
	if len(fields) != 6 {
		return nil, fmt.Errorf("Error: la línea debe contener exactamente 6 campos separados por comas")
	}

	return fields, nil
}

func main() {
	time.Sleep(35 * time.Second)
	// Configurar conexión al servidor
	conn, err := grpc.Dial("dist041:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar al servidor de Logística: %v", err)
	}
	defer conn.Close()

	client := proto.NewPedidoServiceClient(conn)

	// Leer y procesar el archivo
	err = processFile("input.txt", client)
	if err != nil {
		log.Fatalf("Error al procesar el archivo: %v", err)
	}
}

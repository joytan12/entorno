package main

import (
	"context"
	"fmt"
	"os"

	pb "PrimaryNode/proto" // Asegúrate de que este es el path correcto generado a partir del .proto
)

// aqui estaran todas los port para las conexiones
var regionales = []string{"continenteserver:8083", "islafile:8084", "continentefolder:8085"}
var dataNodes = []string{"datanode1:8081", "datanode2:8082"}

const (
	port = ":8080" // Puerto donde escucha el PrimaryNode
)

// Implementación del servidor gRPC
type server struct {
	pb.UnimplementedDigimonServiceServer
}

// Implementación del método GetDigimonStatus
func (s *server) GetDigimonStatus(ctx context.Context, req *pb.DigimonRequest) (*pb.DigimonResponse, error) {
	// Guardar la información del Digimon en un archivo local (INFO_1.txt o INFO_2.txt)
	if req.Opt == "1" {
		id := GenerarID()
		dataNode := SeleccionarDataNode(req.Todo)
		infoTxt, mensajeDataNode := FormatearInformaciónDigimon(req.Todo, id, dataNode)
		err := GuardarLineaEnArchivo(infoTxt, "INFO.txt")
		if err != nil {
			return nil, fmt.Errorf("error al guardar la información del Digimon: %v", err)
		}

		conn, _ := Conectar(dataNodes[dataNode])
		res, _ := EnviarConsulta(conn, "1", mensajeDataNode)

		// Imprimir la respuesta del servidor
		fmt.Printf("Respuesta del servidor: %s\n", res)

		// Enviar confirmación al cliente
		respuesta := fmt.Sprintf("Información de %s almacenada correctamente en el datanode%d", req.Todo, dataNode+1)
		return &pb.DigimonResponse{Status: respuesta}, nil
	}
	if req.Opt == "2" { //con la option 2, son mensajes recibidos por el nodetai
		idsSacrificados, _ := LeerDigimonsSacrificados("INFO.txt")
		// fmt.Println("Sacrificados: ", idsSacrificados)
		atributos := ForSimple(idsSacrificados, dataNodes)
		// fmt.Println("Atributos de los sacrificados", atributos)
		danno := SumaDatosTransf(atributos)
		return &pb.DigimonResponse{Status: danno}, nil
	}
	if req.Opt == "3" { //con la option 3 deberian finalizar los servidores
		// ConectarAServidores()
		porcentajeSacrificados, porcentajeNoSacrificados, _ := CalcularPorcentajes("INFO.txt")
		fmt.Printf("Porcentaje NO Sacrificados: %s\n", porcentajeNoSacrificados)
		fmt.Printf("Porcentaje Sacrificados: %s\n", porcentajeSacrificados)
		ConectarAServidores(regionales)
		ConectarAServidores(dataNodes)
		os.Exit(0)
	}

	return &pb.DigimonResponse{Status: "chupalo manito"}, nil
}

func main() {
	disponible(dataNodes[0], 5)
	disponible(dataNodes[0], 5)
	startServer(port)
}

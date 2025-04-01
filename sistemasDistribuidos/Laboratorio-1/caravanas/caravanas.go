package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sort"
	"strings"
	"time"

	"caravanas/proto"

	"google.golang.org/grpc"
)

var TiempoDeEspera int = 1
var TiempoDestino1 int = 3
var TiempoDestino2 int = 4
var TiempoDestino3 int = 5
var TiempoDestino4 int = 6
var horario *time.Location

func init() {
	// Cargar la zona horaria de Chile al iniciar el programa
	var err error
	horario, err = time.LoadLocation("America/Santiago")
	if err != nil {
		log.Fatalf("Error cargando zona horaria: %v", err)
	}
}

func crearCaravanas() []Caravana {
	caravanaOstronitas1 := Caravana{
		ID:          "CaravanaOstronitas1",
		Tipo:        "Ostronitas",
		Estado:      "En Cetus", // Inicialmente
		Paquetes:    []Paquete{},
		MaxPaquetes: 2,
	}

	caravanaOstronitas2 := Caravana{
		ID:          "CaravanaOstronitas2",
		Tipo:        "Ostronitas",
		Estado:      "En Cetus", // Inicialmente
		Paquetes:    []Paquete{},
		MaxPaquetes: 2,
	}

	caravanaGeneral := Caravana{
		ID:          "CaravanaGeneral",
		Tipo:        "General",
		Estado:      "En Cetus", // Inicialmente
		Paquetes:    []Paquete{},
		MaxPaquetes: 2,
	}

	return []Caravana{caravanaOstronitas1, caravanaOstronitas2, caravanaGeneral}
}

func asignarPaqueteACaravana(paquete Paquete, caravanas []Caravana, caravanaID int) bool {
	caravanas[caravanaID].Paquetes = append(caravanas[caravanaID].Paquetes, paquete)
	return true
	// // Verificar el tipo del paquete para asignarlo a la caravana adecuada
	// switch paquete.Tipo {
	// case "Ostronitas":
	// 	// Intentar asignar el paquete a cualquiera de las caravanas Ostronitas
	// 	for i := 0; i < len(caravanas); i++ {
	// 		if caravanas[i].Tipo == "Ostronitas" && len(caravanas[i].Paquetes) < caravanas[i].MaxPaquetes {
	// 			caravanas[i].Paquetes = append(caravanas[i].Paquetes, paquete)
	// 			fmt.Printf("Paquete %s asignado a la %s\n", paquete.IDpaquete, caravanas[i].ID)
	// 			if len(caravanas[i].Paquetes) == 1 {
	// 				caravanas[i].esperarPaquete()
	// 			}
	// 			return true
	// 		}
	// 	}
	// case "Prioritario":
	// 	// Intentar asignar primero a las caravanas Ostronitas
	// 	for i := 0; i < len(caravanas); i++ {
	// 		if len(caravanas[i].Paquetes) < caravanas[i].MaxPaquetes {
	// 			// Se asignan prioritarios a cualquier caravana con espacio
	// 			caravanas[i].Paquetes = append(caravanas[i].Paquetes, paquete)
	// 			fmt.Printf("Paquete %s (Prioritario) asignado a la %s\n", paquete.IDpaquete, caravanas[i].ID)
	// 			if len(caravanas[i].Paquetes) == 1 {
	// 				caravanas[i].esperarPaquete()
	// 			}
	// 			return true
	// 		}
	// 	}
	// case "Normal":
	// 	// Intentar asignar el paquete a la CaravanaGeneral
	// 	for i := 0; i < len(caravanas); i++ {
	// 		if caravanas[i].ID == "CaravanaGeneral" && len(caravanas[i].Paquetes) < caravanas[i].MaxPaquetes {
	// 			caravanas[i].Paquetes = append(caravanas[i].Paquetes, paquete)
	// 			fmt.Printf("Paquete %s (Normal) asignado a la %s\n", paquete.IDpaquete, caravanas[i].ID)
	// 			if len(caravanas[i].Paquetes) == 1 {
	// 				caravanas[i].esperarPaquete()
	// 			}
	// 			return true
	// 		}
	// 	}
	// default:
	// 	fmt.Printf("Tipo de paquete desconocido: %s\n", paquete.Tipo)
	// 	return false
	// }

	// fmt.Printf("No se pudo asignar el paquete %s. Todas las caravanas están llenas o no disponibles.\n", paquete.IDpaquete)
	// return false
}

func (c *Caravana) esperarPaquete() {
	fmt.Printf("La caravana %s está esperando por un segundo paquete.\n", c.ID)
	time.Sleep(time.Duration(TiempoDeEspera) * time.Second)
	fmt.Printf("La caravana %s ha dejado de esperar.\n", c.ID)
}

type Paquete struct {
	IDpaquete   string
	Tipo        string // Ostronita, Prioritario o Normal
	Valor       string
	Escolta     string
	Estado      string
	Intentos    int
	Seguimiento string
	Destino     string
}

type Caravana struct {
	ID          string
	Tipo        string // Ostronitas o General
	Estado      string
	Paquetes    []Paquete
	MaxPaquetes int
}

type Registro struct {
	IDpaquete    string
	Tipo         string
	Valor        string
	Escolta      string
	Destino      string
	Intentos     int
	FechaEntrega time.Time
}

var registroDePaquetes = make(map[string]Registro)

type RegistroEntrega struct {
	IDPaquete    string
	Tipo         string
	Valor        int
	Escolta      string
	Destino      string
	Intentos     int
	FechaEntrega time.Time
}

type CaravanaServiceServer struct {
	proto.UnimplementedPreguntarEstadoCaravanasServer
	proto.UnimplementedIngresoCaravanaServer
	proto.UnimplementedEstadoPaquetesCaravanasServer
}

// Función para enviar el estado de los paquetes al servidor
func enviarEstadoPaqueteAlServidor(paquete Registro) {
	// Configura la conexión gRPC con el servidor
	conn, err := grpc.Dial("dist041:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar con el servidor: %v", err)
	}
	defer conn.Close()

	client := proto.NewEstadoPaquetesCaravanasClient(conn)

	// Crea el registro
	registro := proto.RecibirEstadoPaquetesCaravanas{
		IDpaquete:    paquete.IDpaquete,
		Tipo:         paquete.Tipo,
		Valor:        paquete.Valor,
		Escolta:      paquete.Escolta,
		Destino:      paquete.Destino,
		Intentos:     fmt.Sprintf("%d", paquete.Intentos), // Conversión de int a string
		FechaEntrega: time.Now().In(horario).Format("02-01-2006 15:04:05"),
	}

	// Enviar la solicitud al servidor
	resp, err := client.GetEstadoPaquetesCaravanas(context.Background(), &registro)
	if err != nil {
		log.Fatalf("Error al enviar el registro al servidor: %v", err)
	}

	// Confirmación de que el servidor recibió el estado
	log.Printf("Respuesta del servidor: %s", resp.Estado)
}

func ContenidoCaravanas(caravanas []Caravana) map[string]int {
	espaciosDisponibles := make(map[string]int)

	for _, caravana := range caravanas {
		// fmt.Printf("ID: %s\n", caravana.ID)
		// fmt.Printf("Tipo: %s\n", caravana.Tipo)
		// fmt.Printf("Max Paquetes: %d\n", caravana.MaxPaquetes)
		// fmt.Println("Paquetes:")
		// for _, paquete := range caravana.Paquetes {
		// 	fmt.Printf("  - ID: %s, Tipo: %s, Valor: %s, Escolta: %s, Destino: %s\n",
		// 		paquete.IDpaquete, paquete.Tipo, paquete.Valor, paquete.Escolta, paquete.Destino)
		// }
		// fmt.Println()
		// Calcular espacios disponibles
		espacios := caravana.MaxPaquetes - len(caravana.Paquetes)
		espaciosDisponibles[caravana.ID] = espacios
		// fmt.Printf("Espacios de ID %s disponibles: %d\n", caravana.ID, espaciosDisponibles[caravana.ID])
	}

	// fmt.Printf("Espacios disponibleseees: %s.\n", espaciosDisponibles)

	return espaciosDisponibles
}

var caravanas = crearCaravanas()

// Implementación del método GetEstadoCaravanas
func (s *CaravanaServiceServer) GetEstadoCaravanas(ctx context.Context, req *proto.SolicitarEstadoCaravanas) (*proto.EstadoCaravanas, error) {
	// fmt.Println("El servidor principal me está preguntando el estado de las caravanas.")
	espacios := ContenidoCaravanas(caravanas)
	// fmt.Println(espacios)
	return &proto.EstadoCaravanas{
		Caravana1: int32(espacios["CaravanaOstronitas1"]),
		Caravana2: int32(espacios["CaravanaOstronitas2"]),
		Caravana3: int32(espacios["CaravanaGeneral"]),
	}, nil
}

func (s *CaravanaServiceServer) GetIngresoCaravana(ctx context.Context, req *proto.IngresarACaravana) (*proto.RespuestaIngreso, error) {
	caravanaID := req.GetCaravana()
	paquete := Paquete{
		IDpaquete:   req.GetIdPaquete(),
		Seguimiento: req.GetNumeroSeguimiento(),
		Tipo:        req.GetTipo(),
		Valor:       req.GetValor(),
		Escolta:     req.GetEscolta(),
		Destino:     req.GetDestino(),
		Intentos:    0,
	}
	// fmt.Printf("ID: %d\n", caravanaID)

	// if caravanaID == 1 {
	// 	caravanaID == "CaravanaOstronitas1"
	// } else if caravanaID == 2 {

	// }

	// Obtener la caravana del mapa
	caravana := caravanas[caravanaID-1]

	// Verificar si hay espacio disponible
	if len(caravana.Paquetes) >= caravana.MaxPaquetes {
		return &proto.RespuestaIngreso{Estado: "No hay espacio disponible en la caravana"}, nil
	}

	// Añadir el paquete a la caravana
	caravana.Paquetes = append(caravana.Paquetes, paquete)

	// Actualizar la caravana en el mapa
	caravanas[caravanaID-1] = caravana

	// fmt.Printf("Espacios disponibles: %d, %d, %d.\n", ContenidoCaravanas(caravanas)["CaravanaOstronitas1"], ContenidoCaravanas(caravanas)["CaravanaOstronitas2"], ContenidoCaravanas(caravanas)["CaravanaGeneral"])
	// fmt.Printf("test %s")
	// Retornar respuesta
	return &proto.RespuestaIngreso{Estado: "Paquete asignado correctamente"}, nil
}

func mostrarRegistroDePaquetes() {
	if len(registroDePaquetes) == 0 {
		fmt.Println("No hay paquetes en el registro.")
		return
	}

	fmt.Printf("%-15s %-15s %-10s %-10s %-10s %-10s %-20s\n", "IDpaquete", "Tipo", "Valor", "Escolta", "Destino", "Intentos", "FechaEntrega")
	fmt.Println(strings.Repeat("-", 80)) // Línea separadora

	for _, registro := range registroDePaquetes {
		fmt.Printf("%-15s %-15s %-10s %-10s %-10s %-10d %-20s\n",
			registro.IDpaquete,
			registro.Tipo,
			registro.Valor,
			registro.Escolta,
			registro.Destino,
			registro.Intentos,
			registro.FechaEntrega.Format("02-01-2006 15:04:05"),
		)
	}
}

func enviarCaravanaADestino(caravana Caravana) {
	// Recordar eliminar paquetes de las caravanas
	fmt.Printf("Enviando %s\n", caravana.ID)
	sort.Slice(caravana.Paquetes, func(i, j int) bool {
		return caravana.Paquetes[i].Valor > caravana.Paquetes[j].Valor
	})

	for i := 0; i < len(caravana.Paquetes); i++ {
		paquete := &caravana.Paquetes[i]
		paquete.Estado = "En camino"
		var tiempoDeLlegada int

		// Determinar el tiempo de llegada según el destino
		switch paquete.Destino {
		case "Destino1":
			tiempoDeLlegada = TiempoDestino1
		case "Destino2":
			tiempoDeLlegada = TiempoDestino2
		case "Destino3":
			tiempoDeLlegada = TiempoDestino3
		case "Destino4":
			tiempoDeLlegada = TiempoDestino4
		default:
			fmt.Printf("Destino desconocido para el paquete %s\n", paquete.IDpaquete)
			continue
		}

		// Enviar el paquete con reintentos en caso de falla
		var entregado bool
		for intento := 1; intento <= 3; intento++ {
			// Simular tiempo de envío
			time.Sleep(time.Duration(tiempoDeLlegada) * time.Second)
			paquete.Intentos = intento
			// Simular entrega con un 85% de probabilidad de éxito
			if rand.Float64() <= 0.85 {
				entregado = true
				break
			}
			fmt.Printf("Cantidad de intentos hasta ahora %d\n", intento)
		}

		// Actualizar estado del paquete
		if entregado {
			paquete.Estado = "Entregado"
		} else {
			paquete.Estado = "No Entregado"
		}

		// Crear registro del paquete
		registro := Registro{
			IDpaquete:    paquete.IDpaquete,
			Tipo:         paquete.Tipo,
			Valor:        paquete.Valor,
			Escolta:      paquete.Escolta,
			Destino:      paquete.Destino,
			Intentos:     paquete.Intentos,
			FechaEntrega: time.Now().In(horario),
		}
		registroDePaquetes[paquete.IDpaquete] = registro

		// Enviar paquete al servidor para que haga lo financiero.
		enviarEstadoPaqueteAlServidor(registro)

		// Eliminar paquete de la caravana
		caravana.Paquetes = append(caravana.Paquetes[:i], caravana.Paquetes[i+1:]...)
		i-- // Ajustar el índice después de la eliminación
	}
	caravana.Estado = "En Cetus"
	// mostrarRegistroDePaquetes()

}

func verificarCaravanas() {
	// Definir el orden de prioridad de las caravanas
	ordenDeVerificacion := []string{"CaravanaOstronitas1", "CaravanaOstronitas2", "CaravanaGeneral"}
	// mostrarRegistroDePaquetes()

	for {
		mostrarRegistroDePaquetes()
		for _, id := range ordenDeVerificacion {
			// Buscar la caravana correspondiente
			for i, caravana := range caravanas {
				if caravana.ID == id {
					if len(caravana.Paquetes) >= caravana.MaxPaquetes {
						fmt.Printf("La caravana %s está llena. Se procederá a entregar los paquetes a destino.\n", caravana.ID)
						enviarCaravanaADestino(caravana)
						// Limpiar los paquetes de la caravana después de enviar
						caravanas[i].Paquetes = []Paquete{}
					}
					// Salir del bucle de búsqueda para pasar a la siguiente caravana en el orden
					break
				}
			}
		}

		time.Sleep(5 * time.Second)
	}
}

func main() {
	rand.Seed(time.Now().In(horario).UnixNano())
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor de Caravanas: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterPreguntarEstadoCaravanasServer(grpcServer, &CaravanaServiceServer{})
	proto.RegisterIngresoCaravanaServer(grpcServer, &CaravanaServiceServer{})
	proto.RegisterEstadoPaquetesCaravanasServer(grpcServer, &CaravanaServiceServer{})

	go verificarCaravanas()

	log.Printf("Servidor de Caravanas escuchando en el puerto 8081")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al ejecutar el servidor de Caravanas: %v", err)
	}

}

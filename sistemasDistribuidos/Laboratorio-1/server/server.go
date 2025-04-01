package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"server/proto" // Importamos el proto de logistica
	"time"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

type Orden struct {
	Timestamp   time.Time
	IDpaquete   string
	Tipo        string
	Nombre      string
	Valor       string
	Escolta     string
	Destino     string
	Seguimiento string
}

type Paquete struct {
	IDpaquete   string
	Seguimiento string
	Tipo        string
	Valor       string
	Intentos    string
	Estado      string
	Destino     string
	Escolta     string
}

var registroDeOrdenes = make(map[string]Orden)
var colaOstronitas []Paquete
var colaPrioritario []Paquete
var colaNormal []Paquete
var gananciasTotales int = 0
var horario *time.Location

func init() {
	// Cargar la zona horaria de Chile al iniciar el programa
	var err error
	horario, err = time.LoadLocation("America/Santiago")
	if err != nil {
		log.Fatalf("Error cargando zona horaria: %v", err)
	}
}

// var contadorPaquetes = 1

// Función para generar IDpaquete en formato "000001"
// func generarIDpaquete() string {

// 	return fmt.Sprintf("%06d", contadorPaquetes)
// }

// Función para recibir la orden y generar el número de seguimiento
func recibirOrden(orden Orden) string {
	fmt.Printf("Se ingresó la orden N°%s\n", orden.IDpaquete)
	orden.Timestamp = time.Now().In(horario)
	orden.Seguimiento = generarNumeroDeSeguimiento()
	registroDeOrdenes[orden.Seguimiento] = orden
	// time.Sleep(2 * time.Second)
	return orden.Seguimiento
}

// Genera un número de seguimiento único basado en el timestamp actual
func generarNumeroDeSeguimiento() string {
	for {
		// Convertir el timestamp en una cadena de texto
		seguimiento := fmt.Sprintf("%d", time.Now().In(horario).UnixNano())

		// Obtener los últimos 7 u 8 caracteres del timestamp
		if len(seguimiento) > 8 {
			seguimiento = seguimiento[len(seguimiento)-8:]
		}

		// Verificar si el número ya está en uso
		if _, exists := registroDeOrdenes[seguimiento]; !exists {
			// Si no existe, retornar el número de seguimiento
			return seguimiento
		}

		// Si existe, se genera otro número en la siguiente iteración
	}
}

func mostrarOrdenes() {
	if len(registroDeOrdenes) == 0 {
		fmt.Println("No hay órdenes en el registro.")
		return
	}
	for _, orden := range registroDeOrdenes {
		fmt.Printf("ID Paquete: %s\n", orden.IDpaquete)
		fmt.Printf("Timestamp: %s\n", orden.Timestamp.Format("02-01-2006 15:04:05"))
		fmt.Printf("Tipo: %s\n", orden.Tipo)
		fmt.Printf("Nombre: %s\n", orden.Nombre)
		fmt.Printf("Valor: %s\n", orden.Valor)
		fmt.Printf("Escolta: %s\n", orden.Escolta)
		fmt.Printf("Destino: %s\n", orden.Destino)
		fmt.Printf("Seguimiento: %s\n", orden.Seguimiento)
		fmt.Println("-----------------------------")
	}
}

func agregarACola(paquete Paquete) {
	switch paquete.Tipo {
	case "Ostronitas":
		colaOstronitas = append(colaOstronitas, paquete)
	case "Prioritario":
		colaPrioritario = append(colaPrioritario, paquete)
	default:
		colaNormal = append(colaNormal, paquete)
	}
}

func procesarOrden(orden *proto.GenerarOrden, numeroSeguimiento string) {
	paquete := Paquete{
		IDpaquete:   orden.GetId(),
		Seguimiento: numeroSeguimiento,
		Tipo:        orden.GetTipo(),
		Valor:       orden.GetValor(),
		Destino:     orden.GetDestino(),
		Escolta:     orden.GetEscolta(),
		Intentos:    "0",         // Inicialmente cero intentos
		Estado:      "pendiente", // Estado inicial
	}

	agregarACola(paquete)
	fmt.Printf("Paquete %s agregado a la cola %s\n", paquete.IDpaquete, paquete.Tipo)
}

// Función para consultar el estado de las caravanas disponibles
func consultarEstadoCaravanas(ctx context.Context) (*proto.EstadoCaravanas, error) {
	// Conectarse al servidor de caravanas en el puerto 8081
	conn, err := grpc.Dial("dist043:8081", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("no se pudo conectar al servidor de Caravanas: %v", err)
	}
	defer conn.Close()

	// Crear el cliente del servicio PreguntarEstadoCaravanas
	caravanaClient := proto.NewPreguntarEstadoCaravanasClient(conn)

	// Realiza la solicitud para consultar el estado de las caravanas
	req := &proto.SolicitarEstadoCaravanas{
		Solicitud: "Estado",
	}

	// Enviar la solicitud al servidor de caravanas
	resp, err := caravanaClient.GetEstadoCaravanas(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error al consultar el estado de las caravanas: %v", err)
	}

	return resp, nil
}

// Función para ingresar el paquete a una caravana
func ingresarPaqueteACaravana(orden Orden, caravana int32) (*proto.RespuestaIngreso, error) {
	conn, err := grpc.Dial("dist043:8081", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("No se pudo conectar al servidor de Caravanas: %v", err)
	}
	defer conn.Close()

	client := proto.NewIngresoCaravanaClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &proto.IngresarACaravana{
		Caravana:          caravana,
		IdPaquete:         orden.IDpaquete,
		Tipo:              orden.Tipo,
		Valor:             orden.Valor,
		Escolta:           orden.Escolta,
		Destino:           orden.Destino,
		NumeroSeguimiento: orden.Seguimiento,
	}

	res, err := client.GetIngresoCaravana(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Error al ingresar paquete a la caravana: %v", err)
	}
	fmt.Printf("Res %s\n", res)
	return res, nil
}

// Generar Orden
type PedidoServiceServer struct {
	proto.UnimplementedPedidoServiceServer
}

type ServerService struct {
	proto.UnimplementedEstadoPaquetesCaravanasServer
}

func (s *ServerService) GetEstadoPaquetesCaravanas(ctx context.Context, req *proto.RecibirEstadoPaquetesCaravanas) (*proto.ConfirmarRecibimientoEstado, error) {
	// Procesar el paquete recibido
	log.Printf("Servidor recibió paquete: ID: %s, Tipo: %s, Valor: %s, Escolta: %s, Destino: %s, Intentos: %s, FechaEntrega: %s",
		req.IDpaquete, req.Tipo, req.Valor, req.Escolta, req.Destino, req.Intentos, req.FechaEntrega)

	// Aquí se conecta con el sistema financiero
	conn := connectRabbitMQ(10, 3*time.Second)
	defer conn.Close()

	body := fmt.Sprintf("%s,%s,%s,%s", req.IDpaquete, req.Tipo, req.Valor, req.Intentos)
	// body := "mensaje de prueba xd"
	// if err := handleFinanciero(body, conn); err != nil {
	// 	log.Fatalf("Error al enviar mensaje: %v", err)
	// }
	handleFinanciero(body, conn)
	// Responder con una confirmación
	return &proto.ConfirmarRecibimientoEstado{
		Estado: "Registro recibido exitosamente",
	}, nil
}

// Conexión con Financiero:

type Estado struct {
	financiero bool
	carabanas  bool
}

func connectRabbitMQ(retries int, delay time.Duration) *amqp.Connection {
	rabbitMQURL := os.Getenv("RABBITMQ_URL")

	var conn *amqp.Connection
	var err error

	// Intentar conectar a RabbitMQ con reintentos
	for i := 0; i < retries; i++ {
		conn, err = amqp.Dial(rabbitMQURL)
		if err == nil {
			fmt.Println("Connected to RabbitMQ")
			return conn
		}

		log.Printf("Failed to connect to RabbitMQ: %v", err)
		log.Printf("Retrying in %v seconds...", delay.Seconds())
		time.Sleep(delay)
	}

	log.Fatalf("Could not connect to RabbitMQ after %d retries", retries)
	return nil
}

func handleFinanciero(body string, conn *amqp.Connection) {
	// Crear un canal de comunicación
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declarar una cola (queue)
	q, err := ch.QueueDeclare(
		"task_queue", // nombre de la cola
		true,         // durable
		false,        // borrar cuando no se use
		false,        // no exclusiva
		false,        // sin espera
		nil,          // argumentos adicionales
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// El cuerpo del mensaje que se va a enviar
	err = ch.Publish(
		"",     // intercambio
		q.Name, // clave de enrutamiento (routing key)
		false,  // obligatorio
		false,  // inmediato
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}
	fmt.Printf(" [x] Sent %s\n", body)

}

func enviarAMensajeria(body string, conn *amqp.Connection) error {
	// Crear un canal de comunicación
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declarar una cola (queue)
	q, err := ch.QueueDeclare(
		"task_queue", // nombre de la cola
		true,         // durable
		false,        // borrar cuando no se use
		false,        // no exclusiva
		false,        // sin espera
		nil,          // argumentos adicionales
	)
	if err != nil {
		return fmt.Errorf("Failed to declare a queue: %v", err)
	}

	// Publicar el mensaje en la cola
	err = ch.Publish(
		"",     // intercambio
		q.Name, // clave de enrutamiento (routing key)
		false,  // obligatorio
		false,  // inmediato
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return fmt.Errorf("Failed to publish a message: %v", err)
	}

	fmt.Printf(" [x] Sent %s\n", body)
	return nil
}

func (s *PedidoServiceServer) GetPedidoStatus(ctx context.Context, req *proto.GenerarOrden) (*proto.RecibirNumeroSeguimiento, error) {
	orden := Orden{
		IDpaquete: req.GetId(),
		Tipo:      req.GetTipo(),
		Nombre:    req.GetNombre(),
		Valor:     req.GetValor(),
		Escolta:   req.GetEscolta(),
		Destino:   req.GetDestino(),
	}

	numeroSeguimiento := recibirOrden(orden) // Se ingresa al registro
	// fmt.Printf("Numero seguimiento: %s\n", numeroSeguimiento)
	procesarOrden(req, numeroSeguimiento)
	// mostrarOrdenes()

	return &proto.RecibirNumeroSeguimiento{
		NumeroSeguimiento: numeroSeguimiento,
	}, nil
}

func paqueteToOrden(paquete Paquete) Orden {
	// Buscar la orden correspondiente en el registro de órdenes
	return Orden{
		IDpaquete:   paquete.IDpaquete,
		Seguimiento: paquete.Seguimiento,
		Tipo:        paquete.Tipo,
		Valor:       paquete.Valor,
		Escolta:     paquete.Escolta,
		Destino:     paquete.Destino,
	}
}

func eliminarDeCola(cola []Paquete, i int) []Paquete {
	return append(cola[:i], cola[i+1:]...)
}

func verificarYAsignarPaquetes() {
	// mostrarOrdenes()
	for {
		ctx := context.Background()
		estadoCaravanas, err := consultarEstadoCaravanas(ctx)
		if err != nil {
			log.Printf("Error al consultar el estado de las caravanas: %v", err)
			continue
		}

		// fmt.Printf("[Server] Espacios disponibles: %d, %d, %d\n", estadoCaravanas.Caravana1, estadoCaravanas.Caravana2, estadoCaravanas.Caravana3)
		// Verificar y asignar paquetes de la cola Ostronitas
		if estadoCaravanas.Caravana1 > 0 || estadoCaravanas.Caravana2 > 0 {
			for i := 0; i < len(colaOstronitas); i++ {
				paquete := colaOstronitas[i]
				if estadoCaravanas.Caravana1 > 0 {
					_, err := ingresarPaqueteACaravana(paqueteToOrden(paquete), 1)
					if err == nil {
						fmt.Printf("Paquete %s asignado a Caravana1\n", paquete.IDpaquete)
						colaOstronitas = eliminarDeCola(colaOstronitas, i) // Eliminar paquete de la cola
						i--                                                // Reducir el índice porque la longitud de la cola ha disminuido
						estadoCaravanas.Caravana1--                        // Reducir espacio disponible en la caravana
						break
					}
				} else if estadoCaravanas.Caravana2 > 0 {
					_, err := ingresarPaqueteACaravana(paqueteToOrden(paquete), 2)
					if err == nil {
						fmt.Printf("Paquete %s asignado a Caravana2\n", paquete.IDpaquete)
						colaOstronitas = eliminarDeCola(colaOstronitas, i) // Eliminar paquete de la cola
						i--                                                // Reducir el índice porque la longitud de la cola ha disminuido
						estadoCaravanas.Caravana2--
						break
					}
				}
				// Si no hay espacios en ninguna caravana prioritaria, salir del loop
				if estadoCaravanas.Caravana1 == 0 && estadoCaravanas.Caravana2 == 0 {
					break
				}
			}
		}

		// Verificar y asignar paquetes de la cola Prioritarios
		if estadoCaravanas.Caravana3 > 0 {
			for i := 0; i < len(colaPrioritario); i++ {
				paquete := colaPrioritario[i]
				_, err := ingresarPaqueteACaravana(paqueteToOrden(paquete), 3)
				if err == nil {
					fmt.Printf("Paquete %s asignado a Caravana3\n", paquete.IDpaquete)
					colaPrioritario = eliminarDeCola(colaPrioritario, i)
					i-- // Reducir el índice porque la longitud de la cola ha disminuido
					estadoCaravanas.Caravana3--
					break
				}
				if estadoCaravanas.Caravana3 == 0 {
					break
				}
			}
		}

		// Verificar y asignar paquetes de la cola Normal
		if estadoCaravanas.Caravana3 > 0 {
			for i := 0; i < len(colaNormal); i++ {
				paquete := colaNormal[i]
				_, err := ingresarPaqueteACaravana(paqueteToOrden(paquete), 3)
				if err == nil {
					fmt.Printf("Paquete %s asignado a Caravana3\n", paquete.IDpaquete)
					colaNormal = eliminarDeCola(colaNormal, i)
					i-- // Reducir el índice porque la longitud de la cola ha disminuido
					estadoCaravanas.Caravana3--
					break
				}
				if estadoCaravanas.Caravana3 == 0 {
					break
				}
			}
		}

		// Esperar 3 segundos antes de volver a verificar
		time.Sleep(3 * time.Second)
	}
}

func main() {
	// Conexión a RabbitMQ
	conn := connectRabbitMQ(10, 5*time.Second)
	defer conn.Close()
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor de Logística: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterPedidoServiceServer(grpcServer, &PedidoServiceServer{})
	proto.RegisterEstadoPaquetesCaravanasServer(grpcServer, &ServerService{})

	go verificarYAsignarPaquetes()

	log.Printf("Servidor de Logística escuchando en el puerto 8080")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al ejecutar el servidor de Logística: %v", err)
	}

	time.Sleep(3)

	// Crear un canal de comunicación
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declarar una cola (queue)
	q, err := ch.QueueDeclare(
		"task_queue", // nombre de la cola
		true,         // durable
		false,        // borrar cuando no se use
		false,        // no exclusiva
		false,        // sin espera
		nil,          // argumentos adicionales
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Consumir los mensajes de la cola
	msgs, err := ch.Consume(
		q.Name, // nombre de la cola
		"",     // nombre del consumidor
		true,   // auto-acknowledge
		false,  // no exclusiva
		false,  // sin espera
		false,  // sin argumentos adicionales
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// Crear un canal para mantener el proceso vivo
	forever := make(chan bool)

	// Iniciar una gorutina para recibir mensajes de la cola
	go func() {
		for d := range msgs {
			fmt.Printf(" [x] Received: %s\n", d.Body)
		}
	}()

	// body := "este mensaje esta llegando de pana banana, para que te quede clarisimo"
	// handleFinanciero(body, conn)
	// handleFinanciero(body, conn)
	// Ejemplo de mensaje que se quiere enviar
	// fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	// body := "mensaje de prueba"
	// if err := enviarAMensajeria(body, conn); err != nil {
	// 	log.Fatalf("Error al enviar mensaje: %v", err)
	// }

	// Mantener el proceso vivo
	<-forever
}

// func main() {
// 	// estado := Estado{
// 	//     financiero: false,
// 	//     carabanas:  false,
// 	// }
// 	// conexion con rabbit
// 	conn := connectRabbitMQ(10, 5*time.Second)
// 	defer conn.Close()

// 	time.Sleep(3)

// 	// Crear un canal de comunicación
// 	ch, err := conn.Channel()
// 	if err != nil {
// 		log.Fatalf("Failed to open a channel: %v", err)
// 	}
// 	defer ch.Close()

// 	// Declarar una cola (queue)
// 	q, err := ch.QueueDeclare(
// 		"task_queue", // nombre de la cola
// 		true,         // durable
// 		false,        // borrar cuando no se use
// 		false,        // no exclusiva
// 		false,        // sin espera
// 		nil,          // argumentos adicionales
// 	)
// 	if err != nil {
// 		log.Fatalf("Failed to declare a queue: %v", err)
// 	}

// 	// Consumir los mensajes de la cola
// 	msgs, err := ch.Consume(
// 		q.Name, // nombre de la cola
// 		"",     // nombre del consumidor
// 		true,   // auto-acknowledge
// 		false,  // no exclusiva
// 		false,  // sin espera
// 		false,  // sin argumentos adicionales
// 		nil,
// 	)
// 	if err != nil {
// 		log.Fatalf("Failed to register a consumer: %v", err)
// 	}

// 	// Crear un canal para mantener el proceso vivo
// 	forever := make(chan bool)

// 	// Iniciar una gorutina para recibir mensajes de la cola
// 	go func() {
// 		for d := range msgs {
// 			fmt.Printf(" [x] Received: %s\n", d.Body)
// 		}
// 	}()
// 	// el body, es el mensaje que vas a querer enviar a financiero
// 	body := "mira ql este mensaje esta llegando de pana banana, para que te quede clarisimo"
// 	handleFinanciero(body, conn)
// 	handleFinanciero(body, conn)

// 	// fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
// 	<-forever
// }

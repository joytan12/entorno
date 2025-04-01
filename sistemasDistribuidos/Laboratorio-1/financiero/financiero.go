package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

type Finanzas struct {
	IDpaquete         string
	Estado            string
	Intentos          int
	CreditosGenerados int
}

var registroDeFinanzas = make(map[string]Finanzas)

func calcularTotalCreditosGenerados() int {
	totalCreditos := 0

	// Recorrer cada entrada del mapa y sumar los créditos generados
	for _, finanza := range registroDeFinanzas {
		totalCreditos += finanza.CreditosGenerados
	}

	return totalCreditos
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

func handleserver(body string, conn *amqp.Connection) {
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

// Función para enviar mensajes a otra cola
func enviarMensaje(conn *amqp.Connection, body string, cola string) {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declarar la cola a la que se enviará el mensaje
	q, err := ch.QueueDeclare(
		cola,  // nombre de la cola
		true,  // durable
		false, // borrar cuando no se use
		false, // no exclusiva
		false, // sin espera
		nil,   // argumentos adicionales
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
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
		log.Fatalf("Failed to publish a message: %v", err)
	}

	fmt.Printf(" [x] Received: %s\n", body)
}

func processMessage(body []byte) string {
	// Convertir el mensaje a string y separarlo por comas
	message := string(body)
	parts := strings.Split(message, ",")

	if len(parts) != 4 {
		// log.Printf("Mensaje malformado: %s", message)
		return ""
	}

	// Parsear los datos
	idPaquete := parts[0]
	tipo := parts[1]
	valor, err := strconv.Atoi(parts[2])
	if err != nil {
		// log.Printf("Error al parsear valor: %v", err)
		return ""
	}
	intentos, err := strconv.Atoi(parts[3])
	if err != nil {
		// log.Printf("Error al parsear intentos: %v", err)
		return ""
	}

	var resultado int

	// Realizar cálculos de acuerdo con el tipo y el número de intentos
	switch tipo {
	case "Ostronitas":
		if intentos == 1 {
			resultado = valor
		} else if intentos > 1 {
			resultado = valor - 100*intentos
		}
	case "Normal":
		if intentos == 1 {
			resultado = valor
		} else if intentos > 1 {
			resultado = valor - 100*intentos
		}
	case "Prioritario":
		if intentos == 1 {
			resultado = valor + int(math.Floor(float64(valor)*0.3))
		} else if intentos > 1 {
			resultado = valor + int(math.Floor(float64(valor)*0.3)) - 100*intentos
		}
	}

	// Aplicar ajustes especiales si los intentos son 3
	if intentos == 3 {
		switch tipo {
		case "Normal":
			resultado = 0
		case "Prioritario":
			resultado = int(math.Floor(float64(valor)*0.3)) - 100*intentos
		case "Ostronitas":
			resultado = valor - 100*intentos
		}
	}

	estado := "Completado"
	if intentos >= 3 {
		estado = "No Completado"
	}

	finanza := Finanzas{
		IDpaquete:         idPaquete,
		Estado:            estado,
		Intentos:          intentos,
		CreditosGenerados: resultado,
	}

	registroDeFinanzas[idPaquete] = finanza

	// Mostrar el resultado
	fmt.Printf("ID Paquete: %s, Resultado: %d\n", idPaquete, resultado)
	resultMessage := fmt.Sprintf("%s,%s,%d,%d,%d", idPaquete, tipo, valor, intentos, resultado)
	return resultMessage
	// return resultado
}

func main() {
	conn := connectRabbitMQ(10, 5*time.Second)
	defer conn.Close()

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
			if len(d.Body) != 0 {
				mensaje := processMessage(d.Body) // Procesar el mensaje
				// fmt.Printf("Mensaje antes de enviarlo: %s\n", mensaje)
				if len(mensaje) != 0 {
					handleserver(mensaje, conn)
					fmt.Printf("Total de créditos generados: %d\n", calcularTotalCreditosGenerados())
				}
			}
			// fmt.Printf(" [x] Received: %s\n", d.Body)
			// -----------------------------------------------------------
			// hace los que quieras con el body en esta parte o logica de financiero
		}
	}()

	// fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

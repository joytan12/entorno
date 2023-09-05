import socket

# Crea un socket TCP IPv4
server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

# Enlaza el socket a una dirección IP y puerto
host = '127.0.0.1'  # Dirección IP del servidor
port = 12345       # Puerto del servidor
server_socket.bind((host, port))

# Pon el socket en modo de escucha
server_socket.listen(1)  # El argumento es el número máximo de conexiones en cola
print(f"El servidor está escuchando en {host}:{port}")

while True:
    # Espera una conexión entrante
    client_socket, client_address = server_socket.accept()
    print(f"Se ha establecido una conexión con {client_address}")

    # Envía un mensaje al cliente
    message_to_client = "¡Hola, cliente!"
    client_socket.send(message_to_client.encode())

    # Aquí puedes manejar la comunicación con el cliente
    # Por ejemplo, puedes recibir y enviar datos utilizando client_socket.recv() y client_socket.send()

    # Cierra la conexión con el cliente cuando hayas terminado
    client_socket.close()

# Cierra el socket del servidor (esto generalmente no se alcanza en este bucle infinito)
server_socket.close()
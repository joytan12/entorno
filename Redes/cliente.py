import socket

# Crea un socket TCP IPv4
client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

# Dirección y puerto del servidor al que te quieres conectar
server_host = '127.0.0.1'
server_port = 12345

# Conecta al servidor
client_socket.connect((server_host, server_port))

'''
escribir codigo para juego
'''
# print('--------Bienvenido al Juango---------')
# print('-Seleccion una opción\n 1-Jugar \n 2-Salir')
# opcion = int(input(''))

# Envía datos al servidor
message = "Hola, servidor"
client_socket.send(message.encode())

# Recibe la respuesta del servidor
data = client_socket.recv(1024)
print("Respuesta del servidor:", data.decode())

# Cierra la conexión
client_socket.close()
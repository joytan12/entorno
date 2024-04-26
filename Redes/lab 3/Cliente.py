import socket as skt
import json

# Se declara dirección del servidor y puerto
serverAddr = 'localhost'
serverPort = 63420

# Se crea un socket para hacer el manejo de la conexión
clientSocket = skt.socket(skt.AF_INET, skt.SOCK_DGRAM)

domainName = ""
IP = ""
TTL = ""
Type = ""

option = -1
nombre = ""

data = {
    "option" : "",
    "domainName": "",
    "IP": "",
    "TTL": "",
    "Type": ""
}

while True:
    print(f"Escoja una opción:\n 1.- Ingresar valores al servidor DNS.\n 2.- Consultar por un servidor dado el nombre.\nOpción elegida:")
    option = input()
    if option == "1":
        data["option"] = "1"

        print("Ingresa el nombre del dominio: ")
        domainName = input()
        data["domainName"] = domainName
        print(f"Nombre del dominio: {domainName}")

        print("Ingresa la IP: ")
        IP = input()
        data["IP"] = IP
        print(f"IP del dominio: {IP}")

        print("TTL del dominio: ")
        TTL = input()
        data["TTL"] = TTL
        print(f"TTL: {TTL}")

        print("Tipo de dominio: ")
        Type = input()
        data["Type"] = Type
        print(f"Tipo: {data["Type"]}")

        print(f"Resumen: {data}")
        # print(f"Resumen:\n Nombre del dominio: {domainName}\n IP del dominio: {IP}\n TTL: {TTL}\n Tipo: {Type}")

        json_data = json.dumps(data)

        # Realizar la conexión al servidor a partir de acá.
        clientSocket.sendto(json_data.encode(), (serverAddr,serverPort))
        print("Programa Finalizado")

    elif option == "2":
        data["option"] = "2"
        print("Ingresa el nombre del dominio para consultar:")
        data["domainName"] = input()

        print(f"Resumen: {data}")
        
        json_data = json.dumps(data)


        # Conexión al servidor.
        clientSocket.sendto(json_data.encode(), (serverAddr,serverPort))
        print("Esperando respuesta del servidor")
        msg, addr = clientSocket.recvfrom(1024)
        print("Respuesta del servidor:", msg.decode()) 
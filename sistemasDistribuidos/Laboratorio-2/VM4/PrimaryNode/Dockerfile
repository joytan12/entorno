# Etapa 1: Compilación
FROM golang:1.23.1 AS build

WORKDIR /app

# Actualizar repositorios, paquetes e instalar dependencias necesarias
RUN apt-get update && apt-get install -y \
    git \
    bash \
    curl \
    protobuf-compiler

# Instalar los plugins para gRPC y Go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

# Añadir el directorio binario de Go al PATH
ENV PATH="/go/bin:${PATH}"

# Copiar los archivos del proyecto
COPY . .

# Obtener dependencias y generar código 
RUN mkdir proto
RUN touch INFO.txt 
RUN go mod init PrimaryNode
RUN go mod tidy
RUN protoc --go_out=. --go-grpc_out=. general.proto

# Compilar el binario
RUN go build -o main .

# Exponer el puerto en el que el servidor escucha
EXPOSE 8080

# Comando para ejecutar el servidor
CMD ["./main"]

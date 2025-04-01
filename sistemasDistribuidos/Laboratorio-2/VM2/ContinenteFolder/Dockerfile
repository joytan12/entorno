# Usar una imagen base de Go
FROM golang:1.23.1 AS builder

# Actualizar los repositorios y paquetes, e instalar las dependencias
RUN apt-get update && apt-get install -y \
    git \
    bash \
    curl \
    protobuf-compiler \
    build-essential

# Establecer el directorio de trabajo
WORKDIR /app

# Instalar los plugins para gRPC y Go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

# Añadir el directorio binario de Go al PATH
ENV PATH="/go/bin:${PATH}"

# Copiar archivos .proto y generar código gRPC
COPY . .
RUN mkdir proto
RUN go mod init ContinenteFolder
RUN go mod tidy
RUN protoc --go_out=. --go-grpc_out=. general.proto

# Construir el binario del servidor
RUN go build -o main .

# Exponer el puerto en el que el servidor escucha
EXPOSE 8085

# Comando para ejecutar el servidor
CMD ["./main"]
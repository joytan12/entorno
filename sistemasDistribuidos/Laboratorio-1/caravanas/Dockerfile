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
RUN protoc --go_out=. --go-grpc_out=. general.proto

# Construir el binario del servidor de caravanas
RUN go build -o caravanas .

# Usar una imagen mínima para la ejecución
# FROM alpine:latest
# WORKDIR /root/

# # Copiar el binario generado
# COPY --from=builder /app/caravanas .

# Exponer el puerto en el que el servidor de caravanas escucha
EXPOSE 8081

# Comando para ejecutar el servidor
CMD ["./caravanas"]

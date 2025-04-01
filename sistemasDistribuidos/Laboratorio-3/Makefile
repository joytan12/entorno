# Iniciando el sistema Broker
docker-Broker:
	@echo "Iniciando el sistema Broker..."
	sudo docker compose up --build broker

# Iniciando el sistema Jayce
docker-Jayce:
	@echo "Iniciando el sistema Jayce..."
	sudo docker compose up --build jayce

# Iniciando el sistema Servidor 1
docker-Servidor1:
	@echo "Iniciando el sistema Servidor 1..."
	sudo docker compose up --build servidor1

# Iniciando el sistema Servidor 2
docker-Servidor2:
	@echo "Iniciando el sistema Servidor 2..."
	sudo docker compose up --build servidor2

# Iniciando el sistema Servidor 3
docker-Servidor3:
	@echo "Iniciando el sistema Servidor 3..."
	sudo docker compose up --build servidor3

# Iniciando el sistema Supervisor 1
docker-Supervisor1:
	@echo "Iniciando el sistema Supervisor 1..."
	sudo docker compose up --build supervisor1

# Iniciando el sistema Supervisor 2
docker-Supervisor2:
	@echo "Iniciando el sistema Supervisor 2..."
	sudo docker compose up --build supervisor2

# Iniciar el sistema VM1 (Servidor 1 y Supervisor 1)
docker-VM1:
	@echo "Iniciando los sistemas Servidor 1 y Supervisor 1..."
	sudo docker compose up --build servidor1 supervisor1

# Iniciar el sistema VM2 (Servidor 2 y Jayce)
docker-VM2:
	@echo "Iniciando los sistemas Servidor 2 y Jayce..."
	sudo docker compose up --build servidor2 jayce

# Iniciar el sistema VM3 (Servidor 3 y Supervisor 2)
docker-VM3:
	@echo "Iniciando los sistemas Servidor 3 y Supervisor 2..."
	sudo docker compose up --build servidor3 supervisor2

# Iniciar el sistema VM4 (Broker)
docker-VM4:
	@echo "Iniciando el sistema Broker..."
	sudo docker compose up --build broker

docker-down:
	@echo "Cerrando los contenedores ya existentes..."
	sudo docker compose down

docker-Todo:
	@echo "Ejecutando todo lol..."
	sudo docker compose up --build

# Ruta a los archivos txt constantes
CONSTANT_FILES_DIR=./ArchivosConstantes/txt/
# Directorios de las VMs
VM_DIRS=./VM1 ./VM2 ./VM3 ./VM4
# Ruta al archivo proto
PROTO_FILE_DIR=./ArchivosConstantes/proto/
# Directorios de los servicios
SERVICES_DIRS=./VM1/Servidor1 ./VM1/Supervisor1 ./VM2/Servidor2 ./VM2/Jayce ./VM3/Servidor3 ./VM3/Supervisor2 ./VM4/Broker

# Copia los proto a las carpetas correspondientes.
copiar-archivos:
	@echo "Copiando archivos constantes a las VMs..."
	@for dir in $(SERVICES_DIRS); do \
		cp -r $(PROTO_FILE_DIR)* $$dir/; \
	done
	@echo "Archivos copiados correctamente."

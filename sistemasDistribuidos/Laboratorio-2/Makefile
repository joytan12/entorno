# Iniciando el sistema Data Node 1
docker-DataNode1:
	@echo "Iniciando el sistema Data Node 1..."
	sudo docker compose up --build datanode1

# Iniciando el sistema Isla File
docker-IslaFile:
	@echo "Iniciando el sistema Isla File..."
	sudo docker compose up --build islafile

# Iniciando el sistema Continente Folder
docker-ContinenteFolder:
	@echo "Iniciando el sistema Continente Folder..."
	sudo docker compose up --build continentefolder

# Iniciando el sistema Diaboromon
docker-Diaboromon:
	@echo "Iniciando el sistema Diaboromon..."
	sudo docker compose up --build diaboromon

# Iniciando el sistema continente Server
docker-ContinenteServer:
	@echo "Iniciando el sistema Continente Server..."
	sudo docker compose up --build continenteserver

# Iniciando el sistema Data Node 2
docker-DataNode2:
	@echo "Iniciando el sistema Data Node 2..."
	sudo docker compose up --build datanode2

# Iniciando el sistema NodeTai
docker-NodeTai:
	@echo "Iniciando el sistema Node Tai..."
	sudo docker compose run --build -p 8086:8086 --rm --name nodetai nodetai

# Iniciando el sistema Primary Node
docker-PrimaryNode:
	@echo "Iniciando el sistema Primary Node..."
	sudo docker compose up --build primarynode

# Iniciar el sistema VM1 (DataNode1 e IslaFile)
docker-VM1:
	@echo "Iniciando los sistemas DataNode1 e IslaFile..."
	sudo docker compose up --build datanode1 islafile

#Iniciar el sistema VM2 (Continente Folder y Diaboromon)
docker-VM2:
	@echo "Iniciando los sistemas Continente Folder y Diaboromon..."
	sudo docker compose up --build continentefolder diaboromon

#Iniciar el sistema VM3 (Continente Server y Data Node 2)
docker-VM3:
	@echo "Iniciando los sistemas Continente Server y Data Node 2..."
	sudo docker compose up --build continenteserver datanode2

#Iniciar el sistema VM4 (Primary Node y Node Tai)
docker-VM4:
	@echo "Iniciando los sistemas Primary Node y Node Tai..."
	sudo docker compose up --build primarynode nodetai

# Ruta a los archivos txt constantes
CONSTANT_FILES_DIR=./ArchivosConstantes/txt/
# Directorios de las VMs
VM_DIRS=./VM1 ./VM2 ./VM3 ./VM4
# Ruta al archivo proto
PROTO_FILE_DIR=./ArchivosConstantes/proto/
# Directorios de los servicios
SERVICES_DIRS=./VM1/IslaFile ./VM1/DataNode1 ./VM2/ContinenteFolder ./VM2/Diaboromon ./VM3/ContinenteServer ./VM3/DataNode2 ./VM4/PrimaryNode ./VM4/NodeTai

# Copia los txt y los proto, sin embargo TODO: Quitar los txt de dentro, ajustar docker-compose para ello.
copiar-archivos:
	@echo "Copiando archivos constantes a las VMs..."
	@for dir in $(SERVICES_DIRS); do \
		cp -r $(PROTO_FILE_DIR)* $$dir/; \
	done
	@echo "Archivos copiados correctamente."

primary_data:
	sudo docker compose up --build primarynode datanode1

todos:
	sudo docker compose up --build
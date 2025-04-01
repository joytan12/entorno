➤ INTEGRANTES

Juan García 202073539-5
Vicente Zapata 202073573-5

➤ INSTRUCCIONES DE EJECUCIÓN
En cada máquina virtual, debes seguir estos pasos:

1.- Accede a la carpeta del proyecto:

cd Laboratorio_1

2.- Ejecuta el make correspondiente a cada máquina:

* Para el servidor principal (Konzu):
make docker-logistica

* Para el cliente (Facciones):
make docker-clientes

* Para las caravanas:
make docker-caravanas

* Para el servidor financiero (Raquis):
make docker-finanzas

3.- Disfruta.


➤ OBSERVACIONES
- La Branch que tiene los archivos Finales de nuestra entrega es "main".
- El archivo input.txt es el encargado de leer las órdenes que dejarán los clientes, por lo que este archivo va en la carpeta client. Este archivo tiene el siguiente formato:
IDpaquete,Tipo,Nombre,Valor,Escolta,Destino, por ejemplo, una orden posible de colocar en input.txt es: 0002,Ostronitas,Pocion,500,Escolta4,Destino2
- Para facilitar el rumbo del programa, se dejaron solamente como opciones posibles para "Destino" las siguientes, "Destino1", "Destino2", "Destino3" y "Destino4, cada una con tiempos de viaje
distintos, los cuales tienen tiempos de 3, 4, 5 y 6 segundos respectivamente. (Se pueden alterar al principio del archivo caravanas.go)
- Cabe destacar, que una vez ejecutado el programa, se colocaron una serie de time.Sleep(), para
darle tiempo a conectarse los servidores entre ellos. El que más se demora en iniciar es el client,
y es quién da inicio a todas las funcionalidades del sistema. Se hizo esto porque no daba tiempo a
conectarse con RabbitMQ, ocasionando problemas.
- Las máquinas virtuales asignadas para nuestro grupo fueron: 
dist041: server
dist042: client 
dist043: caravanas
dist044: financiero
En estas máquinas se deben ejecutar cada uno de los servidores concretos, dado que las ip's fueron puestas dependiendo la máquina.
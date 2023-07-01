from threading import Thread, Semaphore, Lock
import datetime, time, random

sem_rusa_capacidad = Semaphore(10)
sem_rusa_fila = Semaphore(10)

sem_terror_capacidad = Semaphore(2)
sem_terror_fila = Semaphore(8)

sem_carrusel_capacidad = Semaphore(5)
sem_carrusel_fila = Semaphore(15)

sem_barco_capacidad = Semaphore(3)
sem_barco_fila = Semaphore(6)

lock_zona_comun = Lock()
lock_rusa = Lock()
lock_terror = Lock()
lock_carrusel = Lock()
lock_barco = Lock()
lock_txt = Lock()

rusa_enjuego = 0
rusa_encola = 0

terror_enjuego = 0
terror_encola  = 0

carrusel_enjuego = 0
carrusel_encola = 0

barco_enjuego = 0
barco_encola = 0

total_personas = 150

def str_Zona_Comun(thr_num, juego_elegido1, juego_elegido2, time_start, time_inqueue1, time_inqueue2):
    juegos = ["Montaña Rusa", "Casa del Terror", "Carrusel", "Barco Pirata"]
    return "Persona{} {}, {}, {}, {}, {}\n".format(thr_num, time_start, juegos[juego_elegido1], time_inqueue1, juegos[juego_elegido2], time_inqueue2)

class Jugador(Thread):
    def __init__(self, thread_num):
        Thread.__init__(self)
        self.thread_num = thread_num+1

    def run(self): #Zona_Comun

        global total_personas
        global rusa_enjuego, terror_enjuego, carrusel_enjuego, barco_enjuego
        global rusa_encola, terror_encola, carrusel_encola, barco_encola

        cant_juegos = 0
        juego_elegido1 = -1
        juego_elegido2 = -1

        time_start = str(datetime.datetime.now().strftime("%H:%M:%S.%f"))

        time_inqueue1 = -1
        time_inqueue2 = -1

        while cant_juegos!=2:

            juego_elegido = random.randint(0,3)

            if (juego_elegido==0): #Montaña rusa
                sem_rusa_fila.acquire() #Entra a la cola

                lock_zona_comun.acquire()

                if cant_juegos==0:
                    juego_elegido1 = juego_elegido
                    time_inqueue1 = str(datetime.datetime.now().strftime("%H:%M:%S.%f"))
                else:
                    juego_elegido2 = juego_elegido
                    time_inqueue2 = str(datetime.datetime.now().strftime("%H:%M:%S.%f"))

                file = open("Zona_Comun.txt","a")

                if cant_juegos!=0:
                    file.write(str_Zona_Comun(str(self.thread_num), juego_elegido1, juego_elegido2, time_start, time_inqueue1, time_inqueue2))

                file.close()
                total_personas -= 1
                lock_zona_comun.release()

                lock_rusa.acquire()
                rusa_encola += 1
                lock_rusa.release()

                sem_rusa_capacidad.acquire()#Entra al juego

                lock_rusa.acquire()
                rusa_encola -= 1
                lock_rusa.release()

                sem_rusa_fila.release()#Sale de la cola

                time_entrarpartida = str(datetime.datetime.now().strftime("%H:%M:%S.%f"))

                lock_rusa.acquire()
                file = open("Montaña_Rusa.txt","a")
                if cant_juegos == 0:
                    file.write("Persona"+str(self.thread_num)+", "+time_inqueue1+", "+time_entrarpartida+"\n")
                else:
                    file.write("Persona"+str(self.thread_num)+", "+time_inqueue2+", "+time_entrarpartida+"\n")
                file.close()
                rusa_enjuego += 1
                lock_rusa.release()

                while rusa_enjuego < 10 and total_personas > (rusa_encola+terror_encola+carrusel_encola+barco_encola):
                    pass
                time.sleep(9)

                lock_rusa.acquire()
                rusa_enjuego -= 1
                lock_rusa.release()

                sem_rusa_capacidad.release()#Sale del juego


            elif(juego_elegido==1): #Casa del terror
                sem_terror_fila.acquire() #Entra a la cola

                lock_zona_comun.acquire()

                if cant_juegos==0:
                    juego_elegido1 = juego_elegido
                    time_inqueue1 = str(datetime.datetime.now().strftime("%H:%M:%S.%f"))
                else:
                    juego_elegido2 = juego_elegido
                    time_inqueue2 = str(datetime.datetime.now().strftime("%H:%M:%S.%f"))

                file = open("Zona_Comun.txt","a")

                if cant_juegos!=0:
                    file.write(str_Zona_Comun(str(self.thread_num), juego_elegido1, juego_elegido2, time_start, time_inqueue1, time_inqueue2))

                file.close()
                total_personas -= 1
                lock_zona_comun.release()

                lock_terror.acquire()
                terror_encola += 1
                lock_terror.release()

                sem_terror_capacidad.acquire()#Entra al juego

                lock_terror.acquire()
                terror_encola -= 1
                lock_terror.release()

                sem_terror_fila.release()#Sale de la cola

                time_entrarpartida = str(datetime.datetime.now().strftime("%H:%M:%S.%f"))

                lock_terror.acquire()
                file = open("Casa_Terror.txt","a")
                if cant_juegos == 0:
                    file.write("Persona"+str(self.thread_num)+", "+time_inqueue1+", "+time_entrarpartida+"\n")
                else:
                    file.write("Persona"+str(self.thread_num)+", "+time_inqueue2+", "+time_entrarpartida+"\n")
                file.close()
                terror_enjuego += 1
                lock_terror.release()

                while terror_enjuego < 2 and total_personas > (rusa_encola+terror_encola+carrusel_encola+barco_encola):
                    pass
                time.sleep(5)

                lock_terror.acquire()
                terror_enjuego -= 1
                lock_terror.release()

                sem_terror_capacidad.release()#Sale del juego


            elif (juego_elegido==2): #Carrusel
                sem_carrusel_fila.acquire() #Entra a la cola

                lock_zona_comun.acquire()

                if cant_juegos==0:
                    juego_elegido1 = juego_elegido
                    time_inqueue1 = str(datetime.datetime.now().strftime("%H:%M:%S.%f"))
                else:
                    juego_elegido2 = juego_elegido
                    time_inqueue2 = str(datetime.datetime.now().strftime("%H:%M:%S.%f"))

                file = open("Zona_Comun.txt","a")

                if cant_juegos!=0:
                    file.write(str_Zona_Comun(str(self.thread_num), juego_elegido1, juego_elegido2, time_start, time_inqueue1, time_inqueue2))

                file.close()
                total_personas -= 1
                lock_zona_comun.release()

                lock_carrusel.acquire()
                carrusel_encola += 1
                lock_carrusel.release()

                sem_carrusel_capacidad.acquire()#Entra al juego

                lock_carrusel.acquire()
                carrusel_encola -= 1
                lock_carrusel.release()

                sem_carrusel_fila.release()#Sale de la cola

                time_entrarpartida = str(datetime.datetime.now().strftime("%H:%M:%S.%f"))

                lock_carrusel.acquire()
                file = open("Carrusel.txt","a")
                if cant_juegos == 0:
                    file.write("Persona"+str(self.thread_num)+", "+time_inqueue1+", "+time_entrarpartida+"\n")
                else:
                    file.write("Persona"+str(self.thread_num)+", "+time_inqueue2+", "+time_entrarpartida+"\n")
                file.close()
                carrusel_enjuego += 1
                lock_carrusel.release()

                while carrusel_enjuego < 5 and total_personas > (rusa_encola+terror_encola+carrusel_encola+barco_encola):
                    pass
                time.sleep(7)

                lock_carrusel.acquire()
                carrusel_enjuego -= 1
                lock_carrusel.release()

                sem_carrusel_capacidad.release()#Sale del juego


            else: #Barco Pirata
                sem_barco_fila.acquire() #Entra a la cola

                lock_zona_comun.acquire()

                if cant_juegos==0:
                    juego_elegido1 = juego_elegido
                    time_inqueue1 = str(datetime.datetime.now().strftime("%H:%M:%S.%f"))
                else:
                    juego_elegido2 = juego_elegido
                    time_inqueue2 = str(datetime.datetime.now().strftime("%H:%M:%S.%f"))

                file = open("Zona_Comun.txt","a")

                if cant_juegos!=0:
                    file.write(str_Zona_Comun(str(self.thread_num), juego_elegido1, juego_elegido2, time_start, time_inqueue1, time_inqueue2))

                file.close()
                total_personas -= 1
                lock_zona_comun.release()

                lock_barco.acquire()
                barco_encola += 1
                lock_barco.release()

                sem_barco_capacidad.acquire()#Entra al juego

                lock_barco.acquire()
                barco_encola -= 1
                lock_barco.release()

                sem_barco_fila.release()#Sale de la cola

                time_entrarpartida = str(datetime.datetime.now().strftime("%H:%M:%S.%f"))

                lock_barco.acquire()
                file = open("Barco_Pirata.txt","a")
                if cant_juegos == 0:
                    file.write("Persona"+str(self.thread_num)+", "+time_inqueue1+", "+time_entrarpartida+"\n")
                else:
                    file.write("Persona"+str(self.thread_num)+", "+time_inqueue2+", "+time_entrarpartida+"\n")
                file.close()
                barco_enjuego += 1
                lock_barco.release()

                while barco_enjuego < 3 and total_personas > (rusa_encola+terror_encola+carrusel_encola+barco_encola):
                    pass
                time.sleep(4)

                lock_barco.acquire()
                carrusel_enjuego -= 1
                lock_barco.release()

                sem_barco_capacidad.release()#Sale del juego


            cant_juegos+=1

        time_salir = str(datetime.datetime.now().strftime("%H:%M:%S.%f"))
        lock_zona_comun.acquire()
        file = open("Salida.txt","a")
        file.write("Persona"+str(self.thread_num)+", "+time_salir+"\n")
        file.close()

        print("Persona"+str(self.thread_num)+", "+time_salir) #UNICO PRINT

        lock_zona_comun.release()

### MAIN ###

#CREACION ARCHIVOS
file = open("Zona_Comun.txt","w")
file.close()
file = open("Montaña_Rusa.txt","w")
file.close()
file = open("Casa_Terror.txt","w")
file.close()
file = open("Carrusel.txt","w")
file.close()
file = open("Barco_Pirata.txt","w")
file.close()
file = open("Salida.txt","w")
file.close()

#THREADS
random.seed()
for i in range(total_personas):
    hilo=Jugador(i)
    hilo.start()
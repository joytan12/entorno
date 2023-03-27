#include <stdio.h>
#include <sys/stat.h>
#include <dirent.h>
#include <string.h>
#include <unistd.h>

// CraerSecciones void => void
// crea las 3 secciones para casa genero que se va agregando segun la ruta del archivo que se pase
void CrearSecciones(char* rutaCompleta) {
  char rutaFinal[256];
  char *carpeta1 = "/menos_a_4000", *carpeta2 = "/entre_4000_y_8000", *carpeta3 = "/Mayor_a_8000";

  sprintf(rutaFinal, "%s%s", rutaCompleta, carpeta1);
  mkdir(rutaFinal, 0700);
  sprintf(rutaFinal, "%s%s", rutaCompleta, carpeta2);
  mkdir(rutaFinal, 0700);
  sprintf(rutaFinal, "%s%s", rutaCompleta, carpeta3);
  mkdir(rutaFinal, 0700);
} 

//esta funcione se dedica a llevar el archivo del juego de una rutaInicial a una rutaFinal,
//mediante la conparacion de sus jugadores, se construye la rutaFinal
void moverArchivo(int jugadores, char* rutaInicial, char* rutaIncompleta){
  DIR* directorio = opendir(rutaIncompleta);
  struct dirent* archivo;
  int cont = 0;
  char rutaFinal[256];

  while ((archivo = readdir(directorio)) != NULL) {
    sprintf(rutaFinal, "%s%s", rutaIncompleta, archivo);
    printf("%s\n", rutaFinal);
    printf("%s\n", rutaInicial);
    if (jugadores < 4000 && cont == 4){
      rename(rutaInicial, rutaFinal);
      closedir(directorio);
      return;
    } else if (4000 <= jugadores && jugadores <= 8000 && cont == 4) {
      rename(rutaInicial, rutaFinal);
      closedir(directorio);
      return;
    } else if (cont == 4) {
      rename(rutaInicial, rutaFinal);
    }
    cont++;
  }
  closedir(directorio);
}

 int CraerGenoro() {  
  //aqui se alige como se quiere ordenar las cossa  
  int opcion;
  printf("ordenas por\n(1) jugarores actuales\n(2) pico de jugadores\n");
  scanf("%d", &opcion);
  if (access("CWD", F_OK) != -1){
    printf("El archivo ya esta creado\n");
  } else {
    mkdir("CWD", 0700);
  }

  //ibrir directorio
  char filenema1[256], filenema2[256]; //filenema1 tiene la ruta en donde se encuentra el archivo
  DIR* directorio;
  char* extencion = ".txt";
  struct dirent* archivo;
  char rutaRelativa[256] = "../tarea1/CWD/", rutaArchivos[256] = "../tarea1/Archivos_Prueba/";

  //abrir el directorio
  directorio = opendir("../tarea1/Archivos_Prueba/");

  // Verificar que se abrio correctamente
  if (directorio == NULL) {
    printf("No se puede abrir el directorio.\n");
    return 1;
  }
 
  while ((archivo = readdir(directorio)) != NULL) {
    if (strstr(archivo->d_name, extencion) != NULL){
      FILE* juego;
      char linea[256];
      int jugadores;
      sprintf(filenema1, "%s%s", rutaArchivos, archivo->d_name);
      juego = fopen(filenema1, "r");
      int cont = 0;
      while (fgets(linea, sizeof(linea), juego)){
        sprintf(filenema2, "%s%s", rutaRelativa, linea);
        if (cont == 0 && opcion == 1){
          jugadores = atoi(linea);
        } else if (cont == 1 && opcion == 2){
          jugadores = atoi(linea);
        } else if(cont == 2){
          printf("%s\n", linea);
          snprintf(filenema2, sizeof(filenema2), "%s%s", rutaRelativa, strtok(linea, "\r\n"));
          if (access(filenema2, F_OK) == -1){
            mkdir(filenema2, 0700);
            CrearSecciones(filenema2);  
          }
          moverArchivo(jugadores, filenema1, filenema2);
        }
        cont++;
      }
      fclose(juego);
    }
  } 
  
  closedir(directorio);
  return 0;
}

int main () {
  CraerGenoro();
  printf("Hola mundo\n");
  return 0;
}

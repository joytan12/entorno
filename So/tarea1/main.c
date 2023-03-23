#include <stdio.h>
#include <sys/stat.h>
#include <dirent.h>
#include <string.h>
#include <unistd.h>

// CraerSecciones void => void
// crea las 3 secciones para casa genero que se va agregando 
void CrearSecciones(char rutaIncompleta) {
  struct dirent* archivo;
  DIR* directorio = opendir(rutaInicial);
  int cont = 0;
  char rutaFinal[256];
  char *carpeta1 = "menos_a_4000", *carpeta2 = "entre_4000_y_8000", *carpeta3 = "Mayor_a_8000";

  while ((archivo = readdir(directorio)) != NULL){
    sprintf(rutaFinal, "%s%s", rutaIncompleta, );
    if (cont == 0){
      mkdir(rutaFinal, 0700);
    } else if (cont == 1){      
      mkdir(rutaFinal, 0700);
    } else if (cont == 2) {
      rename(rutaInicial, rutaFinal);
    }
    cont++;
  }
} 

//esta funcione se de
void moverArchivo(int jugadores, char rutaInicial, char rutaIncompleta){
  DIR* directorio = opendir(rutaInicial);
  struct dirent* archivo;
  int cont = 0;
  char rutaFinal[256];
  while ((archivo = readdir(directorio)) != NULL) {
    sprintf(rutaFinal, "%s%s", rutaIncompleta, archivo);
    if (jugarores < 4000 && cont == 0){
      rename(rutaInicial, rutaFinal);
    } else if (4000 <= jugadores && jugarores <= 8000 && cont == 1) {
      rename(rutaInicial, rutaFinal);
    } else if (cont == 2) {
      rename(rutaInicial, rutaFinal);
    }
    cont++;
  }
  closedir(directorio)
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
  char filenema1[256], filenema2[256];
  DIR* directorio;
  char* extencion = ".txt";
  struct dirent* archivo;
  char rutaRelativa[100] = "../tarea1/CWD/", rutaArchivos[100] = "../tarea1/Archivos_Prueba/";

  //abrir el directorio
  directorio = opendir("../tarea1/Archivos_Prueba/");

  // Verificar que se abrio correctamente
  if (directorio == NULL) {
  printf("No se puede abrir el directorio.\n");
    return 1;
  }

  while ((archivo = readdir(directorio)) != NULL) {
    if (strstr(archivo->d_name, extencion) != NULL){
      printf("entro a ver archivos de juegos\n");
      FILE* juego;
      char linea[50];
      int jugarores;
      sprintf(filenema1, "%s%s", rutaArchivos, archivo->d_name);
      juego = fopen(filenema1, "r");
      int cont = 0;
      while (fgets(linea, sizeof(linea), juego)){
        sprintf(filenema2, "%s%s", rutaRelativa, archivo->d_name);
        if (cont == 0 && opcion == 1){
          jugarores = atoi(linea);
        } else if (cont == 1 && opcion == 2){
          jugarores = atoi(linea);
        } else if(cont == 3){
          //se tiene que revisar las direcciones para los archivos
          if (access(filenema2, F_OK) != -1){
            moverArchivo(jugarores, filenema1, filenema2);
          } else {
            mkdir(filenema2, 0700);
            rename(filenema1, filenema2);
          }
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
  //CrearSecciones();
  printf("entro al programa\n");
  CraerGenoro();
  printf("Hola mundo\n");
  return 0;
}

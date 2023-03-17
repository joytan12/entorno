#include <stdio.h>
#include <sys/stat.h>
#include <dirent.h>

// CraerSecciones void => void
// crea las 3 secciones para casa genero que se va agregando 
void CrearSecciones() {
  char *carpeta1 = "menos_a_4000", *carpeta2 = "entre_4000_y_8000", *carpeta3 = "Mayor_a_8000";
  int maldito = mkdir(carpeta1, 0700);
  mkdir(carpeta2, 0700);
  mkdir(carpeta3, 0700);
  printf("%d\n", maldito);
}

int CraerGenoro() {
  mkdir("CWD", 0700)
  DIR* directorio;
  char* extencion = ".txt";
  struct dirent* archivo;

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
      char linea[100];

      juego = fopen(archivo->d_name, "r");
      int cont = 0;
      char* genero, picoJugadores, actualJugadores;

      while (fgets(linea, sizeof(linea), juego)){
        if (cont == 0){
          picoJugadores = linea;
        } else if (cont == 1) {
          actualJugadores = linea;
        } else {
          genero = linea;
        }
        cont++;
      }
      mkdir("./CWM/" + genero, 0700);
      //aqui deberia moverse el archivo de un lado a otro
    }
  }

  closedir(directorio);
  return 0;
}

int main () {
  //CrearSecciones();
  CraerGenoro();
  printf("Hola munao\n");
  return 0;
}

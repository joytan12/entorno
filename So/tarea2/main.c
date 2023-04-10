#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#define NUM_CARTAS 44


void generarMazo(int *cartas){
    int i, temp, indice_aleatorio;
    // Llenar el arreglo con los números del 1 al 44
    for (i = 0; i < NUM_CARTAS; i++) {
        cartas[i] = i + 1;
    }
    
    // Desordenar el mazo
    srand(time(NULL)); // Inicializar la semilla del generador de números aleatorios
    for (i = 0; i < NUM_CARTAS; i++) {
        indice_aleatorio = rand() % NUM_CARTAS;
        temp = cartas[i];
        cartas[i] = cartas[indice_aleatorio];
        cartas[indice_aleatorio] = temp;
    }
}

int main () {
    //generar mazo
    int cartas[NUM_CARTAS];
    int i;

    generarMazo(&cartas);
    printf("Mazo desordenado:\n");
    for (i = 0; i < NUM_CARTAS; i++) {
        printf("%d ", cartas[i]);
    }
    printf("\n");

    pid_t pid;
    int status;
    int num_hijos = 4;
    int variable = 0;
    
    for (i = 1; i <= num_hijos; i++) {
        pid = fork();
        if (pid == -1) {
            printf("Error al crear hijo.\n");
            return 1;
        } else if (pid == 0) {
            // Código de cada hijo
            while (variable < 10) {
                printf("Hijo %d, variable = %d\n", i, variable);
                variable++;
            }
            return 0;
        }
    }
    
    // Código del padre
    for (i = 1; i <= num_hijos; i++) {
        wait(&status);
        if (WIFEXITED(status)) {
            printf("El hijo %d ha terminado con status %d.\n", i, WEXITSTATUS(status));
        }
    }
    
    printf("Variable final: %d\n", variable);
    return 0;
}
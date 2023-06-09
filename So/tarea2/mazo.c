#include "mazo.h"

void generarMazo(int *cartas){
    int i, temp, indice_aleatorio;
    // Llenar el arreglo con los números del 1 al 44
    for (i = 0; i < NUM_CARTA; i++) {
        cartas[i] = i + 1;
    }
    
    // Desordenar el mazo
    srand(time(NULL)); // Inicializar la semilla del generador de números aleatorios
    for (i = 0; i < NUM_CARTA; i++) {
        indice_aleatorio = rand() % NUM_CARTA;
        temp = cartas[i];
        cartas[i] = cartas[indice_aleatorio];
        cartas[indice_aleatorio] = temp;
    }
}

void repartirMano(Tablero *tablero, Mano *manos, int mazo[]){
    tablero[0].pila[0] = mazo[40];
    tablero[0].top = mazo[40];
    tablero[0].cartas = 1;
    tablero[1].pila[0] = mazo[41];
    tablero[1].top = mazo[41];
    tablero[1].cartas = 1;
    tablero[2].pila[0] = mazo[42];
    tablero[2].top = mazo[42];
    tablero[2].cartas = 1;
    tablero[3].pila[0] = mazo[43];
    tablero[3].top = mazo[43];
    tablero[3].cartas = 1;
    for (int i = 0; i < NUM_HIJOS; i++){
        for (int j = 0; j < 10; j++){
            manos[i].mano[j] = mazo[(i * 10) + j];
        }
        manos[i].puntaje = 0;
    }
}

int jugarCarta(int carta, Tablero *tablero, int nPila){
    if (tablero[nPila].cartas == 5 || tablero[nPila].pila[tablero->cartas] > carta){
        int puntos = 0;
        for (int i = 0; i < tablero[nPila].cartas; i++)
        {
            if (tablero[nPila].pila[i]%11 == 0) {
                puntos += 5;
            } else if (tablero[nPila].pila[i]%10 == 0) {
                puntos += 3;
            } else if (tablero[nPila].pila[i]%5 == 0) {
                puntos += 2;
            } else {
                puntos += 1;
            }
        }
        tablero[nPila].pila[0] = carta;
        tablero[nPila].cartas == 1;
        return puntos;
    } else {
        tablero[nPila].pila[tablero[nPila].cartas] = carta;
        tablero[nPila].cartas++;
        return 0;
    }
}
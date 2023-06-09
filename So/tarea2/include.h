#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <signal.h>
#include <time.h>

#define NUM_CARTA 44
#define NUM_HIJOS 4
#define BUFFER_SIZE sizeof(int)

#ifndef structs
#define structs

typedef struct{
    int mano[10];
    int puntaje;
} Mano;


typedef struct{
    int pila[NUM_CARTA];
    int top;
    int cartas;
} Tablero;

#endif
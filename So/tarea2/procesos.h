#include "include.h"

void cicloJugador(int *pipe, Tablero *tablero, Mano *mano);

void cicloBot(int *pipe, Tablero *tablero, Mano *mano, int jugador);

void cicloPadre(Tablero *tablero, Mano *manos, int *pipe1, int *pipe2, int *pipe3, int *pipe4);

void cicloPadre1(Tablero *tablero, Mano *manos, int *pipe1);

long int ordenar(int *jugada, int estado, Tablero *tablero);

void actualizar(int intruccion, Tablero *tablero, Mano *mano);

void bubbleSort(int *arr, int n);

void sacarCarta(Mano *mano, int jugador, int carta);

void verMano(Mano *manos, int indice);

void verFilas(Tablero *tablero);

int eligeJugador(Mano *manos, Tablero *tablero, int indice, int *arr);

int eligeBot(Mano *manos, int indice, int menor, int *arr);
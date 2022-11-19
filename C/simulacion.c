#include <stdio.h>
#include <stdlib.h>
#include "animal.h"
#define M 1000

int main(){
    animal** mundo;
    mundo = (animal**)malloc(M*sizeof(animal*));
    for(int i=0; i<M; i++){
        mundo[i] = (animal*)malloc(M*sizeof(animal));
    }
    int n;
    do {
        printf("elija la accion que quiera hacer\nsi quieres crear un animal elija 1\nsi quieres ver el mundo 2\nsi quieres avanzar una iteracion de tiempo 3\nsi quieres terminar el programa precione -1\n");
        scanf(" %d", &n);
        switch (n){
            int q, y;
            case 1:
                printf("las cordenadas en donde va a poner el animal seran:\n");
                do{
                    scanf(" %i", &q);
                    scanf(" %i", &y);
                    if (mundo[q][y].fuerza != NULL) printf("ingresa de nuevo las cordenadas");
                } while(mundo[q][y].fuerza != NULL);
                crearAnimal(&mundo[q][y]);
                break;
            case 2:
                mostrarMundo(mundo);
                break;
            case 3:
                // iteracion(mundo);
                printf("deveria iterar pero no lo hace");
                break;
            default:
                break;
        }
    } while (n > 0);
    borrarMundo(mundo);
    for(int i=0; i<M; i++){
        free(mundo[i]);
    }
    free(mundo);
    return 0;
}
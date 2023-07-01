#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

//Inicializaciones
void moverElementos(int *arreglo, int indice1, int indice2);
int oRecursivo(int arreglo[], int cantidad, int actual, int aux);
int verificarOrden(int *arreglo, int tamano);
void imprimirArreglo(int *arreglo, int tamano);
void swap(int *a, int *b);

void swap(int *a, int *b){
  int temp = *a;
  *a = *b;
  *b = temp;
}

void imprimirArreglo(int *arreglo, int tamano){
    for(int i=0; i <= tamano; i++){
        printf("%d ",arreglo[i]);
    }
    printf("\n");
}

int verificarOrden(int *arreglo, int tamano){
    int verificacion = 0; //0 si esta ordenado, !=0 si no lo esta
    for(int i = 0; i < tamano; i++){
        if (arreglo[i] > arreglo[i+1]){
            verificacion++;
        }
    }
    return verificacion;
}

void moverElementos(int *arreglo, int indice1, int indice2) {
    if (indice1 != indice2){
        int aux = arreglo[indice1];
        if(indice1 >= indice2){
            for (int k = indice1; k >= indice2; k--){
                arreglo[k] = arreglo[k-1];
            }
        } else { //indice1 < indice2
            for(int k = indice1; k < indice2; k++){
                arreglo[k] = arreglo[k+1];
            }
        }
        arreglo[indice2] = aux;
    }
}

int oRecursivo(int arreglo[], int cantidad, int actual, int aux){
    int menor = 0;
    for(int i = actual; i < cantidad; i++){
        printf("%d, %d\n", arreglo[actual] > arreglo[i+1]);
        if (arreglo[actual] > arreglo[i+1]){
            moverElementos(arreglo, actual, i+1);
            menor++;
            break;
        }
    }
    imprimirArreglo(arreglo, cantidad);
    // printf("%d, %d\n", menor, actual);
    if (actual <= cantidad){
        aux = oRecursivo(arreglo, cantidad, actual + 1, aux + menor);
    }
    return aux;
}

int main() {
    int arreglo[] = {1, 8, 9, 2};
    int cantidad = 3; //<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< OJO
    int min;
    min = oRecursivo(arreglo, cantidad, 0, 0);
    printf("min = %d\n", min);
    int indice1 = 0;
    int indice2 = 2;
    // moverElementos(arreglo, indice1, indice2);
    // for(int i=0; i <= cantidad; i++){
    //     printf("%d ",arreglo[i]);
    // }
    // printf("\n");J
    return 0;
}
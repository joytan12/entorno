#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <time.h>
#include "animal.h"
#define M 1000

void crearAnimal(animal* a){
    /*
    en crear animal recibe el puntero a nanimal y se asigna memoria
    o se pide los valores correspondientes
    */
    printf("ingrese el tipo de fuerza:\n");
    scanf(" %c", &a->tipoFuerza);
    printf("ingrese la fuerza:\n");
    switch (a->tipoFuerza){
        case 'e':
            a->fuerza = malloc(sizeof(int));
            scanf(" %i", (int*)a->fuerza);
            break;
        case 'c':
            a->fuerza = malloc(sizeof(char));
            scanf(" %c", (char*)a->fuerza);
            break;
        case 'f':
            a->fuerza = malloc(sizeof(float));
            scanf(" %f", (float*)a->fuerza);
            break;
        default:
            break;
    }
    printf("ingrese el tipo de velocidad:\n");
    scanf(" %c", &a->tipoVelocidad);
    printf("ingrese la velocidad:\n");
    switch (a->tipoVelocidad){
        case 'e':
            a->velocidad = malloc(sizeof(int));
            scanf(" %i", (int*)a->velocidad);
            break;
        case 'c':
            a->velocidad = malloc(sizeof(char));
            scanf(" %c", (char*)a->velocidad);
            break;
        case 'f':
            a->velocidad = malloc(sizeof(float));
            scanf(" %f", (float*)a->velocidad);
            break;
        default:
            break;
    }
    printf("ingrese el tipo de resistencia:\n");
    scanf(" %c", &a->tipoResistecia);
    printf("ingrese la resistencia:\n");
    switch (a->tipoResistecia){
        case 'e':
            a->resistencia = malloc(sizeof(int));
            scanf(" %i", (int*)a->resistencia);
            break;
        case 'c':
            a->resistencia = malloc(sizeof(char));
            scanf(" %c", (char*)a->resistencia);
            break;
        case 'f':
            a->resistencia = malloc(sizeof(float));
            scanf(" %f", (float*)a->resistencia);
            break;
        default:
            break;
    }
    int aux;
    printf("que tipo de reproduccion\nsimple(1)\ncruzada(2)\n");
    scanf(" %d", &aux);
    switch (aux){
    case 1:
        a->reproduccion = &reproduccionSimple;
        break;
    case 2:
        a->reproduccion = &reproduccionCruzada;
        break;
    default:
        break;
    }
    printf("que tipo de comr o huir quiere\ncomerSiempre(1)\nhuirSiempre(2)\ncomerAleatorio(3)\n");
    scanf(" %d", &aux);
    switch (aux){
    case 1:
        a->comerHuir = &comerSiempre;
        break;
    case 2:
        a->comerHuir = &huirSiempre;
        break;
    case 3:
        a->comerHuir = &comerAleatorio;
        break;
    default:
        break;
    }
}

void borrar(animal* a){
    /*libera las memoria del los punteros 
    void de el animal*/
    free(a->fuerza);
    a->fuerza = NULL;
    free(a->velocidad);
    a->velocidad = NULL;
    free(a->resistencia);
    a->resistencia = NULL;
}

void mostrarAnimal(animal* a){
    /*esta funcion se dedica a identificar que tipos de datos 
    se van a printear por pantalla lo los print*/
    if (a->fuerza != NULL){
        printf("fuerza: ");
        if (a->tipoFuerza == 'e'){
            printf("%d\n", *(int*)a->fuerza);
        } else if (a->tipoFuerza == 'c'){
            printf("%c\n", *(char*)a->fuerza);
        } else{
            printf("%f\n", *(float*)a->fuerza);
        }
        printf("velocidad: ");
        if (a->tipoVelocidad == 'e'){
            printf("%d\n", *(int*)a->velocidad);
        } else if (a->tipoVelocidad == 'c'){
            printf("%c\n", *(char*)a->velocidad);
        } else{
            printf("%f\n", *(float*)a->velocidad);
        }
        printf("resistecia: ");
        if (a->tipoResistecia == 'e'){
            printf("%d\n", *(int*)a->resistencia);
        } else if (a->tipoResistecia == 'c'){
            printf("%c\n", *(char*)a->resistencia);
        } else{
            printf("%f\n", *(float*)a->resistencia);
        }
    }
}

void reproduccion(animal* padre, animal* madre, animal* hijo){
    /*esta funcion se dedica a recibir los tres punderos despues compara 
    al puntero padre y madre hereda los atrivutos al puntero hijo*/
    srand(time(NULL));
    int x;
    x = rand()%2;
    if (x == 1){
        padre->reproduccion(padre, madre, hijo);
    } else {
        madre->reproduccion(padre, madre, hijo);
    }
}

void comerOhuir(animal* a1, animal* a2){
    srand(time(NULL));
    int x;
    x = rand()%2;
    if (x == 1){
        a1->comerHuir(a1, a2);
    } else {
        a2->comerHuir(a1, a2);
    }
}

int comparar(animal* a1, animal* a2){
    int aux=0;
    int aux1, aux2; 
    switch (a1->tipoFuerza){
        case 'e':
            aux1 = *(int*)a1->fuerza;
            break;
        case 'c':
            aux1 = (int)(*(char*)(a1->fuerza)/4);
            break;
        case 'f':
            aux1 = roundf((int)*(float*)a1->fuerza);
            break;
        default:
            break;
    }

    switch (a2->tipoFuerza){
        case 'e':
            aux2 = *(int*)a2->fuerza;
            break;
        case 'c':
            aux2 = (int)(*(char*)(a2->fuerza)/4);
            break;
        case 'f':
            aux2 = roundf((int)*(float*)a2->fuerza);
            break;
        default:
            break;
    }
    if (aux2 < aux1){
        aux++;
    }

    switch (a1->tipoVelocidad){
        case 'e':
            aux1 = *(int*)a1->velocidad;
            break;
        case 'c':
            aux1 = (int)(*(char*)(a1->velocidad)/4);
            break;
        case 'f':
            aux1 = roundf((int)*(float*)a1->velocidad);
            break;
        default:
            break;
    }

    switch (a2->tipoVelocidad){
        case 'e':
            aux2 = *(int*)a2->velocidad;
            break;
        case 'c':
            aux2 = (int)(*(char*)(a2->velocidad)/4);
            break;
        case 'f':
            aux2 = roundf((int)*(float*)a2->velocidad);
            break;
        default:
            break;
    }
    if (aux2 < aux1){
        aux++;
    }

    switch (a1->tipoResistecia){
        case 'e':
            aux1 = *(int*)a1->resistencia;
            break;
        case 'c':
            aux1 = (int)(*(char*)(a1->resistencia)/4);
            break;
        case 'f':
            aux1 = roundf((int)*(float*)a1->resistencia);
            break;
        default:
            break;
    }

    switch (a2->tipoResistecia){
        case 'e':
            aux2 = *(int*)a2->resistencia;
            break;
        case 'c':
            aux2 = (int)(*(char*)(a2->resistencia)/4);
            break;
        case 'f':
            aux2 = roundf((int)*(float*)a2->resistencia);
            break;
        default:
            break;
    }
    if (aux2 < aux1){
        aux++;
    }
    if (aux>1){
        return 0;
    }
    return 1;
}

void borrarMundo(animal** mundo){
    /*borra el todas las variable que tiene variable que tiene memoria dinamica*/
    for (int i=0; i<M; i++){

        for (int j=0; j<M; j++){
            free(mundo[i][j].resistencia);
            free(mundo[i][j].fuerza);
            free(mundo[i][j].velocidad);
        }   

    }
}

void mostrarMundo(animal** mundo){
    /*printea el mundo como una matriz*/
    for (int i=0; i<M; i++){
        printf("[");
        for (int j=0; j<M; j++){
            if (mundo[i][j].fuerza == NULL){
                printf("0 ");
            } else {
                printf("1 ");
            }
        }   
        printf("]\n");
    }
}

void reproduccionSimple(animal* padre, animal* madre, animal* hijo){
    /*en la reproduccion simple se le da a lo hijos hereda los atrivutos de 
    el animal correspondiente*/
    int aux = comparar(padre, madre);
    if (aux == 0){
        hijo->tipoFuerza = padre->tipoFuerza;
        hijo->tipoResistecia = padre->tipoResistecia;
        hijo->tipoVelocidad = padre->tipoVelocidad;
        hijo->reproduccion = padre->reproduccion;
        hijo->comerHuir = padre->comerHuir;
        int* e;
        char* c;
        float* f;
        switch (padre->tipoFuerza){
        case 'e':
            e = (int*)malloc(sizeof(int));
            *e = *(int*)(padre->fuerza);
            hijo->fuerza = e;
            break;
        case 'c':
            c = (char*)malloc(sizeof(char));
            *c = *(char*)(padre->fuerza);
            hijo->fuerza = c;
            break;    
        case 'f':
            f = (float*)malloc(sizeof(float));
            *f = *(float*)(padre->fuerza);
            hijo->fuerza = f;
            break;
        default:
            break;
        }
        e = NULL;
        c = NULL;
        f = NULL;
        switch (padre->tipoResistecia){
        case 'e':
            e = (int*)malloc(sizeof(int));
            *e = *(int*)(padre->resistencia);
            hijo->resistencia = e;
            break;
        case 'c':
            c = (char*)malloc(sizeof(char));
            *c = *(char*)(padre->resistencia);
            hijo->resistencia = c;
            break;    
        case 'f':
            f = (float*)malloc(sizeof(float));
            *f = *(float*)(padre->resistencia);
            hijo->resistencia = f;
            break;
        default:
            break;
        }
        e = NULL;
        c = NULL;
        f = NULL;
        switch (padre->tipoVelocidad){
        case 'e':
            e = (int*)malloc(sizeof(int));
            *e = *(int*)(padre->velocidad);
            hijo->velocidad = e;
            break;
        case 'c':
            c = (char*)malloc(sizeof(char));
            *c = *(char*)(padre->velocidad);
            hijo->velocidad = c;
            break;    
        case 'f':
            f = (float*)malloc(sizeof(float));
            *f = *(float*)(padre->velocidad);
            hijo->velocidad = f;
            break;
        default:
            break;
        }
    }else{
        hijo->tipoFuerza = madre->tipoFuerza;
        hijo->tipoResistecia = madre->tipoResistecia;
        hijo->tipoVelocidad = madre->tipoVelocidad;
        hijo->reproduccion = madre->reproduccion;
        hijo->comerHuir = madre->comerHuir;
        int* e;
        char* c;
        float* f;
        switch (madre->tipoFuerza){
        case 'e':
            e = (int*)malloc(sizeof(int));
            *e = *(int*)(madre->fuerza);
            hijo->fuerza = e;
            break;
        case 'c':
            c = (char*)malloc(sizeof(char));
            *c = *(char*)(madre->fuerza);
            hijo->fuerza = c;
            break;    
        case 'f':
            f = (float*)malloc(sizeof(float));
            *f = *(float*)(madre->fuerza);
            hijo->fuerza = f;
            break;
        default:
            break;
        }
        e = NULL;
        c = NULL;
        f = NULL;
        switch (madre->tipoResistecia){
        case 'e':
            e = (int*)malloc(sizeof(int));
            *e = *(int*)(madre->resistencia);
            hijo->resistencia = e;
            break;
        case 'c':
            c = (char*)malloc(sizeof(char));
            *c = *(char*)(madre->resistencia);
            hijo->resistencia = c;
            break;    
        case 'f':
            f = (float*)malloc(sizeof(float));
            *f = *(float*)(madre->resistencia);
            hijo->resistencia = f;
            break;
        default:
            break;
        }
        e = NULL;
        c = NULL;
        f = NULL;
        switch (madre->tipoVelocidad){
        case 'e':
            e = (int*)malloc(sizeof(int));
            *e = *(int*)(madre->velocidad);
            hijo->velocidad = e;
            break;
        case 'c':
            c = (char*)malloc(sizeof(char));
            *c = *(char*)(madre->velocidad);
            hijo->velocidad = c;
            break;    
        case 'f':
            f = (float*)malloc(sizeof(float));
            *f = *(float*)(madre->velocidad);
            hijo->velocidad = f;
            break;
        default:
            break;
        }
    }
} 

void reproduccionCruzada(animal* padre, animal* madre, animal* hijo){
    /*tomas los valores de los punteros padre y madre se ve cual es superior
    le le asigna los valores ya establecidos en el pdf*/
    int aux = comparar(padre, madre);
    if (aux == 0){
        hijo->tipoFuerza = padre->tipoFuerza;
        hijo->tipoResistecia = madre->tipoResistecia;
        hijo->tipoVelocidad = padre->tipoVelocidad;
        hijo->reproduccion = padre->reproduccion;
        hijo->comerHuir = madre->comerHuir;
        int* e;
        char* c;
        float* f;
        switch (padre->tipoFuerza){
        case 'e':
            e = (int*)malloc(sizeof(int));
            *e = *(int*)(padre->fuerza);
            hijo->fuerza = e;
            break;
        case 'c':
            c = (char*)malloc(sizeof(char));
            *c = *(char*)(padre->fuerza);
            hijo->fuerza = c;
            break;    
        case 'f':
            f = (float*)malloc(sizeof(float));
            *f = *(float*)(padre->fuerza);
            hijo->fuerza = f;
            break;
        default:
            break;
        }
        e = NULL;
        c = NULL;
        f = NULL;
        switch (madre->tipoResistecia){
        case 'e':
            e = (int*)malloc(sizeof(int));
            *e = *(int*)(madre->resistencia);
            hijo->resistencia = e;
            break;
        case 'c':
            c = (char*)malloc(sizeof(char));
            *c = *(char*)(madre->resistencia);
            hijo->resistencia = c;
            break;    
        case 'f':
            f = (float*)malloc(sizeof(float));
            *f = *(float*)(madre->resistencia);
            hijo->resistencia = f;
            break;
        default:
            break;
        }
        e = NULL;
        c = NULL;
        f = NULL;
        switch (padre->tipoVelocidad){
        case 'e':
            e = (int*)malloc(sizeof(int));
            *e = *(int*)(padre->velocidad);
            hijo->velocidad = e;
            break;
        case 'c':
            c = (char*)malloc(sizeof(char));
            *c = *(char*)(padre->velocidad);
            hijo->velocidad = c;
            break;    
        case 'f':
            f = (float*)malloc(sizeof(float));
            *f = *(float*)(padre->velocidad);
            hijo->velocidad = f;
            break;
        default:
            break;
        }
    }else{
        hijo->tipoFuerza = madre->tipoFuerza;
        hijo->tipoResistecia = padre->tipoResistecia;
        hijo->tipoVelocidad = madre->tipoVelocidad;
        hijo->reproduccion = madre->reproduccion;
        hijo->comerHuir = padre->comerHuir;
        int* e;
        char* c;
        float* f;
        switch (madre->tipoFuerza){
        case 'e':
            e = (int*)malloc(sizeof(int));
            *e = *(int*)(madre->fuerza);
            hijo->fuerza = e;
            break;
        case 'c':
            c = (char*)malloc(sizeof(char));
            *c = *(char*)(madre->fuerza);
            hijo->fuerza = c;
            break;    
        case 'f':
            f = (float*)malloc(sizeof(float));
            *f = *(float*)(madre->fuerza);
            hijo->fuerza = f;
            break;
        default:
            break;
        }
        e = NULL;
        c = NULL;
        f = NULL;
        switch (madre->tipoResistecia){
        case 'e':
            e = (int*)malloc(sizeof(int));
            *e = *(int*)(madre->resistencia);
            hijo->resistencia = e;
            break;
        case 'c':
            c = (char*)malloc(sizeof(char));
            *c = *(char*)(madre->resistencia);
            hijo->resistencia = c;
            break;    
        case 'f':
            f = (float*)malloc(sizeof(float));
            *f = *(float*)(madre->resistencia);
            hijo->resistencia = f;
            break;
        default:
            break;
        }
        e = NULL;
        c = NULL;
        f = NULL;
        switch (madre->tipoVelocidad){
        case 'e':
            e = (int*)malloc(sizeof(int));
            *e = *(int*)(madre->velocidad);
            hijo->velocidad = e;
            break;
        case 'c':
            c = (char*)malloc(sizeof(char));
            *c = *(char*)(madre->velocidad);
            hijo->velocidad = c;
            break;    
        case 'f':
            f = (float*)malloc(sizeof(float));
            *f = *(float*)(madre->velocidad);
            hijo->velocidad = f;
            break;
        default:
            break;
        }
    }
}

void comerSiempre(animal* a1, animal* a2){
    /*se dedica a comparar las fueszas haciendo las tranformaciones nacasarias*/
    int fuarza1, fuerza2;
    switch (a1->tipoFuerza){
        case 'e':
            fuarza1 = *(int*)a1->fuerza;
            break;
        case 'c':
            fuarza1 = (int)(*(char*)(a1->fuerza)/4);
            break;
        case 'f':
            fuarza1 = roundf((int)*(float*)a1->fuerza);
            break;
        default:
            break;
    }
    switch (a2->tipoResistecia){
        case 'e':
            fuerza2 = *(int*)a1->resistencia;
            break;
        case 'c':
            fuerza2 = (int)(*(char*)(a1->resistencia)/4);
            break;
        case 'f':
            fuerza2 = roundf((int)*(float*)a1->resistencia);
            break;
        default:
            break;
    }
    if (fuarza1 > fuerza2){
        borrar(a2);
    }else{
        borrar(a1);
    }
}

void huirSiempre(animal* a1, animal* a2){
    /*se dadica a hacer las tranformacione correspondientes de 
    los datos velocidad y decide si el animal 1 huye o muera y es*/
    int aux1, aux2;
    switch (a1->tipoVelocidad){
        case 'e':
            aux1 = *(int*)a1->velocidad;
            break;
        case 'c':
            aux1 = (int)(*(char*)(a1->velocidad)/4);
            break;
        case 'f':
            aux1 = roundf((int)*(float*)a1->velocidad);
            break;
        default:
            break;
    }
    switch (a1->tipoVelocidad){
        case 'e':
            aux1 = *(int*)a1->velocidad;
            break;
        case 'c':
            aux1 = (int)(*(char*)(a1->velocidad)/4);
            break;
        case 'f':
            aux1 = roundf((int)*(float*)a1->velocidad);
            break;
        default:
            break;
    }
    switch (a2->tipoVelocidad){
        case 'e':
            aux2 = *(int*)a2->velocidad;
            break;
        case 'c':
            aux2 = (int)(*(char*)(a2->velocidad)/4);
            break;
        case 'f':
            aux2 = roundf((int)*(float*)a2->velocidad);
            break;
        default:
            break;
    }
    switch (a2->tipoVelocidad){
        case 'e':
            aux2 = *(int*)a2->velocidad;
            break;
        case 'c':
            aux2 = (int)(*(char*)(a2->velocidad)/4);
            break;
        case 'f':
            aux2 = roundf((int)*(float*)a2->velocidad);
            break;
        default:
            break;
    }
    if (aux1 <= aux2){
        borrar(a1);
    }
}

void comerAleatorio(animal* a1, animal* a2){
    /*se dedica a defini las variables establecida 
    y loego se comparan*/
    srand(time(NULL));
    int x, aux1, aux2;
    x = rand()%3;
    switch (x){
    case 0:
        switch (a1->tipoFuerza){
        case 'e':
            aux1 = *(int*)a1->fuerza;
            break;
        case 'c':
            aux1 = (int)(*(char*)(a1->fuerza)/4);
            break;
        case 'f':
            aux1 = roundf((int)*(float*)a1->fuerza);
            break;
        default:
            break;
        }
        break;
    case 1:
        switch (a1->tipoResistecia){
        case 'e':
            aux1 = *(int*)a1->resistencia;
            break;
        case 'c':
            aux1 = (int)(*(char*)(a1->resistencia)/4);
            break;
        case 'f':
            aux1 = roundf((int)*(float*)a1->resistencia);
            break;
        default:
            break;
        }
        break;
    case 2:
        switch (a1->tipoVelocidad){
        case 'e':
            aux1 = *(int*)a1->velocidad;
            break;
        case 'c':
            aux1 = (int)(*(char*)(a1->velocidad)/4);
            break;
        case 'f':
            aux1 = roundf((int)*(float*)a1->velocidad);
            break;
        default:
            break;
        }
        break;

    default:
        break;
    }
    srand(time(NULL));
    x = rand()%3;
    switch (x){
    case 0:
        switch (a2->tipoFuerza){
        case 'e':
            aux2 = *(int*)a2->fuerza;
            break;
        case 'c':
            aux2 = (int)(*(char*)(a2->fuerza)/4);
            break;
        case 'f':
            aux2 = roundf((int)*(float*)a2->fuerza);
            break;
        default:
            break;
        }
        break;
    case 1:
        switch (a2->tipoResistecia){
        case 'e':
            aux2 = *(int*)a2->resistencia;
            break;
        case 'c':
            aux2 = (int)(*(char*)(a2->resistencia)/4);
            break;
        case 'f':
            aux2 = roundf((int)*(float*)a2->resistencia);
            break;
        default:
            break;
        }
        break;
    case 2:
        switch (a2->tipoVelocidad){
        case 'e':
            aux2 = *(int*)a2->velocidad;
            break;
        case 'c':
            aux2 = (int)(*(char*)(a2->velocidad)/4);
            break;
        case 'f':
            aux2 = roundf((int)*(float*)a2->velocidad);
            break;
        default:
            break;
        }
        break;

    default:
        break;
    }
    if (aux1 > aux2){
        borrar(a1);
    } else if (a1 < a2){
        borrar(a2);
    } else {
        borrar(a1);
    }
}

// void iteracion(animal** mundo){
//     int x, cont = 0;
//     for (int i=0; i<M; i++){
//         for (int j=0; j<M; j++){
//             if (mundo[i][j].fuerza != NULL){
//                 cont++;
//             }
//         }   
//     }
//     int co = 0;
//     for (int i=0; i<M; i++){
//         for (int j=0; j<M; j++){
//             if (mundo[i][j].fuerza != NULL){
//                 srand(i+j);
//                 moverAnimal(mundo, mundo[i][j], i, j)
//             }
//         }   
//     }
// }
// void moverAnimal(animal** mundo, animal* a1, int j, int j){
//     int x;
//     x = rand()%4;
//     animal* aux;
//     int x, y;
//     switch (x){
//         case 0:
//             y = i++;
//             x = j;
//             if (i>M){
//                 y = i%M;
//             }
//             break;
//         case 1:
//             y = i;
//             x = j++;
//             if (j>M){
//                 x = j%M;
//             }
//             break;
//         case 2:
//             y = i--;
//             x = j;
//             if (i<0){
//                 y = M;
//             }
//             break;
//         case 3:
//             y = i;
//             x = j--;
//             if (j<0){
//                 x = M;
//             }
//             break;
//     default:
//         break;
//     }
// }
#include <stdio.h>
#include <math.h>
#include <stdlib.h>

typedef struct animal{
    void* fuerza;
    char tipoFuerza;
    void* velocidad;
    char tipoVelocidad;
    void* resistencia;
    char tipoResistecia;
    // void (*reproduccion)(struct animal* , struct animal* , struct animal* );
    // void (*comerHuir)(struct animal* ,struct animal* );
} animal;

void mostrarAnimal(animal* a){
    if (a->fuerza != NULL){
        printf("fuerza: ");
        if (a->tipoFuerza == 'e'){
            printf("%d\n", *(int*)a->fuerza);
        } else if (a->tipoFuerza == 'c'){
            printf("%c\n", *(char*)a->fuerza);
        } else if ((a->tipoFuerza == 'f')){
            float* aux = a->fuerza;
            printf("%f\n", *aux);
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

void reproduccionSimple(animal* padre, animal* madre, animal* hijo){
    int aux = comparar(padre, madre);
    if (aux = 1){
        hijo->tipoFuerza = padre->tipoFuerza;
        hijo->tipoResistecia = padre->tipoResistecia;
        hijo->tipoVelocidad = padre->tipoVelocidad;
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

void borrar(animal* a){
    free(a->fuerza);
    a->fuerza = NULL;
    free(a->velocidad);
    a->velocidad = NULL;
    free(a->resistencia);
    a->resistencia = NULL;
}


int main(){
    int* x = (int*)malloc(sizeof(int));
    float* z = (float*)malloc(sizeof(float));
    char* y = (char*)malloc(sizeof(char));
    *x = 12;
    *z = 3.14;
    *y = 'e';
    animal legoshi = {x, 'e', z, 'f', y, 'c'};
    animal haru = {z, 'f', y, 'c', x, 'e'};
    animal* haroshi = malloc(sizeof(animal));
    reproduccionSimple(&haru, &legoshi, &haroshi);
    mostrarAnimal(&haroshi);
    borrar(&haru);
    mostrarAnimal(&haroshi);
    return 1;
}
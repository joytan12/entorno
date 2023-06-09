#include "procesos.h"
#include "mazo.h"

#define juega 10000
#define fin 20000

// 4 a 5 digitos que representan lo siguiente
// por ejemplo 1234 
// esta intruccion dice que la carta 12 fue jugada en la fila 3 por el jugador 4 

void cicloJugador(int *pipe, Tablero *tablero, Mano *mano){
    printf("entre al ciclo jugador");
    int orden, jugadas[4];
    for (int i = 0; i < 10; i++)
    {
        for (int j = 0; j < 1; j++)
        {
            read(pipe[0], &orden, BUFFER_SIZE);
            jugadas[j] = orden;
            actualizar(jugadas[j], tablero, mano);
        }
        if ((jugadas[1] / juega) == 1){
            int arr[2];
            eligeJugador(mano, tablero, 0, &arr);
            int jugada = arr[0] * 100 + arr[1] * 10 + 1;
            write(pipe[1], &jugada, BUFFER_SIZE);//se manda el valor jugado en un conjugo de 4 numeros 1243
        } else if ((jugadas[1] / fin) == 2){
            int menor = 99999;
            int pos = -1;
            for (int i = 0; i < NUM_HIJOS; i++)
            {
                if (menor > mano[i].puntaje){
                    menor = mano[i].puntaje;
                    pos = i;
                }
            }
            if(pos == 0){
                printf("ganaste!!!!");
            }
            printf("el ganador es el jugador %d, con %d", pos, menor);
            break;
        }
    }
}

void cicloBot(int *pipe, Tablero *tablero, Mano *mano, int jugador){
    printf("SOY EL JUGADOR %d\n", jugador);
    long int orden;
    long int pre = 0;
    while (1)
    {   
        read(pipe[0], &orden, BUFFER_SIZE);
        if ((orden / juega) == 1 && pre != orden){
            pre = orden;
            int act[4]; //son acciones que se hicienron en el jeugo anterio por los demas procesos
            for (int i; i < 4; i++){
                act[i] = (orden / (10000 * i) % 10000);
                actualizar(*act, tablero, mano);
            }
            int arr[2], pila;
            srand(time(0)); // inicializar la semilla para la función rand()
            pila = rand() % 4 + 1;
            eligeBot(mano, pila, tablero[pila].top, &arr);
            arr[1] = pila;
            int jugada = arr[0] * 100 + arr[1] * 10 + jugador; //mismo formato de cuatro numeros
            write(pipe[1], &jugada, BUFFER_SIZE);
        } else if ((orden / fin) == 2 && pre != orden){// fin del juego
            break;
        }
    }
}

void cicloPadre(Tablero *tablero, Mano *manos, int *pipe1, int *pipe2, int *pipe3, int *pipe4){
    printf("entro al padre");
    int jugada[4]; 
    for (int i = 0; i < 10; i++){
        read(pipe1[0], &jugada, BUFFER_SIZE);
        actualizar(jugada, tablero, manos);
        jugada[1] = 10000;
        if (i == 9){
            jugada[1] = 20000;
        }
        write(pipe1[1], &jugada, BUFFER_SIZE);
    }
}

void cicloPadre1(Tablero *tablero, Mano *manos, int *pipe1){
    printf("entro al padre");
    int jugada = 10000; 
    for (int i = 0; i < 10; i++){
        write(pipe1[1], &jugada, BUFFER_SIZE);
        read(pipe1[0], &jugada, BUFFER_SIZE);
        actualizar(jugada, tablero, manos);
        if (i == 9){
            jugada = 20000;
        } else {
            //pasar las jugadas de los demas
            jugada = 10000;
        }
        
    }
}

void bubbleSort(int *arr, int n) {
    int i, j, temp;
    for (i = 0; i < n - 1; i++) {
        for (j = 0; j < n - i - 1; j++) {
            if (arr[j] > arr[j+1]) {
                // intercambiar elementos
                temp = arr[j];
                arr[j] = arr[j+1];
                arr[j+1] = temp;
            }
        }
    }
}

long int ordenar(int *jugada, int estado, Tablero *tablero){
    bubbleSort(jugada, NUM_HIJOS);
    long int orden = 0;
    long int base = 10000;
    for (int i = 0; i < NUM_HIJOS; i++)
    {   
        long int potencia = 1;
        printf("%d", i);
        for (int j = 0; j < i; j++)
        {
            potencia = potencia * base;
        }
        orden += jugada[i] * potencia;
    }
    return orden;
}

// actualiza el tablero con la jugadas anteriares 
void actualizar(int intruccion, Tablero *tablero, Mano *mano){
    int carta, fila, jugador, puntos;
    jugador = intruccion % 10;
    fila = (intruccion / 10) % 10;
    carta = (intruccion / 100) % 100;
    puntos = jugarCarta(&carta, tablero, fila);
    mano[jugador].puntaje += puntos;
    sacarCarta(mano, &jugador, &carta);
}

// esta funcion saca la carta de la mano del jugador, esta ya esta jugada y no se 
// considera en el juego
void sacarCarta(Mano *mano, int jugador, int carta){
    for (int i = 0; i < 10; i++)
    {
        if (mano[jugador].mano[i] == carta){
            mano[jugador].mano[i] == 46;
            break;
        }
    }
    
}

void verMano(Mano *manos, int indice){
    for (int i = 0; i < 10; i++)
    {   
        if (manos[indice].mano[i] < 45){
            printf("(%d | %d) ", manos[indice].mano[i], i);
        }
    }
    printf("\n");
}

void verFilas(Tablero *tablero){
    for (int i = 0; i < NUM_HIJOS; i++)
    {  
        for (int j = 0; j < tablero[i].cartas; j++)
        {
            printf("|%d| ", tablero[i].pila[j]);
        }
        printf("\n");
    }
    printf("\n");
}

int eligeJugador(Mano *manos, Tablero *tablero, int indice, int *arr){
    int opcion = 1;
    int aux, pos, carta;
    while (opcion == 1)
    {   
        printf("opciones \nver mano(1)\nver topes de las filas(2)\nelegir carta(3)\n");
        scanf("%d", &aux);
        if (aux == 1){
            verMano(manos, indice);
        } else if(aux == 2){
            verFilas(tablero);
        } else if(aux == 3){
            printf("que posicion escoges: ");
            scanf("%d", &pos);
            opcion = 0;
        }
    }

    carta = manos[indice].mano[pos];
    manos[indice].mano[pos] = 46;

    int act, post; 

    for (int i = 0; i < 10; i++)
    {
        for (int j = i; j < 9; j++)
        {
            act = manos[indice].mano[j];
            post = manos[indice].mano[j+1];
            if (act == 46){
                manos[indice].mano[j] = post;
                manos[indice].mano[j+1] = act;
            }
        }   
    }

    // este for comprueba si la carta elegida es mayor a los topes de todas las 
    // filas, si es asi se da a elegir en que fila se quiere dejar la carta
    int tope = 0;
    for (int i = 0; i < 4; i++)
    {
        if (tablero[i].top < carta){
            tope++;
        } 
    }
    int f = -1;
    if (tope == 3){
        verFilas(tablero);
        printf("elige la fila en la que quieres poner tu carta");
        scanf("%d", f);
    }
    arr[1] = f;
    arr[0] = carta;
}

int eligeBot(Mano *manos, int indice, int menor, int *arr){
    int pos = -1;
    for (int i = 0; i < 10; i++)
    {
        if (manos[indice].mano[i] < menor){
            menor = manos[indice].mano[i];
            pos = i;
        }
    }
    if(pos == -1){
        menor = 45;
        for (int i = 0; i < 10; i++)
        {
            if (manos[indice].mano[i] < menor){
                menor = manos[indice].mano[i];
                pos = i;
            }
        }
        
    }
    manos[indice].mano[pos] = 46;
    arr[0] = menor;
    arr[1] = pos;
}
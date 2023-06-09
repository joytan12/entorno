#include "mazo.h"
#include "procesos.h"

int main () {
    int cartas[NUM_CARTA];
    Tablero tablerito[NUM_HIJOS];
    Mano manuel[NUM_HIJOS];


    generarMazo(&cartas);
    repartirMano(&tablerito, &manuel, cartas);
    int pipe1[2], pipe2[2], pipe3[2], pipe4[2];
    pipe(pipe1);
    // pipe(pipe2);
    // pipe(pipe3);
    // pipe(pipe4);
    // for (int i = 0; i < NUM_HIJOS; i++)
    // {
    //     pid = fork();
    //     if (pid == 0){
    //         // codigo de los hijos
    //         if (i == 0){
    //             cicloJugador(&pipe1, &tablerito, &manuel);
    //             close(pipe1[0]);
    //         } else if (i == 1){
    //             cicloBot(&pipe2, &tablerito, &manuel, i);
    //             close(pipe2[0]);
    //         } else if (i == 2){
    //             cicloBot(&pipe3, &tablerito, &manuel, i);
    //             close(pipe3[0]);
    //         } else {
    //             cicloBot(&pipe4, &tablerito, &manuel, i);
    //             close(pipe4[0]);
    //         }
    //         exit(EXIT_SUCCESS);
    //     } else {
    //         cicloPadre(&tablerito, &manuel, &pipe1, &pipe2, &pipe3, &pipe4);
    //         close(pipe1[0]);
    //         close(pipe1[1]);
    //         close(pipe2[0]);
    //         close(pipe2[1]);
    //         close(pipe3[0]);
    //         close(pipe3[1]);
    //         close(pipe4[0]);
    //         close(pipe4[1]); 
    //     }
    // }


    // --------------------------------------------------------------------------------
    // printf("hola mano");
    // if (fork() == 0){
    //     printf("hola mano");
    //     cicloJugador(&pipe1, &tablerito, &manuel);
    // } else {
    //     printf("hola mano");
    //     cicloPadre1(&tablerito, &manuel, &pipe1);    
    // }
    printf("cartas\n");
    for (int i = 0; i < NUM_CARTA; i++)
    {
        printf("|%d| ", cartas[i]);
    }
    printf("\n");
    
    for (int i = 0; i < 10; i++)
    {   
        printf("el turno es %d\n", (i+1));
        int arr[2], opcion, carta;
        carta = eligeJugador(&manuel, &tablerito, 0, &arr);
        int pos = arr[1];
        if(arr[1] != -1){
            jugarCarta(arr[0], &tablerito, arr[1]);
        } else {
            int diferencia = 45, pos = arr[1];
            for (int j = 0; j < 4; j++)
            {
                if((tablerito[j].top - carta) < diferencia){
                    diferencia = (tablerito[j].top - carta);
                    pos = j;
                }
            }
            jugarCarta(arr[0], &tablerito, pos);
        }
        printf("las cartas de la fila son %d ", tablerito[pos].cartas);
    }
    

    return 0; 
}
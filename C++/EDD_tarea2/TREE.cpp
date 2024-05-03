#include <fstream>
#include <iostream>
#include <array>
#include <math.h>
#include <string>
#include <sstream>
#include <cstdlib>



struct tElementoarbol{ //elemento de un arbol
    int x;
    int y;
};
struct tEval{ //estructura para sacar datos de EVALUAR
    int x;
    float y;
};

struct tNodoarbol{ //estructura que almacena el nodo de un arbol
    tElementoarbol info;
    tNodoarbol* left;
    tNodoarbol* right;
};

struct treeBinar{ //semilla del arbol
    tNodoarbol* root;
    int tElemns;
};

class aBinario{ //IMPLEMENTACION DEL ARBOL
private:
    treeBinar* start; //puntero al 1er nodo
    int i; //definicion auxiliar
public:
    aBinario(){
        start = new treeBinar{NULL,0}; 
    };
    ~aBinario(){
        tNodoarbol* array[start->tElemns];
        i = 0;
        inOrden(start->root, array, i);
        int elem = start->tElemns;
        delete[] start;
        for (int i=0; i<elem; i++){
            delete[] array[i];
        }
    };
    //operaciones get
    int size(){
        /*retorna la cantidad de elementos*/
        return start->tElemns;//retorna la cantidad de elementos del arbol
    }
    tNodoarbol* getFirs(){
        /*retorna el primer elemento del arbol*/
        return start->root;//retorna un puntero al primer nodo
    }
    //operaciones con arbol
    void append(tElementoarbol item, tNodoarbol* aux){
        /*esta funcion ubica el monomio en la pocicion que 
        deberia ir segun su exponente */
        if (aux == NULL){
            start->root = new tNodoarbol{item,NULL,NULL};
            start->tElemns++;
        } else {
            if (item.y > aux->info.y){
                if (aux->right==NULL){
                    aux->right = new tNodoarbol{item,NULL,NULL};
                    start->tElemns++;
                } else {
                    append(item, aux->right);
                }
            } else {
                if (aux->left==NULL){
                    aux->left = new tNodoarbol{item,NULL,NULL};
                    start->tElemns++;
                } else {
                    append(item, aux->left);
                }
            }
        }
    }
    void inOrden(tNodoarbol* aux, tNodoarbol* array[], int &i){
        /*recorrido del arbol en in orden
        no retorna nada, ya que modifica variables 
        las cuales se le entregan a la funcion*/
        if (aux != NULL){
            inOrden(aux->left, array, i);
            array[i] = aux;
            // cout << array[i] <<' '<< aux->info.x << ' ' << aux->info.y << " en la pocison " << i << endl;
            i++;
            inOrden(aux->right, array, i);
        }
    };
    void print(){
        /*esta funcion nos ayuda a printear la lista, 
        y ya que no altera nada en la lista,
        no se retorna nada*/
        tNodoarbol* array[start->tElemns];
        i = 0;
        inOrden(start->root, array, i);
        for (int i=0; i<start->tElemns; i++){
            cout << array[i]->info.x << ' ' << array[i]->info.y << ' ' ;
        }
        cout << endl;
    }
};

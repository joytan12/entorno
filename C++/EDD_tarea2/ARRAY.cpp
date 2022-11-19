#include <fstream>
#include <iostream>
#include <array>
#include <math.h>
#include <string>
#include <sstream>
#include <cstdlib>

using namespace std;

struct tElemento{ //estructura que almacenara los monomios
        unsigned int x;
        int y;
};
struct tNodo{ //estructura nodo
    tElemento info;
    tNodo* sig;
};
struct sCoeficiente{ //estructura que se usara para estudiar cuando COEFICIENTE aparezca en el txt
    int coefC;
    int nPolCoef;
};
struct sEvaluar{ //estructura que se usara para estudiar cuando EVALUAR aparezca en el txt
    int nPolEval;
    float xToEval;
};

class tLista{
private:
    tNodo* head; //cabezera de la lista
    tNodo* tail; //final de la lista
    tNodo* curr; //cursor de la lista
    tNodo* aux; //variable para no definir nuevos punteros en cada funcion
    unsigned int listSize; //cantidad de elementos de la lista
    unsigned int pos; //posicion del cursor
public:
    tLista(tElemento A[] = {}, unsigned int size = 0, unsigned int n = 0){
        tElemento vacio; //variabla vacia para la cabezara
        pos = n;
        if (size > 0){
            tail = new tNodo{A[static_cast<int>(size)-1], NULL};
            curr = tail;
            for(int i=0; i<static_cast<int>(size)-1; i++){
                aux = new tNodo{A[i], curr};
                curr = aux;
            }
            head = new tNodo{vacio, curr};
        } else if (size == 1){
            tail = new tNodo{A[1], NULL};
            curr = tail;
            head = new tNodo{vacio, curr};
        }else {
            curr = head = tail = new tNodo{vacio, NULL}; 
        }
        listSize = size;
        if (pos != 0){//estoy arreglando el pos y curr
            aux = head; 
            for(int i=0; i<static_cast<int>(pos); i++){
                curr = aux->sig;
                aux = curr;
            }
        }
    }

    ~tLista(){
        aux = head->sig;
        delete head;
        do{
            curr = aux->sig;
            delete aux;
            aux = curr;
        } while(aux->sig != NULL);
        delete aux;
    };

    //operaciones con lista
    int append(tElemento item){
        /*
        se agrega un elemento al final de la lista y 
        retorna el nuevo largo de la lista
        */
        tNodo* aux = new tNodo{item, NULL};
        tail->sig = aux;
        tail = aux;
        listSize++;
        return listSize; //largo de la lista
    }
    int moveToStart(){
        // mueve el curr al inicio de la lista
        curr = head->sig;
        pos = 1;
        return pos;
    }
    int moveToEnd(){
        // mueve el curr al final de la lista
        curr = tail;
        pos = listSize;
        return pos;
    }
    int insert(tElemento item){ 
        /*inserta un item tipo tElemento en la posicion 
        en donde se ubica en curr*/
        tNodo* aux = new tNodo{item, curr->sig}; //se crea un nuevo struct
        curr->sig = aux; 
        listSize++;
        return pos;//retorna la posicion del cursor
    }

    void delet(){
        /*esta funcion se dedica a eliminar el elemento
        al que esta apuntando el cursor*/
        aux = head;
        int i; //condicion para el for
        for (i = 1; i < static_cast<int>(pos); i++){
            aux = aux->sig;
        }
        aux->sig = curr->sig;
        delete curr;
        if (curr == tail){
            aux->sig = NULL;
            tail = curr = aux;
            pos--;
        } else {
            curr = aux;
        }
        listSize--;
    }
    unsigned int length(){
        return listSize;
    }

    tElemento get(int n){
        /*la funcion se dedica a obtener el dato
        que apunta el cursor, devuelve el tElemento*/
        if (n > static_cast<int>(listSize)){
            exit(1);
        }
        aux = head;
        for (int i = 0; i < n; i++){
            aux = aux->sig;
        }
        return aux->info; //retorna el struct tElemento
    }
    // movimientos del cursor
    unsigned int getPos(){
        /*devuelve el pos*/
        return pos; //retorna la posicion
    }

    tElemento getValue(){
        /*esta funcion retorna el tElemento
        al que apunta el curr*/
        return curr->info; //retorna el item
    }

    int next(){
        /*el curr pasa a la siguiente posicion*/
        if (pos == listSize){
            exit(1);
        }
        curr = curr->sig;
        pos++;
        return pos; //retorna la posicion
    }

    int previous(){
        /*ubiaca el curr en el item anterior de la lista, 
        tambien pone el pos en su respectivo lugar*/
        if(getPos() == 0){
            exit(1);
        }
        curr = head; 
        for(unsigned int i=1; i<pos; i++){
            curr = curr->sig;
        }
        pos--;
        return pos;//retorna la posicion actualizada
    }
    // recorrer lista
    void print(){
        //esta funcion nos ayudara a printear la lista, 
        //ya que no altera nada en la lista 
        //no se retorna nada        
        aux = head;
        do{
            aux = aux->sig;
            cout << aux->info.x << ' ' << aux->info.y << ' ' ;
        } while (aux->sig != NULL);
        cout << endl;
    }
};




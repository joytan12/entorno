#include <fstream>
#include <iostream>
#include <array>
#include <math.h>

#define VACIA -1 // variable para definir una casilla vacia

using namespace std;

typedef int tipoClave;

struct oferta {
    int cod_producto;
    int cantidad_descuento;
    int descuento;
    int productos_equivalentes[10];
};

struct ranuraOferta {
    tipoClave k;
    oferta I; 
};

bool esPrime(int z){
    if (z == 0 || z == 1 || z == 4) return false;
    for (int i=2; i<z/2; i++){
        if ((z % i) == 0) return false;
    }
    return true;
}

class hashingOferta{
private:
    ranuraOferta *HT;
    int *colisiones;
    int M;
    int prime;
public:
    hashingOferta(int M);
    ~hashingOferta();
    int h(int k);
    int h2(int k);
    int p(tipoClave k, int i);
    int hashInsert(tipoClave k, oferta I);
    oferta hashSearch(tipoClave k);
};

hashingOferta::hashingOferta(int n){
    M = n;
    HT = new ranuraOferta[M];
    int i;
    for (i=0; i<M; i++){
        if (esPrimo(i)){
            prime = i;
        }
    }
    colisiones = new int[M];
    for (int i=0; i < M; i++){
        HT[i].k = VACIA;
        colisiones[i] = 0;
    }
}

hashingOferta::~hashingOferta(){
    delete[] HT;
    delete[] colisiones;
}

int hashingOferta::h(int k){
    return k % M;
}

int hashingOferta::h2(int k){
    return (prime  - (k % prime ));
}

int hashingOferta::p(tipoClave k, int i){
    if (i == 0) return 0;
    return i*h2(k);
}

int hashingOferta::hashInsert(tipoClave k, oferta I){
    int inicio, i;
    int pos = inicio = h(k);
    int cols = 0;
    for (i = 1; HT[pos].k != VACIA && HT[pos].k != k; i++){
        pos = (inicio + p(k, i)) % M;
        cols++;
    }
    // cout << HT[pos].k << " " << k << endl;
    if (HT[pos].k == k){
        return 0;
    } else {
        HT[pos].k = k;
        HT[pos].I = I;
        colisiones[pos] = cols;
        return 1;
    }
}

oferta hashingOferta::hashSearch(tipoClave k){
    int inicio, i;
    int pos = inicio = h(k);
    for (i=1; HT[pos].k != VACIA && HT[pos].k != k; i++){
        pos = (inicio + p(k, i)) % M;
    }
    if (HT[pos].k == k){
        return HT[pos].I;
    } else {
        int i[10]={-1,-1,-1,-1,-1,-1,-1,-1,-1,-1};
        oferta invalido = {0, 0, 0, i[10]};
        return invalido;//arreglar final 
    }
}
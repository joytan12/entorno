#include <fstream>
#include <iostream>
#include <array>
#include <math.h>

#define VACIA -1 // variable para definir una casilla vacia

using namespace std;

typedef int tipoClave;

struct producto {
    int  cod_producto;
    char nombre_producto[31];
    int  precio;
};

struct ranura {
    tipoClave k;
    producto I; 
};

bool esPrimo(int z){
    if (z == 0 || z == 1 || z == 4) return false;
    for (int i=2; i<z/2; i++){
        if ((z % i) == 0) return false;
    }
    return true;
}

class hashing{
private:
    ranura *HT;
    int *colisiones;
    int M;
    int prime;
public:
    hashing(int n);
    ~hashing();
    int h(int k);
    int h2(int k);
    int p(tipoClave k, int i);
    int hashInsert(tipoClave k, producto I);
    producto hashSearch(tipoClave k);
    void recorrer();
};

hashing::hashing(int n){
    M = n;
    HT = new ranura[M];
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

hashing::~hashing(){
    delete[] HT;
    delete[] colisiones;
}

int hashing::h(int k){
    return k % M;
}

int hashing::h2(int k){
    return (prime - (k % prime));
}

int hashing::p(tipoClave k, int i){
    if (i == 0) return 0;
    return i*h2(k);
}

int hashing::hashInsert(tipoClave k, producto I){
    int inicio, i;
    int pos = inicio = h(k);
    int cols = 0;
    for (i = 1; HT[pos].k != VACIA && HT[pos].k != k; i++){
        pos = (inicio + p(k, i)) % M;
        cols++;
    }
    if (HT[pos].k == k){
        return 0;
    } else {
        HT[pos].k = k;
        HT[pos].I = I;
        colisiones[pos] = cols;
        return 1;
    }
}

producto hashing::hashSearch(tipoClave k){
    int inicio, i;
    int pos = inicio = h(k);
    for (i=1; HT[pos].k != VACIA && HT[pos].k != k; i++){
        pos = (inicio + p(k, i)) % M;
    }
    if (HT[pos].k == k){
        return HT[pos].I;
    } else {
        producto invalido = {0, " ", 0};
        return invalido;//arreglar final 
    }
}
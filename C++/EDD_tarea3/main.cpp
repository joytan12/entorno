#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include <cstdlib>
#include "hashing.cpp"
#include "hashingOferta.cpp"

using namespace std;

int main(){
    int j;
    fstream fp;
    string linea, cantproductos, cantClientes, line;
    oferta tem;
    producto aux;
    fp.open("productos.dat", ios::in | ios::binary);
    fp.read((char*)&j, sizeof(int));
    j = (j / 0.7) + 1;
    hashing productos(j);
    for (int i=0; i<j; i++){
        fp.read((char *)&aux, sizeof(producto));
        productos.hashInsert(aux.cod_producto, aux);
    }
    fp.close();
    fp.open("ofertas.dat", ios::in | ios::binary);
    fp.read((char*)&j, sizeof(int));
    j = (j / 0.7) + 1;
    hashingOferta ofertas(j);
    for (int i=0; i<j; i++){
        fp.read((char *)&tem, sizeof(oferta));
        ofertas.hashInsert(tem.cod_producto, tem);
    }
    fp.close();
    fp.open("compras.txt", ios::in);
    getline(fp, cantClientes);
    // cout << "Numero de clientes = " << cantClientes << endl;
    int boleta[stoi(cantClientes)];
    for(int c=0; c<stoi(cantClientes); c++){
        getline(fp, cantproductos);
        int compras[stoi(cantproductos)];
        // cout << "Cantidad de productos del cliente " << c << " = " << cantproductos << endl;
        for(int p=0; p<stoi(cantproductos); p++){
            getline(fp, line);
            compras[p] = stoi(line);
            // cout << "Producto comprado código = " << stoi(line) << endl;
            // cout << productos.hashSearch(stoi(line)).nombre_producto << endl;
        }
        int comprasOrdenadas[stoi(cantproductos)], menor, marca;
        for (int i=0; i<stoi(cantproductos); i++){ // ordena las compras
            menor = compras[i];
            marca = i;
            int j;
            for (j=0; j<stoi(cantproductos); j++){
                if (menor < compras[j]){
                    menor = compras[j];
                    marca = j;
                }
            }
            compras[marca] = -1; 
            comprasOrdenadas[i] = menor;
        }
        for (int i=0; i<stoi(cantproductos); i++){
            cout << comprasOrdenadas[i] << " ";
        }
        cout << endl;
        int total = 0, cantidad = 1, contOferta = cantidad;
        if (stoi(cantproductos) == 1){
            total = productos.hashSearch(comprasOrdenadas[0]).precio;
            boleta[0] = total;
        }
        for (int i=1; i<stoi(cantproductos); i++){
            if (comprasOrdenadas[i] != comprasOrdenadas[i-1] && comprasOrdenadas[i-1] != -1 && comprasOrdenadas[i] != -1){
                int cons = comprasOrdenadas[i];
                total += cantidad*productos.hashSearch(comprasOrdenadas[i-1]).precio;
                cantidad = 0;
                for (int j=i; j<stoi(cantproductos); j++){
                    cout << "\n" << j << endl;
                    cout << comprasOrdenadas[i-1] << " la compra anterior se compara con "<< comprasOrdenadas[j] << " la cantidad es " << cantidad << endl;
                    if (comprasOrdenadas[j] != cons){
                        total += cantidad*productos.hashSearch(cons).precio;
                        cantidad = 0;
                        cons = comprasOrdenadas[j];
                    }
                    for (int x=0; x<10; x++){
                        int tmp = ofertas.hashSearch(comprasOrdenadas[i-1]).productos_equivalentes[x];
                        if (tmp != -1 && tmp == comprasOrdenadas[j]){
                            cantidad++;
                            contOferta++;
                            comprasOrdenadas[j] = -1;
                        }
                        cout << ofertas.hashSearch(comprasOrdenadas[i-1]).productos_equivalentes[x] << " ";  
                    }
                    cout << endl;
                }
                if (ofertas.hashSearch(comprasOrdenadas[i-1]).cantidad_descuento != 0){
                    total -= (contOferta / ofertas.hashSearch(comprasOrdenadas[i-1]).cantidad_descuento) * ofertas.hashSearch(comprasOrdenadas[i-1]).descuento;
                }
                contOferta = 0;
                for (int p=0; p<i; p++){
                    comprasOrdenadas[p] = -1;
                }
                // for (int i=0; i<stoi(cantproductos); i++){cout << comprasOrdenadas[i] << " ";}cout << endl;
                cantidad = 0;
            } else if (comprasOrdenadas[i] != -1) {
                cantidad++;
                contOferta++;
                if (i == stoi(cantproductos)-1){
                    total += cantidad*productos.hashSearch(comprasOrdenadas[i]).precio;
                    if (ofertas.hashSearch(comprasOrdenadas[i]).cantidad_descuento != 0){
                        total -= (contOferta / ofertas.hashSearch(comprasOrdenadas[i]).cantidad_descuento) * ofertas.hashSearch(comprasOrdenadas[i]).descuento;
                    }
                }
            }
        }
        boleta[c] = total;
    }
    fp.close();
    fp.open("boletas.txt");
    for (int i=0; i<stoi(cantClientes); i++){
        fp << i+1 << endl;
        fp << boleta[i] << endl;
    }
    fp.close();
    return 1;
}
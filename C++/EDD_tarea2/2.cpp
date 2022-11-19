#include <fstream>
#include <iostream>
#include <array>
#include <math.h>
#include <string>
#include <sstream>
#include <cstdlib>
#include "TREE.cpp"

using namespace std;

int TreePolinomios(){ //CALCULA COEFICIENTE Y EVALUAR DADOS EN EL TXT PARA LOS POLINOMIOS DADOS EN EL TXT TAMBIEN, ALMACENA ESTOS POLINOMIOS EN ARBOLES
    ifstream texto;
    string line, lineastring, elemento;
    int contador = 0, cantidadPolinomios = 0, nMonomios = 0, nPolinomio = 0; //cantidad es la cantidad de polinomios.
    texto.open("entradaPolinomio.txt", ifstream::in);
    getline(texto,line,'\n');
    cantidadPolinomios = stoi(line);
    aBinario arboles[cantidadPolinomios];
    while (getline(texto,line,'\n')){
        stringstream lineastring(line);
        if (line[0] != 'C' && line[0] != 'E') {
            nMonomios = stoi(line);
            for (int k = 1; k <= nMonomios; k++){ //Obtener los valores de cada monomio
                getline(texto,line,'\n');
                stringstream lineastring(line);
                tElementoarbol valores;
                for (int i = 0; getline(lineastring,elemento,' '); i++){
                    if (i == 0) {
                        valores.y = stoi(elemento);
                    } else if (i == 1) {
                        valores.x = stoi(elemento);
                    }
                }
                arboles[nPolinomio].append(valores, arboles[nPolinomio].getFirs());
                contador++;
            }
            nPolinomio++;
        }
    }
    texto.close();

    string str;
    int totalPolinomios = 0;
    ifstream archivo;
    archivo.open("entradaPolinomio.txt", ifstream::in);
    getline(archivo, str,'\n');
    totalPolinomios = stoi(str);
    archivo.close();
    for (int i = 0; i<totalPolinomios; i++){

        int x = 0, maxExp = 0; //Variable para obtener exponente maximo de los monomios.
        tNodoarbol* recorrer[arboles[i].size()];
        arboles[i].inOrden(arboles[i].getFirs(), recorrer, x);
        for (int p = arboles[i].size()-1; p>=0; p--){ //Recorrer el arbol de manera invertida, para encontrar el maximo exponente.
            if (p == arboles[i].size()-1){
                maxExp = recorrer[p]->info.y; //Obtenemos el maximo exponente, para asi agregar los faltantes.
            }
        }
        for(int c = 0; c<=maxExp; c++){
            int x = 0;
            tNodoarbol* recorrer[arboles[i].size()];
            arboles[i].inOrden(arboles[i].getFirs(), recorrer, x);
            if(c<recorrer[c]->info.y){
                tElementoarbol valores{0,c};
                arboles[i].append(valores, arboles[i].getFirs());
            }
        }
        arboles[i].print();
    }

    ifstream txt;
    ofstream wrt;
    txt.open("entradaPolinomio.txt", ifstream::in); //abrimos el txt de nuevo para estudiar los COEFICIENTE y EVALUAR
    wrt.open("salidaPolinomio.txt"); //abrimos archivo de escritura
    while (getline(txt,str,'\n')){
        if (str[0] == 'C'){

            tElementoarbol coef; //COEFICIENTE
            coef.x = stoi(str.substr(str.find(' '),str.rfind(' ')));           
            coef.y = stof(str.substr(str.rfind(' '),str.size()));
            int resCoef=0; //resultado del calculo para COEFICIENTE
            for (int i=0; i<arboles[coef.x].size(); i++){
                int h = 0;
                tNodoarbol* recorrer[arboles[coef.x].size()];
                arboles[coef.x].inOrden(arboles[coef.x].getFirs(), recorrer, h);
                
                for (int z=0; z<arboles[coef.x].size() ; z++){
                    if (recorrer[z]->info.y == coef.y){
                        resCoef = recorrer[z]->info.x;
                    };
                }
            }
            wrt << resCoef << endl;

        } else if (str[0] == 'E'){
            tEval eval; //EVALUAR 
            eval.x = stoi(str.substr(str.find(' '),str.rfind(' ')));           
            eval.y = stof(str.substr(str.rfind(' '),str.size()));
            float resEval=0; //resultado del calculo para EVALUAR

            cout << arboles[eval.x].size() << endl;
            tNodoarbol* recorrer[arboles[eval.x].size()]; //nodo arbol auxiliar para recorrer el polinomio 
            for (int i=arboles[eval.x].size()-1; i>=0; i--){
                int h = 0; //entero auxiliar
                arboles[eval.x].inOrden(arboles[eval.x].getFirs(), recorrer, h);
                resEval*=eval.y;
                resEval+=recorrer[i]->info.x;

                
            }
            wrt << resEval << endl;
        } 
    }
    return 0;
}
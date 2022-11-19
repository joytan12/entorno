#include <fstream>
#include <iostream>
#include <array>
#include <math.h>
#include <string>
#include <sstream>
#include <cstdlib>
#include "ARRAY.cpp"


int TDApolinomios(){ //CALCULA COEFICIENTE Y EVALUAR DADOS EN EL TXT PARA LOS POLINOMIOS DADOS EN EL TXT TAMBIEN, ALMACENA ESTOS POLINOMIOS EN ARREGLOS

    /* ---------- DEFINICIONES y lectura txt ---------- */
    ifstream txt; 
    string str; //line: linea del .txt | str: linea del .txt convertida para operar
    int totalPolinomios; //totalPolinomios: numero tot de polinomios
    unsigned int nMonomios; //nMonomios: numero de monomios dentro de un polinomio
    int nPolinomio = 0; //numero de polinomio estudiado
    sCoeficiente coeficiente; //coelficiente: guarda los vaores dados en la linea del txt COEFICIENTE
    sEvaluar evaluar; //evaluar: guarda los valores dados en la linea del txt EVALUAR
    tElemento monomio; // monomio: almacena 1 monomio
    txt.open("entradaPolinomio.txt", ifstream::in);
    getline(txt, str,'\n');
    totalPolinomios = stoi(str);
    tLista polinomios[totalPolinomios]; //polinomios: arreglo con todos los polinomios dentro del txt
    while(!txt.eof()){
        getline(txt, str,'\n');
        
        
        if (str[0]!='C' && str[0]!= 'E') {
            nMonomios = stoi(str);
            for (unsigned int i=0 ; i < nMonomios ; i++){

                getline(txt, str,'\n');
                int spc = str.rfind(' ');
                monomio.x = stoi(str.substr(0,spc));
                monomio.y = stoi(str.substr(spc,str.length()));
                polinomios[nPolinomio].append(monomio);
                
            };//ENDFOR
            
            nPolinomio+=1;
        }
        
    } //ENDWHILE
    for (int i=0; i<totalPolinomios; i++){
        bool condicion = true; //condicion para el while
        polinomios[i].moveToStart();
        while (condicion){
            tElemento pre, next; // esto elementos son para evaluar en anterior y despues
            pre = polinomios[i].getValue();
            polinomios[i].next();
            next = polinomios[i].getValue();
            if (pre.x > next.x){    
                polinomios[i].previous();
                polinomios[i].previous();
                polinomios[i].insert(next);
                polinomios[i].next();
                polinomios[i].next();
                polinomios[i].next();
                polinomios[i].delet();                
            }
            if (polinomios[i].getPos() == polinomios[i].length()){condicion = false;}
        }
    }
    for (int i=0; i<3; i++){ 
        polinomios[i].moveToStart();
        bool condicion = true; //condicion para el while
        unsigned int z;
        if (polinomios[i].getValue().x < 0){
            z = polinomios[i].getValue().x; // condicion para agregar los monomios que faltan
        } else {
            z = 0;
        }
        while (condicion){
            unsigned int y1 = polinomios[i].getValue().x, y2 = z;
            if (y1 > y2){ 

                tElemento auxT{y2,0};
                polinomios[i].previous();
                polinomios[i].insert(auxT);
                polinomios[i].next();
            }
            
            polinomios[i].next();
            z++;

            if (polinomios[i].getPos() == polinomios[i].length()){
                if (z >= polinomios[i].getValue().x){
                    condicion = false;
                }
            }
        }
    }
    //for (int i=0; i<totalPolinomios; i++){polinomios[i].print();}; //printea los polinomios
    txt.close();

    txt.open("entradaPolinomio.txt", ifstream::in); //abrimos el txt de nuevo para estudiar los COEFICIENTE y EVALUAR
    ofstream wrt;
    wrt.open("salidaPolinomio.txt"); //abrimos archivo de escritura
    while(getline(txt, str,'\n')){
        /* ---------- COEFICIENTE ---------- */
        if (str[0] == 'C'){ //REALIZA LAS OPERACIONES PARA CALCULAR COEFICIENTE CON LOS DATOS DADOS
            coeficiente.coefC = stoi(str.substr(str.find(' '),str.rfind(' ')));
            coeficiente.nPolCoef = stoi(str.substr(str.rfind(' '),str.size()));

            for (unsigned int i=0; i<polinomios[coeficiente.coefC].length(); i++){
                if (polinomios[coeficiente.coefC].getValue().x == i) {

                    wrt << polinomios[coeficiente.coefC].getValue().y << endl; //escribe el coeficiente encontrado en salidaPolinomio.txt
                };


            };
            
        /* ---------- EVALUAR ---------- */
        } else if (str[0] == 'E'){ //REALIZA LAS OPERACIONES PARA CALCULAR EVALUAR CON LOS DATOS SOLICITADOS
            evaluar.nPolEval = stoi(str.substr(str.find(' '),str.rfind(' ')));
            evaluar.xToEval = stof(str.substr(str.rfind(' '),str.size()));
            tLista polAux = polinomios[evaluar.nPolEval];
            float eval=0; //resultado del algoritmo de horner en float
            for (polAux.moveToEnd(); polAux.getPos()>0;polAux.previous()){ //algoritmo de horner
                eval*= evaluar.xToEval;
                eval+= static_cast<float>(polAux.getValue().y);
            };
            wrt << eval << endl;
        }
    };
    wrt.close();
    txt.close();
    return 0;
    }
    


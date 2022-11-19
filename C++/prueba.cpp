#include <map>
#include <vector>
#include <iostream>

using namespace std;

int main(){
    map<int, vector<int>> lista;

    vector<int> numeros {10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110};

    for(int j=0, i=1; j < 5; i++){
        vector<int> vectorAux;
        for(int aux=0; aux < 5; aux++){
            if ((aux + j) >= numeros.size()){
                break;
            }
            vectorAux.push_back(numeros[j+aux]);
        }
        lista.insert(make_pair(i, vectorAux));
        j += 5;
    }

    for (auto it1 { lista.begin() }; it1 != lista.end(); ++it1) {
        cout << "llave: " << it1->first << endl; 
        cout << "Elementos: ";
        for (auto it2 { it1->second.begin() }; it2 != it1->second.end(); ++it2){    
            cout << *it2 << " ";
        }
        cout << endl;
    }
    return 0;
}
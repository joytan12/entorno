% recibe una lista y la Suma de de estod numeros
% mediante recorcion se va recorriendo la lista y sumando 
% los valores de esta y asignandolo a suma
sumar([], 0).
sumar([L1|L2], Suma):-
    sumar(L2, X),
    Suma is L1 + X.

% recive una lista y aux el cual sera el promedio de la lista
% esta funcion llama a suma para saber cual sera la suma de los numeros de 
% la lista y luego los divide por el largo de la lista
promedio(L1, Aux):-
    sumar(L1, Suma),
    length(L1, X),
    Aux is Suma / X.
    
% recive un entero
% identifica si el numero es par
par(Largo):-
    Largo mod 2 =:= 0.

% recive un entero
% identifica si el numero es impar
inpar(Largo):-
    Largo mod 2 =:= 1.

% recive la lista y la mediana
% la lista se ordena con un sort y se verifica si es de largo par o impar
% si es impar toma el valor de el centro de la lista, si no pasa a la otra mediana
mediana(L1, Mediana):-
    msort(L1, Aux),
    length(L1, X),
    inpar(X),
    Pos is X//2+1,
    nth1(Pos, Aux, Mediana).

% recive la lista y la mediana
% la lista se ordena con un sort y se verifica si es de largo par o impar
% si es par toma los valor de el centro de la lista y se promedian.
mediana(L1, Mediana):-
    msort(L1, Aux),
    length(L1, X),
    par(X),
    Pos1 is X // 2 + 1,
    Pos2 is Pos1 - 1,
    nth1(Pos1, Aux, V1),
    nth1(Pos2, Aux, V2),
    Mediana is (V1 + V2) / 2.

% recive valor1 y valor2
% los compara luego devuelve true si v1 > v2 y false con lo contrario
comparar(V1, V2, Bool):-
    V1 > V2,
    Bool = true.

% recive valor1 y valor2
% los compara luego devuelve true si v1 > v2 y false con lo contrario
comparar(V1, V2, Bool):-
    V1 =< V2,
    Bool = false.

% recive la lista de listas y devuelve la lista respuesta
% recorre la lista de listas y obtiene el promedio y mediana de la lista
% se compara ambos valores y se devuelve true o falso segun se nececita
% luego se agrega el valor a la lista y se sigue el proseso
bondad2([], _).
bondad2([Cabeza|Cola], L1):-
    promedio(Cabeza, Promedio),
    mediana(Cabeza, Mediana),
    comparar(Promedio, Mediana, Bool),
    append(L2, [Bool], L1),
    bondad2(Cola, L2).

% recive la lista de listas y devuelve la lista respuesta
% llama a bondad 2 para tener la lista de respuesta y le aplica un reverse
% antes de entregar la respuesta, ya que bondad2 entrega los datos alreves
bondad(L1, Aux) :-
    bondad2(L1, L),
    reverse(L, Aux).
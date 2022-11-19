ancestro(p2,p1).
ancestro(p3,p1).
ancestro(p4,p1).
ancestro(p5,p1).

ancestro(p2,p10).
ancestro(p3,p10).
ancestro(p4,p10).
ancestro(p6,p10).
ancestro(p7,p10).

ancestro(p5,p2).
ancestro(p10,p2).
ancestro(p6,p2).

ancestro(p5,p3).
ancestro(p9,p3).

ancestro(p6,p4).

ancestro(p8,p5).

ancestro(p7,p6).
ancestro(p9,p6).

ancestro(p8,p7).

ancestro(p1,p9).

% recibe una lista y un nodo
% recorre la lista con recurrcion y verifica que el nodo no esta en la lista
notIn([],_).
notIn([L1|L2], Nodo):-
    \+ L1 = Nodo,
    notIn(L2, Nodo).

% recibe una lista y un aux que sera la lista final
% recorre la lista con recursividad y agrega todos los ancestros de los
% nodos que esta en la lista a una lista unica, da igual si se repiten
agregar([], []).
agregar([Cabeza|Cola], Aux):-
    findall(X, ancestro(X, Cabeza), L1),
    agregar(Cola, R),
    append(R, L1, Aux).

% recibe una lista y lo que va a ser la respuesta
% obtiene una lista y saca la cabeza y la devuelve
ultimoElemento([Cabeza|Cola], End):-
    End = Cabeza.

% recibe una lista de listas, el nodo final y la respuesta.
% es un funcion que toma la ultima lista de vecinos de la lista de listas
% luego crea una lista nueva con todos los vecionos y conprueba si esta el 
% que se busca, si se encuantra se va a aux va al otro aux y a las 
% RES se le asigna la lugitud de la lista de listas
aux(Lista, Final, RES):-
    ultimoElemento(Lista, End),
    agregar(End, L1),
    notIn(L1, Final),
    append([L1], Lista, R),
    aux(R, Final, RES).

aux(Lista, Final, RES):-
    ultimoElemento(Lista, End),
    agregar(End, L1),
    append([L1], Lista, Distancia),
    length(Distancia, RES).

% recive ambos nodos y la repuesta
% seca la primera lista de vecinos del nodo inicial y el la respuesta
% primero se verifica si el final no esta en los primeros vecinos, si esta
% se retorna uno, tambien se verifica si existen vecinos, si no existen
% significa que no hay vecinos y retorna 0
ancestrocidad(Final, Padre, RES):-
    findall(X, ancestro(X, Padre), L1),
    notIn(L1, Final),
    length(L1, X),
    X =\= 0,
    append([L1], [], L2),
    aux(L2, Final, RES).

ancestrocidad(Final, Padre, RES):-
    findall(X, ancestro(X, Padre), L1),
    length(L1, X),
    X =\= 0,
    RES is 1.

ancestrocidad(Final, Padre, RES):-
    RES is -1.
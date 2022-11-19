% recive y el str de la operacion y los dos numeros que se utilizara, y devuelve el resultado
% se recualve la operacion y devuleve el resultado
operador("+", X, Y, RES):-
    RES is X + Y.

operador("-", X, Y, RES):-
    RES is X - Y.

operador("*", X, Y, RES):-
    RES is X * Y.

operador("/", X, Y, RES):-
    RES is X / Y.

operador("^", X, Y, RES):-
    RES is X ^ Y.

operador("mod", X, Y, RES):-
    RES is X mod Y.

% se recive la lista y devulve el resultado
% se obtiene los valores de las 3 pociciones de la lista y se verifica si 
% estos valosres son integer, si es que alguno no es un integer se pasa 
% a las otras mates en donde en problema se resulve con recurcividad
% se identifica si alguno no es integer o si los 2 no son integer
% el el caso de que no es integer se llama a matematicas pasando esa lista
matematica([], 0).
matematica(L1, R):-
    nth0(0, L1, V1),
    nth0(1, L1, V2),
    nth0(2, L1, V3),
    integer(V1),
    integer(V3),
    operador(V2, V1, V3, R).

matematica(L1, R):-
    nth0(0, L1, V1),
    nth0(1, L1, V2),
    nth0(2, L1, V3),
    \+ integer(V1),
    integer(V3),
    matematica(V1, RES),
    operador(V2, RES, V3, R).

matematica(L1, R):-
    nth0(0, L1, V1),
    nth0(1, L1, V2),
    nth0(2, L1, V3),
    integer(V1),
    \+ integer(V3),
    matematica(V3, RES),
    operador(V2, V1, RES, R).

matematica(L1, R):-
    nth0(0, L1, V1),
    nth0(1, L1, V2),
    nth0(2, L1, V3),
    matematica(V1, RES1),
    matematica(V3, RES2),
    operador(V2, RES1, RES2, R).
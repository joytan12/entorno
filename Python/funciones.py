from typing import NamedTuple
import re

class Token(NamedTuple):
    type: str
    value: str
    line: int
    column: int

def crearMatriz(n):
    '''crearMatriz
    n : int
    crea la matriz la cual se untilizara en el juago, retorna la matriz'''
    matriz = []
    for x in range(n):
        linea = []
        for i in range(n):
            linea.append(0)
        matriz.append(linea)
    return matriz
# ------------------------

def tokenize(code):
    '''tokenize
    code : str
    esta funcion recibe linea por linea del archivo buscando los patrones
     ya establecidos en la tarea y guandandolos en token con su
    respectivo grupo, si se encuentra un patron que no pertemece al 
    grupo se guardara como un error, retorna token'''
    tokenSpecification = [
        ('modificar', r'A+|B+'),
        ('movimiento', r'[>|<|U|D]([1-9][0-9]*|0)'),
        ('reset', r'[R|Z]'),
        ('mate', r'[X|Y]([>|<|U|D]([1-9][0-9]*|0))+'),
        ('mostrar', r'[S|L][e|c]'),
        ('parentecis', r'[()]'),
        ('comando', r'[?]'),
        ('error', r'.'), #carecter que representa el error se sintaxis
    ]
    tok_regex = '|'.join('(?P<%s>%s)' % pair for pair in tokenSpecification)
    line_num = 1
    line_start = 0
    for mo in re.finditer(tok_regex, code):
        kind = mo.lastgroup
        value = mo.group()
        column = mo.start() - line_start
        if kind == 'NEWLINE':
            line_start = mo.end()
            line_num += 1
            continue
        elif kind == 'SKIP':
            continue
        yield Token(kind, value, line_num, column)

def comprobar(lista, inicio):
    '''comprobar
    lista : lista con token
    inicio : int
    esta funcion se dedica a comprobar si el comando que comiensa con ?
    esta bien escrito, retorna un True en el caso de que este bien
    escrito y False si esta mal'''
    if lista[inicio+1] == 'movimiento' or lista[inicio+1].type == 'mate' or lista[inicio+1].type == 'modificar' or lista[inicio+1].type == 'reset' or lista[inicio+1].type == 'mostrar':
        return True
    elif lista[inicio+1] == 'comando':
        if lista[inicio].type == 'comando' and lista[inicio+1].type == 'movimiento':
            return comprobar(lista, inicio+1)
        else:
            return False
    else:
        return False

def modificar(matriz, pos, cadena):
    '''modificar
    matriz : lista
    pos : lista
    cadena : str
    resibe una cadena de A o B y se hace la operacion correspondiente en
    la matriz, no retorna nada'''
    for x in cadena:
        if x == 'A':
            matriz[pos[0]%len(matriz)][pos[1]%len(matriz)] += 1
        else:
            matriz[pos[0]%len(matriz)][pos[1]%len(matriz)] -= 1

def movimiento(pos, cadena):
    '''movimiento
    pos : lista
    cadena : str
    recibe una cadena con el la direccion y cuanto se tiene que mover
    se le suma o resta a la pos como undique la direccion, tomamos en
    cuenta que el primer numero de pos es X y el segundo en Y
    '''
    direccion = cadena[0]
    distancia = int(cadena[1:])
    if direccion == 'U':
        pos[0] -= distancia
    elif direccion == 'D':
        pos[0] += distancia
    elif direccion == '<':
        pos[1] -= distancia
    elif direccion == '>':
        pos[1] += distancia
    # print(pos)

def reset(matriz, pos, cadena):
    '''reset
    matriz : lista
    pos : lista
    cadena : str
    se dedica a resetear la matriz o resetar la casilla segun
    lo que indique la letra
    '''
    if cadena == "R":
        matriz[pos[0]%len(matriz)][pos[1]%len(matriz)] = 0
    else:
        for x in range(len(matriz)):
            for i in range(len(matriz)):
                matriz[x][i] = 0
        
def mate(matriz, pos, cadena):
    '''reset
    matriz : lista
    pos : lista
    cadena : str
    se dedica a resetear la matriz o resetar la casilla segun
    lo que indique la letra
    '''
    indicador = cadena[0]
    cadena = cadena[1:]
    newPos = pos
    for x in re.findall(r'(>|<|U|D)([1-9][/d]*|0)', cadena):
        direccion = x[0]
        distancia = int(x[1])
        if direccion == 'U':
            newPos = [newPos[0]-distancia, newPos[1]] 
        elif direccion == 'D':
            newPos = [newPos[0]+distancia, newPos[1]]
        elif direccion == '<':
            newPos = [newPos[0], newPos[1]-distancia]
        elif direccion == '>':
            newPos = [newPos[0], newPos[1]+distancia]
    if indicador == 'X':
        matriz[pos[0]%len(matriz)][pos[1]%len(matriz)] *= matriz[newPos[0]%len(matriz)][newPos[1]%len(matriz)]
    elif matriz[newPos[0]%len(matriz)][newPos[1]%len(matriz)] > 0:
        matriz[pos[0]%len(matriz)][pos[1]%len(matriz)] = matriz[pos[0]%len(matriz)][pos[1]%len(matriz)] / matriz[newPos[0]%len(matriz)][newPos[1]%len(matriz)]

def mostrar(matriz, pos, cadena):
    '''mostrar
    matriz : lista
    pos : lista
    cadena : str
    se dedica a imprimir la matris segun las posibilidades L[e|c] o 
    S[e|c], en donde si la opcion es S se concadena hasta completar la
    matriz y inprimirlo 
    '''
    if cadena[0] == 'L':
        if cadena[1] == 'c':
            if 32 < matriz[pos[0]%len(matriz)][pos[1]%len(matriz)] and matriz[pos[0]%len(matriz)][pos[1]%len(matriz)] < 127:
                print(chr(matriz[pos[0]%len(matriz)][pos[1]%len(matriz)]))
            elif matriz[pos[0]%len(matriz)][pos[1]%len(matriz)] <= 127:
                print('')
        else:
            print(matriz[pos[0]%len(matriz)][pos[1]%len(matriz)], end='')
    else:
        dato = ''
        for x in range(len(matriz)):
            for i in range(len(matriz)):
                if cadena[1] == 'c':
                    if (32 < matriz[x][i] and matriz[x][i] < 127):
                        dato += chr(matriz[x][i])
                    elif matriz[x][i] == 127:
                        print('')
                else:
                    dato += str(matriz[x][i])
        print(dato, end='')
            
def comando(matriz, lista, pocicion, pos):
    '''comando
    matriz : lista
    pos : lista
    lista : lista
    pocicion : int
    esta funcion resbe la lista de comandos y la posicion en donde se 
    comienza el ?, utilizando recurcion se va leyendo el comando hasta 
    que que se termine ?, retorna la pocicion en donde termina dicho 
    comando
    '''
    direccion = lista[pocicion+1].value[0]
    distacia = int(lista[pocicion+1].value[1:])
    if direccion == 'U':
        newPos = [pos[0], pos[1]-distacia] 
    elif direccion == 'D':
        newPos = [pos[0], pos[1]+distacia]
    elif direccion == '<':
        newPos = [pos[0]-distacia, pos[1]]
    elif direccion == '>':
        newPos = [pos[0]+distacia, pos[1]]
    if matriz[newPos[0]%len(matriz)][newPos[1]%len(matriz)] < 0:
        return (pocicion+2)
    elif lista[pocicion+2].type == 'comando':
        return comando(matriz, lista, pocicion+2, pos)
    elif lista[pocicion+2].type == 'movimiento' or lista[pocicion+2].type == 'mate' or lista[pocicion+2].type == 'modificar' or lista[pocicion+2].type == 'reset' or lista[pocicion+2].type == 'mostrar':
        operacion(matriz, pos, lista[pocicion+2], lista, pocicion)
        return (pocicion + 2)
    else:
        return (pocicion + 2)

def operacion(matriz, pos, token, lista, inicio):
    '''operacion
    matriz : lista
    pos : lista
    token : token (class)
    lista : lista
    inicio : int
    esta funcion se dedica a recolectar toda la informacion requerida
    para los comandos, tomando el valor token type, se puede disernir 
    que funcion es util para esta tipo de comando pasandole los datos requeridos, en el caso de la funcion comando retorna una nueva
    pocicion para seguir desde ahi, esto se debe a la naturaleza 
    recursiva de ?
    '''
    if token.type == 'modificar':
        modificar(matriz, pos, token.value)
    elif token.type == 'movimiento':
        movimiento(pos, token.value)
    elif token.type == 'reset':
        reset(matriz, pos, token.value)
    elif token.type == 'mate':
        mate(matriz, pos, token.value)
    elif token.type == 'mostrar':
        mostrar(matriz, pos, token.value)
    elif token.type == 'comando':
        return comando(matriz, lista, inicio, pos)

def operaciones(matriz, pos, lista, star=0, end=-1):
    '''operaciones
    matriz : lista
    pos : lista
    lista : lista
    star : int
    end : int
    esta fincion se encarga se ejecutar los comando de izquierda a derecha, priorizando los parantacicis por eso estan los valores
    star y end, teniendo este orden encuanta llama a la funcion 
    operacion operacion y luego elimina los comandos ejecutados
    '''
    inicio = star
    while star != end:
        # cadena = ''
        # for x in lista:
        #     cadena += x.value
        # print(cadena)
        parentecis = False
        if lista[end].type == 'parentecis':
            parentecis = True
        if lista[star].type == 'comando':
            star = operacion(matriz, pos, lista[star], lista, star)
        else:
            operacion(matriz, pos, lista[star], lista, inicio)
        star += 1
        if star == len(lista):
            end = len(lista)
        
    if star < len(lista):
        operacion(matriz, pos, lista[star], lista, inicio)

    aux = inicio
    if parentecis:
        while inicio <= end:
            del lista[aux]
            inicio += 1
    else:    
        while inicio < end:
            del lista[0]
            inicio += 1
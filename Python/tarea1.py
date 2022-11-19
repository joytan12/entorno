from os import error, truncate
import re
from typing import NamedTuple
from funciones import *

# nombre = input('que test quiere abrir')

pos = [0,0]

archivo = open('ejemplo.txt', 'r')

errores = open('errores.txt', 'w')

primera = int(archivo.readline())

matriz = crearMatriz(primera)

errorEnELarchivo = False
todoMalo = True
i = 2
for linea in archivo:
    acciones = []
    lineasMalas = []
    errorEnLaLinea = True
    for token in tokenize(linea):
        if token.type == 'error':
            errorEnELarchivo = True
            errorEnLaLinea = False
        else:
            acciones.append(token)
    equilibrio = [0, 0]
    if errorEnLaLinea:
        for x in acciones:
            if x.value == '(':
                equilibrio[0] += 1
            if x.value == ')':
                equilibrio[1] += 1
    if equilibrio[0] != equilibrio[1]:
        errorEnLaLinea = False
    for x in range(len(acciones)):
        if acciones[x].type == 'comando' and acciones[x+1].type == 'movimiento':
            errorEnLaLinea = comprobar(acciones, x+1)
        elif acciones[x].type == 'comando':
            errorEnLaLinea = False
            errorEnELarchivo = True
    if errorEnLaLinea:
        todoMalo = False
        i = 0
        star = -1
        end = -1
        while len(acciones):
            if acciones[i].value == '(':
                star = i
            if acciones[i].value == ')':
                end = i
            if end > 0 and star >= 0:
                operaciones(matriz, pos, acciones, star, end)
                i = -1
                star = -1
                end = -1
            elif (i+1) == len(acciones):
                operaciones(matriz, pos, acciones)
            i += 1
    else:
        errores.write(str(i) + ' ' + linea)
    i +=1
if todoMalo:
    errores.write('No hay lineas correctas :c')
elif errorEnELarchivo == False:
    errores.write('No hay errores!')
errores.close()
archivo.close()
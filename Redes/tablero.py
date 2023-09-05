'''
ideas para la comunicacion, el cliento, servidor y bot tendran un tablero y solo se mandaran las cordenas
y lo hiran actualizando cada uno.
''' 

def dibujarTablero(tablero):
    for _ in range(6):
        print(' ' ,_, '', end='')
    print('')
    for fila in tablero:
        for casilla in fila:
            if casilla == 0:
                print("   ", end="|")
            elif casilla == 1:
                print(" X ", end="|")
            elif casilla == 2:
                print(" O ", end="|")
        print("\n" + "---+" * 6)

def verificarGanador(tablero):
    # Verificar filas
    for fila in tablero:
        for i in range(3):
            if fila[i] == fila[i + 1] == fila[i + 2] == fila[i + 3] != 0:
                return fila[i]

    # Verificar columnas
    for col in range(6):
        for i in range(3):
            if tablero[i][col] == tablero[i + 1][col] == tablero[i + 2][col] == tablero[i + 3][col] != 0:
                return tablero[i][col]

    # Verificar diagonales hacia arriba
    for i in range(3):
        for j in range(3, 6):
            if tablero[i][j] == tablero[i + 1][j - 1] == tablero[i + 2][j - 2] == tablero[i + 3][j - 3] != 0:
                return tablero[i][j]

    # Verificar diagonales hacia abajo
    for i in range(3):
        for j in range(3):
            if tablero[i][j] == tablero[i + 1][j + 1] == tablero[i + 2][j + 2] == tablero[i + 3][j + 3] != 0:
                return tablero[i][j]

    return 0  # Si no hay ganador

def jugarEnColumna(columna, tablero, jugador):
    # Verificar que la columna sea válida
    if columna < 0 or columna > 5:
        print("Columna inválida. Debe ser un número entre 0 y 5.")
        return False

    # Encontrar la fila disponible en la columna
    fila_disponible = -1
    for fila in range(5, -1, -1):
        if tablero[fila][columna] == 0:
            fila_disponible = fila
            break

    # Si no se encontró una fila disponible en la columna, está llena
    if fila_disponible == -1:
        print("La columna está llena. Elige otra columna.")
        return False

    # Realizar el movimiento del jugador en la columna
    tablero[fila_disponible][columna] = jugador
    return True

tablero = [
    [0, 0, 0, 0, 0, 0],
    [0, 1, 0, 0, 0, 0],
    [0, 0, 1, 0, 0, 0],
    [0, 0, 0, 1, 2, 0],
    [0, 0, 0, 1, 1, 0],
    [0, 0, 2, 2, 2, 0]
]
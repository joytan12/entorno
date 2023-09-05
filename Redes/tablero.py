def dibujar_tablero(tablero):
    for fila in tablero:
        print("|", end='')
        for casilla in fila:
            if casilla == 0:
                print("   ", end="|")
            elif casilla == 1:
                print(" X ", end="|")
            elif casilla == 2:
                print(" O ", end="|")
        print("\n" + "---+" * 6)

def verificar_ganador(tablero):
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

tablero = [
    [0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0],
    [0, 0, 0, 1, 2, 0],
    [0, 0, 0, 1, 1, 0],
    [0, 0, 2, 2, 2, 0]
]

dibujar_tablero(tablero)
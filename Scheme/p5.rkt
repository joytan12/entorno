#lang scheme

; racket p5.rkt para compilar
; lista recibe una lista y un numero
; esta funcion devuelve el valor de una pocicion la cual indica el numero dado
; devuelve un numero
(define (lista ls cont) 
    (if (= cont 1)
        (car ls)
    ;else
        (lista (cdr ls) (- cont 1))
    )
)

; length recive una lista
; con recurcion cuanta cuantos elementos tiene la lista
; devuelve el largo de la lista
(define (length ls)
    (let rc([ls ls] [cont 0])
        (if (null? ls)
            cont
        ;else
            (rc (cdr ls) (+ cont 1))
        )
    )
)

; busqueda recibe un contador lista nodo grafo
; se dedica a buscar en el nodo entregado en el grafo
; devuelve una lista con el nodo y sus vecinos
(define (busqueda i aux nodo grafo)
    (if (<= i (length grafo))
        (if (= (car (lista grafo i)) nodo)
            (append aux (append (list (car (lista grafo i))) (cadr (lista grafo i))))
        ;else
            (busqueda (+ i 1) aux nodo grafo)
            ; (displayln "recurcion")
        )
    ;else
        aux
    )
)

; vecinos lista grafo inicio
; recorre la lista y por cada vecino del contagio se busca y guarda los vecinos de los vecinos
; devuelve una lista con los vecinos de los vecinos
(define (vecinos ls grafo inicia)
    (let rc ([aux ls] [gr grafo] [i inicia])
        (if (< i (length ls))
            (rc (append aux (cdr (busqueda 1 '() (lista aux (+ i 1)) gr))) gr (+ i 1) )
        ;else
            aux
        )
    )
)
; repetidos lista n cantidad
; se intruduce la lista, el numero que se avelue y un 0 para inicializar y si se repite la condicion que devuelve false
; retorna true si no se repite y false si se repite
(define (repetidos lista n cantidad)
    (cond 
        [(null? lista)
            (cond 
                [(= cantidad 1)
                    #f
                ]

                [else
                    #t
                ]
            )
        ]

        [else
            (cond 
                [(= (car lista) n)
                    (repetidos (cdr lista) n (+ cantidad 1))
                ]

                [else
                    (repetidos (cdr lista) n cantidad)
                ]
            )
        ]
    )
)
; respuesta lista
; se recorre la lista y los numeros que no se repiten se agregan a una lista aux la cual devuelve
; devuelve una lista sin repetidos
(define (respuesta lista)
    (let rC ([aux lista] [aux2 '()])
        (cond
            [(null? aux)
                aux2
            ]

            [(eq? #t (repetidos aux2 (car aux) 0))
                (rC (cdr aux) (append aux2 (list(car aux))))
            ]

            [else
                (rC (cdr aux) aux2)
            ]
        )
    )
)


(define (contagio grafo n d)
    (let rc ([gr grafo] [inicia 1] [ciclo d] [aux (busqueda 1 '() n grafo)])
        (if (> ciclo 1)
            (rc gr (length aux) (- ciclo 1) (vecinos aux gr inicia))
        ;else
            (respuesta aux)
        )
    )
)

(contagio '((2 (1 3 4)) (1 (2)) (3 (2)) (4 (2))) 2 3)
; (contagio '((2 (5 6)) (5 (2 10)) (10 (5)) (6 (2 11 8)) (8 (3 6 1)) (3 (8)) (1 (8)) (11 (6 15 4)) (15 (11)) (4 (11))) 2 2)
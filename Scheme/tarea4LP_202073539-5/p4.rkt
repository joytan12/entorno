#lang scheme

;;vida h que es el nodo que hay que buscar y el arbol
;;el arbol se va recorrien tomando en cuanto 4 estados en los cuales se puede saber si se visito derecha y lo que eso segnifica eso, se va agregando todos los nodos al camino y se van a eliminar los nodos que no corresponde a los padres del nodo buscado
;;devuelve la lista con los antecesores

(define (vida h arbol)
    (let recorrer ([aux arbol] [estado 0] [padres '()] [camino arbol]) 
        (let ([n (car aux)] [l (cadr aux)] [r (caddr aux)])            
            (cond 
                [(or [= n h] [null? arbol] [null? aux])
                    (reverse padres)
                ]

                [(= estado 4)
                    (if (pair? (cadr camino))
                        (recorrer (cadr camino) 2 (cdr padres) (cdr camino) )

                    ;else
                        (recorrer (cdr camino) 2 (cdr padres) (cdr camino) )
                    )
                ]

                [(and [not (null? l)] (or [= estado 0] [= estado 3]))
                    (recorrer l 0 (append [list (car aux)] padres) (append (list l) camino) )
                ]

                [(and [null? l] [null? r] [not (= n h)] [= estado 0] )
                    (recorrer (cadr camino) 2 (cdr padres) (cdr camino) )
                ]

                [(and [not (null? r)])
                    (recorrer r 3 (append [list (car aux)] padres) (append (list r) camino) )
                ]

                [(and [null? l] [null? r] [not (= n h)])
                    (recorrer (cadr camino) 4 (cdr padres) (cdr camino) )
                ]
            )
        )
    )
)

(vida 2 '(5 (3 (2 () ()) (4 () ())) (8 (7 () ()) ())))
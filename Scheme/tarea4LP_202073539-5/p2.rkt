#lang scheme

;;tranformacion recibe las dos funciones lambda y una lista de numeros
;;se hace una recurcion simple hasta llegar al final de la lista original, en cada llamado recurcivo se hece la aplicacion de las funciones como dice el enunciado y se comparan agregando a una lista aux el mayor
;;devuelve la lista con los mayores entre las convinaciones de ambas funciones

(define (transformacion funcion1 funcion2 numeros)
    (let rc ([ls numeros] [aux '()])
        (cond 
        [(null? ls)
        aux
        ]
        [else 
        (let solo ([r1 (funcion2 (funcion1 (car ls)))] [r2 (funcion1 (funcion2 (car ls)))])
                (cond
                [(null? ls)
                    aux
                ]
                [(> r1 r2)
                    (rc (cdr ls) (append aux (list r1)) )
                ]
                [(< r1 r2)
                    (rc (cdr ls) (append aux (list r2)) )
                ]
                [else
                    (rc (cdr ls) (append aux (list r1)) )
                ]
                )
            )
        ]
        )
    )
)

(transformacion (lambda (x) (+ 2 x)) (lambda (x) (/ x 2)) '(2 3 4))
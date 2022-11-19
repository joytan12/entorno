#lang scheme

;;mazo recive la lista y el divisor 
;;se va a ir sacando los elementos de la lista e ir comprobando si cumplen la condicion de ser divisor para agregarlos a la nueva lista la cua sera retornada
;;devuelve la lista con todos los nueros que son divididos por el paramtro ya entregado

(define mazo
    (lambda (carta div)
        (let recurs ([ls carta] [aux '()])
            (cond
            [(null? ls)
                aux
            ]
            [(= (modulo (car ls) div) 0)
                (recurs (cdr ls) (append aux (list (car ls))))
            ]
            [else
                (recurs (cdr ls) aux)
            ]
            )
        )
    )
)

(mazo '(1 2 3 4 5) 3)

(mazo '(1 2 3 4 5) 2)
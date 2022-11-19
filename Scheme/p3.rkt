#lang scheme

;;zeta_simple y zeta_cola reciven dos numeros que seran utilzados en la ecuacion definida en el enunciado
;;zeta_simple suma el resultado de la operacion con el futuro operacion de la sumatoria, zeta_cola 
;;ambos devuelven el resultado de la sumatoria

(define (zeta_cola i s)
    (let rc ([suma 0] [j i])
        (cond 
        [(= 0 j)
            suma
        ]
        [else
            (rc (+ suma (/ 1.0 (expt j s))) (- j 1) )
        ]
        )
    )
)

(define (zeta_simple i s)
    (if (= i 0)
        0
    ;else
        (+ (/ 1.0 (expt i s)) (zeta_simple (- i 1) s) )
    )    
)

(zeta_simple 3 2)

(zeta_cola 3 2)
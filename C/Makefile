test: simulacion.o animal.o
	gcc -Wall simulacion.o animal.o -o test -lm && ./test

test2: simulacion.o animal.o;
	clean
	gcc -Wall simulacion.o animal.o -o test -lm
	(echo 1; echo 0; echo 0; echo e; echo 14; echo c; echo a; echo e; echo 3; echo 1; echo 1; echo 1; echo 0; echo 2; echo e; echo 6; echo c; echo r; echo f; echo 3.34; echo 1; echo 2;echo 2; echo -1;) | ./test

simulacion.o: simulacion.c animal.h
	gcc -c simulacion.c

animal.o: animal.c animal.h
	gcc -c animal.c

.PHONY: clean
clean:
	rm -f *.o test
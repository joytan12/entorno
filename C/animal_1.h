typedef struct animal{
    void* fuerza;
    char tipoFuerza;
    void* velocidad;
    char tipoVelocidad;
    void* resistencia;
    char tipoResistecia;
    void (*reproduccion)(struct animal* , struct animal* , struct animal* );
    void (*comerHuir)(struct animal* ,struct animal* );
} animal;

void crearAnimal(animal* a);
void borrar(animal* a);
void mostrarAnimal(animal* a);
void reproducir(animal* padre, animal* madre, animal* hijo);
void comeOhuir(animal* a1, animal* a2);
int comparar(animal* a1, animal* a2);
void mostrarMundo(animal** mundo);
void borrarMundo(animal** mundo);
void reproduccionSimple(animal* padre, animal* madre, animal* hijo);
void reproduccionCruzada(animal* padre, animal* madre, animal* hijo);
void comerSiempre(animal* a1, animal* a2);
void huirSiempre(animal* a1, animal* a2);
void comerAleatorio(animal* a1, animal* a2);
// void iteracion(animal** mundo);
// void moverAnimal(animal** mundo, animal* a1, animal* a2)
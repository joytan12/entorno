import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class Data extends Thread {
    private int valor;
    private MatrixFileReader reader;

    public Data(int valor, MatrixFileReader reader) {
        this.valor = valor;
        this.reader = reader;
    }

    @Override
    public void run() {
        if (valor == 0 && (reader.getDim() != reader.getAux())){
            test0();
        } else if(reader.getDim() == reader.getAux()) {
            Resultado dim = new Resultado(reader.getDim());
            busqueda(dim);
            // busca el la matriz original 
        } else {
            // aqui se hace la busqueda especializada
            Resultado dim = new Resultado(reader.getDim());
            // Scanner scanner = new Scanner(System.in);
            int aux = valor, div = 0;
            while (aux > 0){
                aux = aux / 10;
                div ++;
            }
            if (dim.getLimX()  / (div * 2) == reader.getAux()){
                //si entro aqui hay buscar si el resoro esta en las dimeciones
                parametrizador(invertirNumero(valor), dim);
                busqueda(dim);
            } else {
                test1(valor);
            }
        }
    }

    public void setValor(int valor){
        this.valor = valor;
    }

    public int invertirNumero(int numero) {
        int numeroInvertido = 0;

        while (numero != 0) {
            int digito = numero % 10;
            numeroInvertido = numeroInvertido * 10 + digito;
            numero /= 10;
        }

        return numeroInvertido;
    }

    public void parametrizador(int valor, Resultado dim){
        
        int aux1 = (dim.getLimX() - dim.getX())/2, aux2 = (dim.getLimY() - dim.getY())/2, tmp = valor % 10;
        valor = valor/10;

        if (tmp == 1 || tmp == 3){
            dim.setLimX(dim.getLimX() - aux1);
        } 
        if (tmp == 1 || tmp == 2){
            dim.setLimY(dim.getLimY() - aux2);
        }
        if (tmp == 4 || tmp == 2){
            dim.setX(dim.getX() + aux1);
        }
        if (tmp == 4 || tmp == 3){
            dim.setY(dim.getY() + aux2);
        }
        
        if (valor > 0){
            parametrizador(valor, dim);
        }
    }

    public void test0(){
        Data thread1 = new Data(1, reader);
        Data thread2 = new Data(2, reader);
        Data thread3 = new Data(3, reader);
        Data thread4 = new Data(4, reader);

        // Iniciar la ejecución del hilo
        thread1.start();
        thread2.start();
        thread3.start();
        thread4.start();
    }

    public void test1(int valor){

        Data thread1 = new Data((valor * 10) + 1, reader);
        Data thread2 = new Data((valor * 10) + 2, reader);
        Data thread3 = new Data((valor * 10) + 3, reader);
        Data thread4 = new Data((valor * 10) + 4, reader);

        // Iniciar la ejecución del hilo
        thread1.start();
        thread2.start();
        thread3.start();
        thread4.start();
    }

    public void busqueda(Resultado dim){
        // aqui va el algoritmo de busqueda en la matriz
        for (int i = dim.getX(); i < dim.getLimX(); i++){
            for (int j = dim.getY(); j < dim.getLimY() ; j++){
                if (reader.getMatriz()[i][j] == 1){
                    System.out.println("La encontramos en la columna " + (i + 1));
                    break;
                }
            }
        }
    }
}
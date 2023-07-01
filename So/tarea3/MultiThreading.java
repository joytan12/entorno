import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class MultiThreading extends Thread {
    private int valor;
    private MatrixFileReader reader;

    public MultiThreading(int valor, MatrixFileReader reader) {
        this.valor = valor;
        this.reader = reader;
    }
    
    long h = System.nanoTime();
    @Override
    public void run() {
        //corre las hebras
        if (valor == 0 && (reader.getDim() != reader.getAux())){
            test0();
        } else if(reader.getDim() == reader.getAux()) {
            Resultado dim = new Resultado(reader.getDim());
            busqueda(dim);
            // busca en la matriz original 
        } else {
            // aqui se hace la busqueda especializada
            Resultado dim = new Resultado(reader.getDim());
            int aux = valor, div = 0;
            while (aux > 0){
                aux = aux / 10;
                div ++;
            }
            if (dim.getLimX()  / (div * 2) == reader.getAux()){
                //si entro aqui hay buscar si el resoro esta en las dimensiones
                parametrizador(invertirNumero(valor), dim);
                busqueda(dim);
            } else {
                test1(valor);
            }
        }
        
    }

    public int invertirNumero(int numero) {
        //invierte un numero dado
        int numeroInvertido = 0;

        while (numero != 0) {
            int digito = numero % 10;
            numeroInvertido = numeroInvertido * 10 + digito;
            numero /= 10;
        }

        return numeroInvertido;
    }

    public void parametrizador(int valor, Resultado dim){
        //genera los subcuadrantes
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
        //primera division de la matriz original
        MultiThreading thread1 = new MultiThreading(1, reader);
        MultiThreading thread2 = new MultiThreading(2, reader);
        MultiThreading thread3 = new MultiThreading(3, reader);
        MultiThreading thread4 = new MultiThreading(4, reader);

        // Iniciar la ejecución del hilo
        thread1.start();
        thread2.start();
        thread3.start();
        thread4.start();
    }

    public void test1(int valor){
        //divisiones de las submatrices
        MultiThreading thread1 = new MultiThreading((valor * 10) + 1, reader);
        MultiThreading thread2 = new MultiThreading((valor * 10) + 2, reader);
        MultiThreading thread3 = new MultiThreading((valor * 10) + 3, reader);
        MultiThreading thread4 = new MultiThreading((valor * 10) + 4, reader);

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
                    System.out.println("fila " + (i + 1) + ", columna " + (j + 1));
                    long durHebra = (System.nanoTime() - h)/1000000;
                    System.out.println("Hebras: " + durHebra + "ms");
                    break;
                }
            }
        }
    }
}
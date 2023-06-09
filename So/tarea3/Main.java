import java.io.Reader;
import java.io.Reader;
import java.util.*;

public class Main 
{
    public static void main(String[] args) {
        
        String filePath = "ubicacion-tesoro.txt";
        MatrixFileReader reader = new MatrixFileReader(filePath);
        int aux = reader.getDim();
        int[][] matriz = reader.getMatriz();

        Data thread = new Data(0, reader);
        thread.start();
        Fork procesos = new Fork(0, reader);
        Resultado dim = new Resultado(aux);
        procesos.padre(dim);
    }
}
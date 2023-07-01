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

        MultiThreading thread = new MultiThreading(0, reader);
        
        thread.start();

        Fork procesos = new Fork(0, reader);
        Resultado dim = new Resultado(aux);
        
        procesos.padre(dim);

        //corresponde a la ejecucion del programa sin fork ni thread
        long d = System.nanoTime();
        for (int i = 0; i < aux; i++){
            for (int j = 0; j < aux ; j++){
                // if (matriz[i][j] == 1){
                    // System.out.println("fila " + (i + 1) + ", columna " + (j + 1));
                    // break;
                // }
                System.out.print(matriz[i][j] + " ");
            }
            System.out.print("\n");
        }
        long durNing = (System.nanoTime() - d)/1000000;
        System.out.println("Ninguno de los anteriores: " + durNing + "ms");
    }
}
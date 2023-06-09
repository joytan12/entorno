import java.util.concurrent.ForkJoinPool;  
import java.util.concurrent.RecursiveTask;

import javax.swing.plaf.DimensionUIResource;

import java.io.IOException;

public class Fork {
    private int valor;
    private MatrixFileReader reader;

    public Fork(int valor, MatrixFileReader reader){
        this.valor = valor;
        this.reader = reader;
    }

    public void padre(Resultado dim){
        if ((dim.getLimX() - dim.getX()) > reader.getAux()){
            try {
                for (int i = 1; i <= 4; i++) {
                    ProcessBuilder processBuilder = new ProcessBuilder("java", "-cp", "ruta_de_clases", "Proceso", Integer.toString(i));
                    Process proceso = processBuilder.start();
                    System.out.println(proceso);
                    if (proceso.isAlive()){
                        hijo(i, dim);
                    }
                }
            } catch (IOException e) {
                e.printStackTrace();
            }
        } 
        // else {
            // aqui va el codigo de busqueda ayuda
        // }
    }

    public void hijo(int valor, Resultado aux){
        Resultado dim = new Resultado(reader.getDim());
        dim.setX(aux.getX());
        dim.setY(aux.getY());
        dim.setLimX(aux.getLimX());
        dim.setLimY(aux.getLimY());
        parametrizador(valor, dim);
        padre(dim);
    }

    public void parametrizador(int valor, Resultado dim){
        
        int aux1 = (dim.getLimX() - dim.getX())/2, aux2 = (dim.getLimY() - dim.getY())/2;

        if (valor == 1 || valor == 3){
            dim.setLimX(dim.getLimX() - aux1);
        } 
        if (valor == 1 || valor == 2){
            dim.setLimY(dim.getLimY() - aux2);
        }
        if (valor == 4 || valor == 2){
            dim.setX(dim.getX() + aux1);
        }
        if (valor == 4 || valor == 3){
            dim.setY(dim.getY() + aux2);
        }
    }
}

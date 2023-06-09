import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class MatrixFileReader {
    private int dim;
    private int aux;
    private int[][] matriz;

    public MatrixFileReader(String filePath) {
        try {
            File file = new File(filePath);
            Scanner scanner = new Scanner(file);

            // Leer las dimensiones de la matriz
            dim = Integer.parseInt(scanner.nextLine());

            // Leer el valor auxiliar
            aux = Integer.parseInt(scanner.nextLine());

            // Leer la matriz
            String matrixLine = scanner.nextLine();
            String[] matrixValues = matrixLine.split(",");
            matriz = new int[dim][dim];
            int index = 0;
            for (int i = 0; i < dim; i++) {
                for (int j = 0; j < dim; j++) {
                    String value = matrixValues[index++]
                        .replace("[", "")
                        .replace("]", "")
                        .replace("x", "1")
                        .trim();
                    matriz[i][j] = Integer.parseInt(value);
                }
            }

            scanner.close();
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }
    }

    public int getDim() {
        return dim;
    }

    public int getAux() {
        return aux;
    }

    public int[][] getMatriz() {
        return matriz;
    }
}
public class Resultado {
    private int x;
    private int limX;
    private int y;
    private int limY;

    public Resultado(int valor) {
        this.x = 0;
        this.limX = valor;
        this.y = 0;
        this.limY = valor;
    }

    // Métodos getter y setter para acceder a los valores
    public int getX() {
        return x;
    }

    public void setX(int x) {
        this.x = x;
    }

    public int getLimX() {
        return limX;
    }

    public void setLimX(int limX) {
        this.limX = limX;
    }

    public int getY() {
        return y;
    }

    public void setY(int y) {
        this.y = y;
    }

    public int getLimY() {
        return limY;
    }

    public void setLimY(int limY) {
        this.limY = limY;
    }
}

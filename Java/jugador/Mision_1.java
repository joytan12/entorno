package jugador;
public class Mision 
{
    // atributos 
    private char requisito;
    private int valor;
    private int cantidad;
    private int recompensa;

    // contrucctor
    public Mision(char var1, int var2, int var3, int var4){
        this.requisito = var1;
        this.valor = var2;
        this.cantidad = var3;
        this.recompensa = var4;
    }

    // metodos
    public boolean verificarRequisito(){
        /*se dedica a verificar si las condiciones de las miesiones
        son cumplidas todo esto diferenciado si son de matar o 
        de llegar a algun lugar */
        boolean condicion = false;
        if(cantidad == valor && requisito == 'v'){
            condicion = true;
        } else if (cantidad >= valor && requisito == 'm'){
            condicion = true;
        }
        return condicion;
    }
    // getter
    public char getRequisito(){
        return this.requisito;
    }

    public int getValor(){
        return this.valor;
    }

    public int getRecompensa(){
        return this.recompensa;
    }

    public int getCantidad(){
        return this.cantidad;
    }

    // setter
    public void setRequisito(char var){
        this.requisito = var;
    }

    public void setValor(int var){
        this.valor = var;
    }

    public void setRecompensda(int var){
        this.recompensa = var;
    }

    public void setCantidad(int var){
        this.cantidad = var;
    }
}
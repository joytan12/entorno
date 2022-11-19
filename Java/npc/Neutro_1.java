package npc;
import java.util.*;
import jugador.Jugador;
import jugador.Mision;

public class Neutro extends NPC{
    // atributo
    private char requisito;
    private int valor;
    private int recompensa;
    private boolean yaDioRecompensa;
    // contructor
    public Neutro(char requisito, int valor, int recompensa, boolean estado){
        this.requisito = requisito;
        this.valor = valor;
        this.recompensa = recompensa;
        this.yaDioRecompensa = estado;
    }
    // metodos
    public void interaccion(Jugador pj){
        /*realiza la interaccion con en jugador */
        if (yaDioRecompensa){
        String msg = String.format("< %s >: ya te di mision, no intentes cuentirme larva", this.nombre);
        System.out.println(msg);
        } else {
            String var1, var2;
            if (this.requisito == 'v'){
                var1 = "llegar a";
                var2 = "del mundo";
            } else {
                var1 = "matar";
                var2 = "de Monstruo";
            }

            String msg = String.format("< %s >: hola, quieres cumplir con esta mision? Deberas < %s > < %d > < %s > y recibiras < %d > de xp", this.nombre, var1, this.valor, var2, this.recompensa);

            System.out.println(msg);

            Scanner sc = new Scanner(System.in);
        
            System.out.println("(1) aceptar\n(2) rechazar");

            int respuesta = sc.nextInt();

            if (1 == respuesta){
                // agregar el metodo
                Mision aux = new Mision(this.requisito, this.valor, 0, this.recompensa);
                pj.agregarMision(aux);
                this.yaDioRecompensa = true;
            }
        }
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

    public boolean getYaDioRecompensa(){
        return this.yaDioRecompensa;
    }
    // setter
    public void setRequisito(char var){
        this.requisito = var;
    }

    public void setValor(int var){
        this.valor = var;
    }

    public void setRecompensa(int var){
        this.recompensa = var;
    }

    public void setYaDioRecompensa(Boolean var){
        this.yaDioRecompensa = var;
    }
}

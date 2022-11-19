package npc;
import jugador.Jugador;

public class Malo extends NPC{
    // atributo
    private int cantidadEnergia;
    private int cantidadMana;

    // constructor
    public Malo(String var1, int var2, int var3){
        this.nombre = var1;
        this.cantidadEnergia = var2;
        this.cantidadMana = var3; 
    }

    // metodos
    public void interaccion(Jugador pj){
        /*realiza la interaccion con en jugador */
        String msg = String.format("< %s >: SOY MALO TE VOY A DISMINUIR TU ENERGIA y MANA EN < %d > y < %d >!", this.nombre, this.cantidadEnergia, this.cantidadMana);
        System.out.println(msg);
        int numero = (int)(Math.random()*10+1);
        if (numero <= 6){
            pj.modEnergia(this.cantidadEnergia);
        }
        numero = (int)(Math.random()*10+1);
        if (numero <= 6){
            pj.modMana(this.cantidadMana);
        }
    }
    // getter
    public int getCantidadEnergia(){
        return this.cantidadEnergia;
    }

    public int getCantidadMana(){
        return this.cantidadMana;
    }

    // setter
    public void setCantidadEnergia(int var){
        this.cantidadEnergia = var;
    }

    public void setCantidadMana(int var){
        this.cantidadMana = var;
    }
}
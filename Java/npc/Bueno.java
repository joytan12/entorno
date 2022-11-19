package npc;
import jugador.Jugador;

public class Bueno extends NPC{
    // atributo
    private String atributo;
    private int cantidad;
    // constructor
    public Bueno(String var1, String var2, int var3){
        this.nombre = var1;
        this.atributo = var2;
        this.cantidad = var3;
    }
    // metodos
    public void interaccion(Jugador pj){
        /*realiza la interaccion con en jugador */
        String msg = String.format("< %s >: Creo que necesitas un poco de ayuda te subire < %s > a tu < %d >!!!!", this.nombre, this.atributo, this.cantidad);
        System.out.println(msg);
        if (this.atributo.equals("vida")){
            pj.modVida(this.cantidad);
        } else if (this.atributo.equals("energia")){
            pj.modEnergia(this.cantidad);
        } else if (this.atributo.equals("mana")){
            pj.modMana(this.cantidad);
        } else if (this.atributo.equals("xp")){
            pj.subirNivel(this.cantidad);
        }
    }
    // getter
    public String getAtributo(){
        return this.atributo;
    }

    public int getCantidad(){
        return this.cantidad;
    }
    // setter
    public void setAtributo(String var){
        this.atributo = var;
    }

    public void setCantidad(int var){
        this.cantidad = var;
    }
}
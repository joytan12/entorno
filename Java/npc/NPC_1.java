package npc;
import jugador.Jugador;

public abstract class NPC 
{
    // atributo
    protected String nombre;

    // metodos
    public abstract void interaccion(Jugador pj);

    // getter
    public String getNombre(){
        return this.nombre;
    }
    // setter
    public void setNombre(String var){
        this.nombre = var;
    }
}
package tierra;
import npc.NPC;
import jugador.Jugador;
import enemigo.*;

public abstract class Tierra 
{
    // atributo
    protected float probabilidadEnemigo;
    protected Monstruo monstruo;
    protected JefeFinal jefeFinal;
    protected NPC npc;

    // motodos abstracto
    public abstract boolean accion(Jugador pj);

    // getter
    public float getProbabilidadEnemigo(){
        return this.probabilidadEnemigo;
    }

    public Monstruo getMonstruo(){
        return this.monstruo;
    }

    public JefeFinal getJefeFinal(){
        return this.jefeFinal;
    }

    public NPC getNpc(){
        return this.npc;
    }

    // setter
    public void setprobabilidadEnemigo(float var){
        this.probabilidadEnemigo = var;
    }

    public void setMonstruo(Monstruo var){
        this.monstruo = var;
    }

    public void setJefeFinal(JefeFinal var){
        this.jefeFinal = var;
    }

    public void setNpc(NPC var){
        this.npc = var;
    }
}
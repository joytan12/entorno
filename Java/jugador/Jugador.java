package jugador;
import java.util.*;

public abstract class Jugador 
{
    // atributos
    protected String nombre;
    protected int vida;
    protected int vidaMax;
    protected int xp;
    protected int fuerza;
    protected int inteligencia;
    protected int energia;
    protected int energiaMax;
    protected int mana;
    protected int manaMax;
    protected ArrayList<Mision> listaMisiones = new ArrayList<Mision>();
    protected int nivelActual;

    // metodos abstractos
    public abstract int ataque();    
    public abstract int hechizo();
    public abstract void subirNivel(int exp);

    // metodos
    public void mostrosJugador(){
        /*esta funcion se dedica a dar por consola los datos del 
        jugagor */
        String smj = String.format("tu nombre <%s>\ntu vada es<%d>/<%d>\ntu experiencia<%d>\ntu fuerza es <%d>\ntu inteligencia<%d>\ntu energia<%d>/<%d>\ntu nivel actual <%d>", this.nombre, this.vida, this.vidaMax, this.xp, this.fuerza, this.inteligencia, this.energia, this.energiaMax, this.nivelActual);
        System.out.println(smj);
    }

    public void modVida(int var){
        /*
        int var 
        modifica la vida y la mantiene es sus parametro segun
        el entero entregado*/
        this.vida += vida; 
        if (this.vida < 0){
            this.vida = 0;
        } else if (this.vidaMax > 0){
            this.vida = vidaMax;
        }
    }

    public void modEnergia(int var){
        /*
        int var 
        modifica la energia y la mantiene es sus parametro segun
        el entero entregado*/
        this.energia += var;
        if (this.energia < 0){
            this.energia = 0;
        } else if (this.energiaMax > 0){
            this.energia = energiaMax;
        }
    }

    public void modMana(int var){
        /*
        int var 
        modifica la mana y la mantiene es sus parametro segun
        el entero entregado*/
        this.mana += var;
        if (this.mana < 0){
            this.mana = 0;
        } else if (this.manaMax > 0){
            this.mana = manaMax;
        }
    }

    public void agregarMision(Mision aux){
        /*se agrega una mision en particular*/
        this.listaMisiones.add(aux);
    }

    public Mision getMision(int i){
        /*se obtiene una miesion en espesifico de la lista*/
        return listaMisiones.get(i);
    }

    public void verificarMsion(){
        /*recorre la lista de misiones verificando si alguna de estas 
        se a cumplido, les da la respectiva exp al jugador y despues las remueve de la lista*/
        for(int i=0; i<this.listaMisiones.size(); i++){
            if (listaMisiones.get(i).verificarRequisito() == true){
                subirNivel(listaMisiones.get(i).getRecompensa());
                listaMisiones.remove(i);
                i--;
            }
        }
    }

    public void avanzarMision(int j, char verificar){
        /*
        int j 
        char requisito
        esta funcion se dedicara a sumar a las misiones y despues llamar
        a las verificacion para ver si se comlieron
        no retorna nada 
        */
        for(int i=0; i<this.listaMisiones.size(); i++){
            if (verificar == 'v'){
                listaMisiones.get(i).setCantidad(j + listaMisiones.get(i).getCantidad());    
            } else {
                listaMisiones.get(i).setCantidad(j + listaMisiones.get(i).getCantidad());    
            }
        }
        verificarMsion();
    }
    // getter
    public String getNombre(){
        return this.nombre;
    }

    public int getVida(){
        return this.vida;
    }

    public int GetVidaMax(){
        return this.vidaMax;
    }

    public int getXp(){
        return this.xp;
    }

    public int GetFuerza(){
        return this.fuerza;
    }

    public int getInteligencia(){
        return this.inteligencia;
    }

    public int getEnergia(){
        return this.energia;
    }

    public int getEnergiaMax(){
        return this.energiaMax;
    }

    public int getMana(){
        return this.mana;
    }

    public int getManaMax(){
        return this.manaMax;
    }

    public List<Mision> getMisiones(){
        return this.listaMisiones;
    }

    public int getNivel(){
        return this.nivelActual;
    }
    
    // setter
    public void setNombre(String var){
        this.nombre = var;
    }

    public void setVida(int var){
        this.vida = var;
    }

    public void setVidaMax(int var){
        this.vidaMax = var;
    }

    public void setXp(int exp){
        this.xp = exp;
    }

    public void setFuerza(int var){
        this.fuerza = var;
    }

    public void setInteligencia(int var){
        this.inteligencia = var;
    }

    public void setEnergia(int var){
        this.energia = var;
    }

    public void setEnergiaMax(int var){
        this.energiaMax = var;
    }

    public void setmMana(int var){
        this.mana = var;
    }

    public void setManaMax(int var){
        this.manaMax = var;
    }

    public void setMana(int var){
        this.mana = var;
    }

    public void setNIvel(int var){
        this.nivelActual = var;
    }
}
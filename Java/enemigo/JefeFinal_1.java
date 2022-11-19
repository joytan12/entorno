package enemigo;

import jugador.Jugador;
import java.util.*;

public class JefeFinal implements Enemigo{
    // atributos
    private String nombre;
    private int vida;
    private int vidaMax;
    private int danoBase;
    //  constructor
    public JefeFinal(String nombre, int vida, int dano){
        this.nombre = nombre;
        this.vida = vida;
        this.vidaMax = vida;
        this.danoBase = dano;
    } 
    // metodos
    public void combate(Jugador pj) {
        /*esta funcion se dedica a desarrollar la batalla contra 
        el jefe final */
        System.out.println("te topaste con un rector elige tu ataque para vencer a la u");
        Scanner sc = new Scanner(System.in);
        if (vidaMax/2 >= vida && vida != 0){
            System.out.println("(1) ataque \n (2) hechizo");
            int ataque = sc.nextInt();
            if (ataque == 1){
                this.vida -= pj.ataque();    
            } else {
                this.vida -= pj.hechizo();
            }
            pj.modVida(-(this.danoBase + 2 * 1));
        }else if (this.vida > 0 && pj.getVida() > 0){
            System.out.println("(1) ataque \n (2) hechizo");
            int ataque = sc.nextInt();
            if (ataque == 1){
                this.vida -= pj.ataque();    
            } else {
                this.vida -= pj.hechizo();
            }
            pj.modVida(-(this.danoBase + 2 * 2));
        } else {
            System.out.println("El rector te siega con su pela y se escapa siguelo hasta matarlo");
        }
    }
    // getter
    public String getNombre(){
        return this.nombre;
    }

    public int getVida(){
        return this.vida;
    }

    public int getVidaMax(){
        return this.vidaMax;
    }

    public int getDanoBase(){
        return this.danoBase;
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

    public void setDanoBase(int var){
        this.danoBase = var;
    }
}

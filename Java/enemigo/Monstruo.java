package enemigo;

import jugador.Jugador;
import java.util.*;

public class Monstruo implements Enemigo{
    // atributo
    private int vida;
    private int dano;
    // constructor
    public Monstruo(int vida, int dano){
        this.vida = vida;
        this.dano = dano;
    } 
    // atributo
    public void combate(Jugador pj) {
        /*se dedica a generar el combate entre en jugador
        y el monstruo*/
        while (this.vida > 0 && pj.getVida() > 0){
            System.out.println("te topaste con un ayudante elige tu ataque");
            System.out.println("(1) ataque \n(2) hechizo");
            Scanner sc = new Scanner(System.in);
            int ataque = sc.nextInt();
            if (ataque == 1){
                setVida(vida - pj.ataque());   
            } else {
                setVida(vida - pj.hechizo());
            }
            pj.modVida(-this.dano);
        }
        if (pj.getVida() > 0){
            pj.avanzarMision(1, 'm');
        }
    }
    // getter
    public int getVida(){
        return this.vida;
    }

    public int getDano(){
        return this.dano;
    }
    // setter
    public void setVida(int var){
        this.vida = var;
    }

    public void setDano(int var){
        this.dano = var;
    }
}
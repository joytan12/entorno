package tierra;

import enemigo.JefeFinal;
import enemigo.Monstruo;
import jugador.*;
import npc.NPC;

public class Bosque extends Tierra {
    // constructor
    public Bosque(float probabilidad, Monstruo aux, JefeFinal temp, NPC var){
        this.probabilidadEnemigo = probabilidad;
        this.monstruo = aux;
        this.jefeFinal = temp;
        this.npc = var;
    }
    // metodos
    public boolean accion(Jugador pj){
        /*verifica los el estado del jugador y ejecuta
        las acciones en orden en el que deberia*/
        System.out.println("bienvenido al bosque los ongos te dan la bienvenida");
        if (pj.getMana() > 0){
            if(pj.getMana() < 3){
                int i = 3 - pj.getMana();
                pj.modVida(-i);
            }
            pj.modMana(-3);
        }
        if (pj.getVida() > 0){
            npc.interaccion(pj);
            // probabilidad se que aparezca tu mama
            int numero = (int)(Math.random()*100+1);
            float n = (float) numero / 100;
            if (n <= probabilidadEnemigo){
                monstruo.combate(pj);
                jefeFinal.combate(pj);
            }
            if (pj.getVida() == 0) return false;
            return true;
        }
        return false;
    }
}
package tierra;

import jugador.*;
import enemigo.*;
import npc.*;

public class Motana extends Tierra {
    // constructor
    public Motana(float probabilidad, Monstruo aux, JefeFinal temp, NPC var){
        this.probabilidadEnemigo = probabilidad;
        this.monstruo = aux;
        this.jefeFinal = temp;
        this.npc = var;
    }
    // metodos
    public boolean accion(Jugador pj){
        System.out.println("llegaste a la montaña");
        /*verifica los el estado del jugador y ejecuta
        las acciones en orden en el que deberia*/
        if (pj.getEnergia() > 0){
            if(pj.getEnergia() < 3){
                int i = 3 - pj.getEnergia();
                pj.modVida(-i);
            }
            pj.modEnergia(-3);
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

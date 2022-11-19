package tierra;

import jugador.*;
import enemigo.*;
import npc.*;

public class Planicie extends Tierra{
    // constructor
    public Planicie(float probabilidad, Monstruo aux, JefeFinal temp, NPC var){
        this.probabilidadEnemigo = probabilidad;
        this.monstruo = aux;
        this.jefeFinal = temp;
        this.npc = var;
    }
    // metodos
    public boolean accion(Jugador pj){
        System.out.println("ves un terreno libre mas adelante, una planicie");
        /*verifica los el estado del jugador y ejecuta
        las acciones en orden en el que deberia*/
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
}

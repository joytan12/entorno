package jugador;

public class Druida extends Jugador {
    // Constructor
    public Druida(String var){
        this.nombre = var;
        this.vida = 15;
        this.vidaMax = 15;
        this.xp = 0;
        this.fuerza = 5;
        this.inteligencia = 5;
        this.energia = 5;
        this.energiaMax = 5;
        this.mana = 5;
        this.manaMax = 5;
        this.nivelActual = 0;
    }
    
    // metodos
    public int ataque(){
        /* se dedica a calcular el daño */
        int danno, max;
        if (energia/3 > mana/2){
            max = energia/3;
        } else {
            max = mana/2;
        }
        danno = (fuerza + inteligencia)/3 * max;
        modEnergia(-3);
        return danno;
    }    

    public int hechizo(){
        /* se dedica a calcular el daño */
        int danno, max;
        if (energia/2 > mana/3){
            max = energia/2;
        } else {
            max = mana/3;
        }
        danno = (fuerza + inteligencia)/3 * max;
        modMana(-3);
        return danno;
    }

    public void subirNivel(int exp){
        this.xp += exp;
        int nivel = 0;
        if (this.xp >= 10){
            nivel = 1;
        } else if (this.xp >= 25){
            nivel = 2;
        } else if (this.xp >= 50){
            nivel = 3;
        } else if (this.xp >= 100){
            nivel = 4;
        } else if (this.xp >= 200){
            nivel = 5;
        } else if (this.xp >= 350){
            nivel = 6;
        } else if (this.xp >= 600){
            nivel = 7;
        } else if (this.xp >= 900){
            nivel = 8;
        }
        if (nivel > this.nivelActual){
            System.out.println("Subiste de nivel pequeño sansano");
            this.vidaMax += nivel;
            this.fuerza += nivel;
            this.inteligencia += nivel;
            this.energiaMax += nivel;
            this.manaMax += nivel;
            this.vida = vidaMax;
            this.energia = energiaMax;
            this.mana = manaMax;
            this.nivelActual = nivel;
        }
    }
}

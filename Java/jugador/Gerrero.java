package jugador;

public class Gerrero extends Jugador {
    // Constructor
    public Gerrero(String var){
        this.nombre = var;
        this.vida = 20;
        this.vidaMax = 20;
        this.xp = 0;
        this.fuerza = 9;
        this.inteligencia = 1;
        this.energia = 10;
        this.energiaMax = 10;
        this.mana = 2;
        this.manaMax = 2;
        this.nivelActual = 0;
    }
    
    // metodos
    public int ataque(){
        /* se dedica a calcular el daño */
        int danno;
        danno = this.fuerza*2+this.energia;
        modEnergia(-5);
        return danno;
    }    

    public int hechizo(){
        /* se dedica a calcular el daño */
        int danno;
        danno = this.inteligencia*(this.fuerza/4)+this.mana;
        modMana(-3);
        return danno;
    }

    public void subirNivel(int exp){
        System.out.println("Subiste de nivel pequeño sansano");
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
            this.vidaMax += 3 * nivel;
            this.fuerza += 5 * nivel;
            this.inteligencia += 1;
            this.energiaMax += 2 * nivel;
            this.manaMax += 1;
            this.vida = vidaMax;
            this.energia = energiaMax;
            this.mana = manaMax;
            this.nivelActual = nivel;
        }
    }
}

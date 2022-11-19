package jugador;
public class Mago extends Jugador {
    // Constructor
    public Mago(String var){
        this.nombre = var;
        this.vida = 10;
        this.vidaMax = 10;
        this.xp = 0;
        this.fuerza = 3;
        this.inteligencia = 10;
        this.energia = 6;
        this.energiaMax = 6;
        this.mana = 15;
        this.manaMax = 15;
        this.nivelActual = 0;
    }
    
    // metodos
    public int ataque(){
        /* se dedica a calcular el daño */
        int danno;
        danno = this.fuerza*(this.inteligencia/4)+this.energia;
        modEnergia(-3);
        return danno;
    }    

    public int hechizo(){
        /* se dedica a calcular el daño */
        int danno;
        danno = this.inteligencia*(this.fuerza/4)+this.mana;
        modMana(-5);
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
            this.vidaMax += 2 * nivel;
            this.fuerza += 1;
            this.inteligencia += 3 * nivel;
            this.energiaMax += 1;
            this.manaMax += 3 * nivel;
            this.vida = vidaMax;
            this.energia = energiaMax;
            this.mana = manaMax;
            this.nivelActual = nivel;
        }
    }
}
import java.util.*;

import javax.security.sasl.Sasl;

import jugador.*;
import tierra.*;
import npc.*;
import enemigo.*;

public class Main 
{
    public static void main(String[] args) {
        
        Scanner sc = new Scanner(System.in);
        
        System.out.println("tamaño del mundo");

        int tamañoMundo = sc.nextInt();

        List<Tierra> mundo = new ArrayList<Tierra>(tamañoMundo);

        /* creacion del mundo */
        for (int xd=0; xd<tamañoMundo; xd++){
            int tipoT, tipoE, i;
            float probabilidad; 
            String nombre;

            System.out.println("Que tipo de tierra quieres\n (1) montaña\n (2) planicie\n (3) bosque");
            tipoT = sc.nextInt();
            System.out.println("Que probabilidad quiere que tenga de aparecer enemigo");
            probabilidad = sc.nextFloat();
            System.out.println("Que tipo de enemigo quiere\n (1) Monstruo\n (2) Jefe Final.");
            tipoE = sc.nextInt();
            Monstruo aux; 
            JefeFinal finDeSemestre;
            // creacion de los enemigos
            if (tipoE == 1){
                int vida, dano;
                System.out.println("imgrese el vida y la daño");
                vida = sc.nextInt();
                dano = sc.nextInt();                
                aux = new Monstruo(vida, dano);
                finDeSemestre = new JefeFinal("ayudante generico n3", 0, 0);
            } else {
                int vida, dano;
                String apodo;
                System.out.println("como se llamara al jefe final");
                apodo = sc.next();
                System.out.println("imgrese el vida y la daño");
                vida = sc.nextInt();
                dano = sc.nextInt();                
                finDeSemestre = new JefeFinal(apodo, vida, dano);
                aux = new Monstruo(0, 0);
            }
            System.out.println("habra un NPC\n(1) si \n(2) no");
            tipoE = sc.nextInt();
            NPC temp;
            if (tipoE == 1){
                System.out.println("el NPC sera\n (1) bueno \n(2) malo\n (3) neutro");
                i = sc.nextInt();
                System.out.println("Que nombre tendra el NPC");
                nombre = sc.next();
                if (i == 1){
                    String atributo;
                    int cantidad;
                    System.out.println("que atributo tendra:");
                    atributo = sc.next();
                    System.out.println("que cantidad tendra:");
                    cantidad = sc.nextInt();
                    temp = new Bueno(nombre, atributo, cantidad);
                } else if (i == 2){
                    int cantidadEnergia, cantidadMana;
                    System.out.println("cuanto energia tendra:");
                    cantidadEnergia = sc.nextInt();
                    System.out.println("cuanto mana tendra:");
                    cantidadMana = sc.nextInt();
                    temp = new Malo(nombre, cantidadEnergia, cantidadMana);
                } else {
                    int valor, recompensa;
                    char requisito;
                    System.out.println("que requisito:\n (1) v \n(2) m");
                    i = sc.nextInt();
                    if (i == 1){
                        requisito = 'v';
                    } else {
                        requisito = 'm';
                    }
                    System.out.println("que valor tendra:");
                    valor = sc.nextInt();
                    System.out.println("que recompensa tendra:");
                    recompensa = sc.nextInt();
                    temp = new Neutro(requisito, valor, recompensa, false);
                }
            } else {
                temp = new Bueno("belisaurio", "xp", 0);
            }
            if (tipoT == 1){
                Tierra aux2 = new Motana(probabilidad, aux, finDeSemestre, temp);
                mundo.add(aux2);
            } else if (tipoT == 2){
                Tierra aux2 = new Planicie(probabilidad, aux, finDeSemestre, temp);
                mundo.add(aux2);
            } else {
                Tierra aux2 = new Bosque(probabilidad, aux, finDeSemestre, temp);
                mundo.add(aux2);
            }
        }
        // menu para el jugador
        String i;
        Jugador sansano;
        System.out.println("cual va a ser su nombre de dungeon master:");
        i = sc.next();
        System.out.println("que clases elegiras:\n (1) Mago \n (2) Guerrero\n (3) Druida");
        int j = sc.nextInt();
        if (j == 1){
            sansano = new Mago(i);
        } else if(j == 2){
            sansano = new Gerrero(i);
        } else {
            sansano = new Druida(i);
        }
        System.out.println("en que casilla se va a enpezar");
        j = sc.nextInt();
        Mision z = new Mision('m', 2, 2, 2);
        sansano.agregarMision(z);
        while(sansano.getVida() > 0){
            int opcion;
            String msj = String.format("tu posicion actual es < %d > en un mundo de < %d > largo", j+1 , tamañoMundo);
            System.out.println(msj);
            System.out.println("que deseas hacer:\n (1) mover izquierda \n (2) mover derecha\n (3) ver atributos de jugador");
            opcion = sc.nextInt();
            
            if (opcion == 1){
                j--;
                if (j < 0){
                    j = tamañoMundo - 1;
                }
                if (mundo.get(j).accion(sansano)){
                    sansano.avanzarMision(-1, 'v');
                }
            } else if(opcion == 2){
                if (mundo.get(j).accion(sansano)){
                    // System.out.println("entro a la wea");
                    j++;
                    if (j >= tamañoMundo){
                        j = 0;
                    }
                    sansano.avanzarMision(1, 'v');
                }
            } else if(opcion == 3) {
                sansano.mostrosJugador();
            } else {
                sansano.setVida(0);
            }
            if (sansano.getVida() == 0){
                System.out.println("game over, a casa puta");
            }
            if (sansano.getVida() > 0){
                sansano.setVida(0);
                System.out.println("felicidades benciste al rector y saliste de la u");
            }
        }
    }
}
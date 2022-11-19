FLAGS = -g
JC = javac
.SUFFIXES: .java .class
.java.class:
	$(JC) $(FLAGS) $*.java

CLASSES = \
	Main.java \

default: classes

classes: $(CLASSES:.java=.class)

legendOfMorio: classes
	clear
	java Main

MUNDOSIZE = echo 1;
MUNDO0 = echo 2; echo 1; echo 1; echo 8; echo 3; echo 1; echo 1; echo panda; echo xp; echo 800; echo joytan; echo 1; echo 0; echo 1; echo 1;  #creacion de mundo
# MUNDO1 = 

test1: classes
	clear
	($(MUNDOSIZE) \
	$(MUNDO0) \
	) | java Main

clean:
	$(RM) *.class
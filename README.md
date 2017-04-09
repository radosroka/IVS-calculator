# IVS-calculator
IVS Kalkulacka

Preklad:

make build
make run-examle alebo make run-calc

Potrebny gcc, make, go, qmake(qt4)

Skusane na fedore, qmake moze robit problem na ubuntu ak sa binarka nevola qmake alebo qmake-qt4.
 - tomas skusal na ubuntu a ide mu to len to chce rucne nainstalovat nove go(>=1.7.5)

# QT bindings:

v src/github.com/visualfc/goqt/ je cele to goqt
 - bin/ skompilovane binarky
 - examples/ nejake priklady s ktorych sa vieme ispirovat ako sa s tym robi je tam aj nejaky kalkulator
 - ui/ go "api" pre QT

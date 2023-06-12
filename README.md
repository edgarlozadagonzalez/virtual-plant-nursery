# virtual-plant-nursery-go

Proyecto Sprint 1

Se debe realizar una aplicación software en el lenguaje GO (Golang) de lado del servidor que disponga de 
servir una API que permita consumir una serie de servicios relacionados a continuación y abordando el 
siguiente problema:
Se quiere realizar una aplicación para un vivero que permita simular el comportamiento y crecimiento de 
una planta a medida que se va alimentando y cuidando de diferentes formas, esto con el objetivo de 
probar de manera experimental sin afectar plantas reales. Todas las plantas tienen propiedades distintas 
y reaccionan diferente a las condiciones a las que se sometan. Por ejemplo, una rosa requiere 2 horas bajo 
el sol directamente y la orquídea requiere 4 horas, ambas duraciones deben ser exactas porque si duran
menos tienden a pudrirse o si duran más, tienden a secarse y mientras más se alejen de su valor ideal más 
peligro tienen de morir. 
Cada planta tiene propiedades como: 
- Grado de supervivencia (GS): Se refiere a un margen de tolerancia a partir de su valor ideal donde 
se toma un porcentaje y se calcula con relación a cada ítem siguiente y permite que la planta no 
muera durante un tiempo. Si se sobrepasa este límite entonces la planta muere. Unidad 
porcentaje (%). Se muestra un ejemplo luego de definir las propiedades.
- Cantidad de agua requerida (AR): Es un valor que representa la cantidad necesaria para que la 
planta tenga un 100% de agua en el sistema y esté completamente satisfecha. Unidad litros (l).
- Cantidad de agua en el sistema (AS): Es un valor de 0 a 100%. Donde el valor va disminuyendo en 
función del tiempo. Este valor disminuye un 1/X por cada unidad de tiempo transcurrida. Donde 
X es grado de hidratación de la planta. El 100% estaría dado por cantidad de agua requerida.
- Grado de hidratación de la planta (GH): Es un valor entre 20 y 100, donde si éste es más bajo 
entonces la planta se deshidrata más rápido y si es mayor, lo hace más lentamente, por lo que 
resiste más tiempo sin requerir nuevamente agua.
- Cantidad de nutrientes requeridos (NR): Es un valor que representa la cantidad necesaria para 
que la planta tenga un 100% de nutrientes en el sistema y esté completamente satisfecha. Unidad 
miligramos (mg).
- Cantidad de nutrientes en el sistema (NS): Es un valor de 0 a 100%. Donde el valor va 
disminuyendo en función del tiempo. Este valor disminuye un 1/Y por cada unidad de tiempo 
transcurrida. Donde Y es grado de nutrición de la planta. El 100% estaría dado por cantidad de 
nutrientes requeridos.
- Grado de nutrición de la planta (GN): Es un valor entre 60 y 100, donde si éste es más bajo 
entonces la planta pierde nutrientes más rápido y si es mayor, lo hace más lentamente, por lo que 
resiste más tiempo sin requerir nuevamente nutrientes.
Ejemplo de Grado de supervivencia: Se tiene una planta con las siguientes características: GS: 20%., AR: 
0.2 l., AS: 100%., GH: 50., NR: 12 mg., NS: 100%, GN: 60.
Caso 1: En este caso, la planta tiene toda el agua requerida y lo mismo los nutrientes. Sin embargo, no se 
le vuelve a agregar agua ni nutrientes. Entonces cada segundo, las cantidades en el sistema de la planta 
disminuyen de la siguiente manera, la planta pierde 1/50 % (0.02%) de agua y 1/60 % (aprox. 0.016%) de 
nutrientes cada segundo. Esto significa que se quedará sin agua más rápido que quedarse sin nutrientes. 
Entonces al transcurrir 5000 segundos (83.3 minutos aprox. o 1.38 horas aprox.) (esto porque se calcula 
en relación con el 100% de agua al momento) la planta tendrá 0% de agua en su sistema. La planta no 
morirá inmediatamente por que tiene un GS del 20%, esto significa que podrá resistir un poco sin agua y 
así permitir agregarle agua. El 20% permitirá que la planta resista como si tuviera un 20% de agua en su 
sistema todavía, pero con el doble de la pérdida, es decir, 2*(1/50) % por lo que tendría 500 segundos (8. 
3 minutos aprox.) para para planta no muriera por este factor.
Caso 2: En este caso, se le agrega un 30% más agua a la planta, aunque tenga el 100% de AS actualmente
(dado que 100% equivale en este caso 0.2 litros, entonces 30% sería 0.06 litros). Es decir, que la planta 
tiene más agua de la que debería, por lo que se hace el mismo cálculo de 2*(1/50) % para calcular el 
tiempo requerido para que la planta pueda estar en 100% y no por encima, si en ese tiempo no logra llegar 
a este valor, la planta se muere por cantidad excesiva de agua. En este caso, la planta tiene 500 segundos 
(8.3 minutos aprox.) para perder el exceso de agua (este cálculo es con el GS definido). Sin embargo, la 
cantidad extra agregada es de 0.06 litros que es un 30% extra de agua, lo que le tomaría realmente 750 
segundos (12.5 minutos). En consecuencia, la planta no alcanza a sobrevivir.
Nota: La información a continuación planteada es inventada para el ejercicio y no significa que el 
comportamiento y biología de las plantas sea así necesariamente.
La aplicación podrá tener varios usuarios, pero todos con un solo tipo rol, no habrá más roles (no 
administradores o cosas así). El estado de todas las plantas cambiará por cada unidad de tiempo 
(segundos) transcurrida.
La API/REST debe cumplir con lo siguiente:
1. Permitir la realización de todos los CRUD pertinentes.
2. Ver el estado de una planta en cuestión.
3. Ver el estado de todas las plantas creadas.
4. Permitir la agregación de agua tanto por una planta de manera individual o grupal. Unidad litros.
5. Permitir la agregación de nutrientes tanto por una planta de manera individual o grupal. Unidad 
miligramos.
La aplicación también deberá simular si una planta muere por deshidratación y/o pérdida de nutrientes
con una alerta previa cuando esté en riesgo extremo de esto. En cada unidad de tiempo se debe registrar 
los cambios en la base de datos.
El estudiante debe proponer e implementar una solución software que cumpla mínimo con lo aquí 
establecido. Si dado el caso, no logra terminar todo, se debe enviar, presentar y sustentar lo realizado. Se 
evaluará la creatividad, la solución propuesta, funcionalidad, buenas prácticas (por ejemplo, código limpio 
y comentado), la lógica y dominio del tema.
Condiciones de entrega: 
Fecha máxima: lunes 12 de junio de 2023, hora 10:00 am.
Asunto: Sprint 0 – Go.
Archivo: El proyecto comprimido en un .rar, no enlaces o comparticiones. Recuerde no agregar ningún .exe 
dentro del .rar. ¡Éxitos

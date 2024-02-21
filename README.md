Este proyecto es una implementación del algoritmo de búsqueda en anchura **(Breadth-First Search, BFS)**
para resolver el **rompecabezas del 8-puzzle**. 

**main.go**: 
Este es el punto de entrada del programa. Aquí se genera el estado inicial del rompecabezas, se verifica 
si es solucionable y luego se llama a la función `BreadthFirstSearch` para encontrar la solución. 
Si se encuentra una solución, se imprime paso a paso.

**node.go**: Este archivo define la estructura de un nodo en el árbol de búsqueda.
 Cada nodo contiene un estado (la disposición actual de las piezas del rompecabezas), una acción 
 (el movimiento que llevó a este estado) y un enlace al nodo padre. También se definen varias funciones para 
 manipular y verificar los nodos.

**queue.go**: Este archivo implementa una cola, que es una estructura de datos necesaria para el 
algoritmo BFS. La cola se utiliza para mantener un registro de los nodos que aún no se han explorado.

**El algoritmo BFS** funciona explorando todos los nodos de un nivel antes de pasar al siguiente. En el contexto 
de este rompecabezas, esto significa que primero se prueban todos los estados que se pueden alcanzar en un solo
movimiento, luego todos los estados que se pueden alcanzar en dos movimientos, y así sucesivamente hasta que se encuentra
una solución o se han probado todos los estados posibles.

La función `isSolvable` verifica si el rompecabezas tiene solución. No todos los estados
iniciales del 8-puzzle son solucionables. La solvabilidad se determina contando las inversiones en el rompecabezas. 
Una inversión es cualquier par de fichas con valores en el orden incorrecto. Si el número de inversiones es impar, 
el rompecabezas no tiene solución.

La función `GenerateSuccessors` genera todos los estados que se pueden alcanzar a partir de un estado dado realizando 
un solo movimiento. Estos estados se añaden a la cola para ser explorados.

La función `BreadthFirstSearch` implementa el algoritmo BFS. Comienza con el estado inicial y
explora todos los estados sucesores, añadiéndolos a la cola. Cuando se encuentra un estado que es el estado objetivo,
el algoritmo se detiene y devuelve el nodo correspondiente. Este nodo contiene un enlace a su nodo padre, por lo que se puede 
rastrear el camino desde el estado inicial al estado objetivo.

Puedes ver el video de demostración [aquí](https://drive.google.com/file/d/1-QztEqwCj1JzsB3i8UnELwXUIy44w3kP/view?usp=sharing).
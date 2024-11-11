# Simple Blockchain

Una implementación simple de Blockchain en Go.

### Cómo ejecutar el proyecto
Será necesario descargar el fichero `SimpleBlockchain.exe` y ejecutarlo. Al hacerlo se abrirá una ventana de Terminal. A continuación, abre una nueva pestaña, o una nueva ventana, y sigue las siguientes instrucciones (o pasos). En caso de querer realizar otras operaciones con los usuarios existentes (user1 y user2) solo habría que modificar los valores correspondientes.

##### Creación de nuevas transacciones:

*El usuario 1 hace una transferencia al usuario 2:*

    curl -Method Post -Uri http://localhost:8080/transaction -Headers @{"Content-Type"="application/json"} -Body '{"from":"user1","to":"user2","amount":10}'

*El usuario 2 hace una transferencia al usuario 1:*

    curl -Method Post -Uri http://localhost:8080/transaction -Headers @{"Content-Type"="application/json"} -Body '{"from":"user2","to":"user2","amount":15}'

##### Ver transacciones pendientes:

    curl http://localhost:8080/pending

Para observar con claridad cuáles son las transferencias pendientes también es posible abrir la url en el navegador: `http://localhost:8080/pending`.

##### Minería de un bloque:

    curl -Method Post -Uri http://localhost:8080/mine -Headers @{"Content-Type"="application/json"} -Body '{"address":"miner1"}'

##### Obtención de los bloques existentes:

    curl http://localhost:8080/blocks

En caso de querer observar con mayor claridad cuáles son los bloques existentes, tenemos la posibilidad de abrir la url en el navegador: `http://localhost:8080/blocks`.

##### Obtención del saldo de una cuenta:

*Se obtiene el saldo del usuario 1:*

    curl http://localhost:8080/balance/user1

*Se obtiene el saldo del usuario 2:*

    curl http://localhost:8080/balance/user2

Si se quisiera ver mejor cuáles son los saldos de cada uno de los usuarios, se podría colocar solo la url en el navegador: `http://localhost:8080/balance/user1` y `http://localhost:8080/balance/user2`.

Para finalizar el programa solo es necesario cerrar la ventana del Terminal.

***

### Futuras implementaciones
- Añadir validación de la cadena completa.
- Implementar ajuste dinámico de dificultad.
- Agregar límite de transacciones por bloque.
- Implementar sistema de nodos y consenso.
- Agregar firmas digitales para las transacciones.
# ⛓ Blockchain en Go ⛓

Una implementación de una blockchain básica en Go diseñada para demostrar los principios fundamentales de una cadena de bloques, como transacciones, minería, y verificación de saldos. Este proyecto sirve como una introducción al desarrollo de blockchain y proporciona una base sólida para expandir sus capacidades.

## ✅ Requisitos Previos

- **Go** (versión 1.16 o superior): Puedes descargarlo desde [golang.org](https://go.dev/dl/).
- **cURL** (opcional): Para hacer solicitudes HTTP desde la terminal, aunque también puedes usar herramientas como Postman o directamente el navegador.

## 📂 Cómo ejecutar el proyecto

1. **Clona este repositorio en tu entorno local:**
   ```bash
   git clone https://github.com/rcp-code/SimpleBlockchain
   ```

2. **Accede a la carpeta raíz del proyecto y compila el código:**
   ```bash
   cd SimpleBlockchain
   go build cmd/main.go
   ```

   Esto generará un archivo ejecutable `main.exe` (o `main` en Linux/Mac) que puedes ejecutar para interactuar con la blockchain.

## 🔛 Uso del Proyecto

Con el proyecto ejecutándose, puedes realizar las siguientes acciones mediante `curl` o desde tu navegador web.

### 1. Crear nuevas transacciones:

**Transferencias entre usuarios**.

- *El usuario 1 hace una transferencia al usuario 2:*
```bash
curl -Method Post -Uri http://localhost:8080/transaction -Headers @{"Content-Type"="application/json"} -Body '{"from":"user1","to":"user2","amount":10}'
```

- *El usuario 2 hace una transferencia al usuario 1:*
```bash
curl -Method Post -Uri http://localhost:8080/transaction -Headers @{"Content-Type"="application/json"} -Body '{"from":"user2","to":"user2","amount":15}'
```

### 2. Ver transacciones pendientes

Consulta las transacciones que están pendientes de ser minadas:

```bash
curl http://localhost:8080/pending
```

**Respuesta esperada:**
```json
[
    {"from": "user1", "to": "user2", "amount": 10},
    {"from": "user2", "to": "user1", "amount": 15}
]
```

También puedes visualizar las transacciones pendientes en `http://localhost:8080/pending`.

### 3. Minar un bloque

Procesa las transacciones pendientes y agrega un bloque a la cadena:

```bash
curl -Method Post -Uri http://localhost:8080/mine -Headers @{"Content-Type"="application/json"} -Body '{"address":"miner1"}'
```

### 4. Ver la cadena de bloques

Obtén la lista de todos los bloques de la cadena:

```bash
curl http://localhost:8080/blocks
```

Puedes acceder desde tu navegador en `http://localhost:8080/blocks` para ver el historial de bloques.

### 5. Consultar el saldo de una cuenta

Verifica el saldo de un usuario:

- **Saldo de Usuario 1:**
   ```bash
   curl http://localhost:8080/balance/user1
   ```

   **Respuesta esperada:**
   ```json
   {"user": "user1", "balance": 25}
   ```

O abre la URL en tu navegador (`http://localhost:8080/balance/user1`) para ver los saldos.

## ❌ Posibles Errores

Si encuentras problemas, revisa los siguientes puntos:

- **Puerto en uso:** Asegúrate de que el puerto `8080` esté libre antes de ejecutar el proyecto.
- **Compilación fallida:** Revisa que Go esté correctamente instalado y configurado en tu `PATH`.

***

## 🛠 Futuras Mejoras

A continuación, algunos planes de desarrollo que pueden llevar este proyecto a un siguiente nivel:

- **Validación completa de la cadena:** Asegurar la integridad de todos los bloques.
- **Ajuste dinámico de la dificultad de minería:** Adaptar la dificultad basada en la tasa de generación de bloques.
- **Límite de transacciones por bloque:** Implementar un tamaño máximo de transacciones por bloque.
- **Sistema de nodos y consenso:** Permitir la comunicación entre múltiples nodos para lograr consenso.
- **Firmas digitales:** Añadir autenticación de transacciones mediante criptografía.

***

## ⁉ Dudas y sugerencias

Si tienes preguntas o necesitas ayuda para ejecutar el proyecto, ¡estoy aquí para ayudarte! Puedes abrir una *issue* en el repositorio describiendo tu problema o pregunta, y responderé lo antes posible. Si además tienes sugerencias para mejorar el proyecto, no dudes en compartirlas.

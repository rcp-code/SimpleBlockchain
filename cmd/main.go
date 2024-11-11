package main

import (
	"fmt"
	"net/http"

	"github.com/rcp-code/SimpleBlockchain/core"
)

func main() {
	// Crear blockchain
	bc := core.CreaBlockchain()

	// Crear algunas cuentas de ejemplo
	bc.CreaCuenta("miner1", 200.0)
	bc.CreaCuenta("user1", 50.0)
	bc.CreaCuenta("user2", 135.0)

	// Crear y configurar servidor
	servidor := core.NuevoServidor(bc)
	fmt.Println("Servidor blockchain iniciado en http://localhost:8080")

	fmt.Println("Endpoints disponibles:")
	fmt.Println("- GET  /blocks            -> Ver todos los bloques")
	fmt.Println("- POST /transaction       -> Crear nueva transacción")
	fmt.Println("- POST /mine              -> Minar un nuevo bloque")
	fmt.Println("- GET  /balance/{address} -> Ver saldo de una cuenta")
	fmt.Println("- GET  /pending           -> Ver transacciones pendientes")

	http.ListenAndServe(":8080", servidor.Router)
}

/*
# Obtener todos los bloques
curl http://localhost:8080/blocks

# Crear nueva transacción
curl -Method Post -Uri http://localhost:8080/transaction -Headers @{"Content-Type"="application/json"} -Body '{"from":"user1","to":"user2","amount":10}'

curl -Method Post -Uri http://localhost:8080/transaction -Headers @{"Content-Type"="application/json"} -Body '{"from":"user2","to":"user2","amount":15}'


# Minería de un bloque
curl -Method Post -Uri http://localhost:8080/mine -Headers @{"Content-Type"="application/json"} -Body '{"address":"miner1"}'

curl -Method Post -Uri http://localhost:8080/mine -Headers @{"Content-Type"="application/json"} -Body '{"address":"user2"}'

# Obtener saldo de una cuenta
curl http://localhost:8080/balance/user1

curl http://localhost:8080/balance/user2
*/

package main

import (
	"fmt"
	"net/http"

	"github.com/rcp-code/SimpleBlockchain/core"
)

func main() {
	// Crear blockchain
	blockchain := core.CreaBlockchain()

	// Crear algunas cuentas de ejemplo
	blockchain.CreaCuenta("miner1", 200.0)
	blockchain.CreaCuenta("user1", 50.0)
	blockchain.CreaCuenta("user2", 135.0)

	// Crear y configurar servidor
	servidor := core.NuevoServidor(blockchain)
	fmt.Println("Servidor blockchain iniciado en http://localhost:8080")

	fmt.Println("Endpoints disponibles:")
	fmt.Println("- GET  /blocks            -> Ver todos los bloques")
	fmt.Println("- POST /transaction       -> Crear nueva transacción")
	fmt.Println("- POST /mine              -> Minar un nuevo bloque")
	fmt.Println("- GET  /balance/{address} -> Ver saldo de una cuenta")
	fmt.Println("- GET  /pending           -> Ver transacciones pendientes")

	http.ListenAndServe(":8080", servidor.Router)
}

package main

import (
	"encoding/json"
	"fmt"

	"github.com/rcp-code/SimpleBlockchain/core"
)

func main() {
	// Crea una nueva blockchain
	bc := core.CreaBlockchain()

	// Crea algunas cuentas
	bc.CreaCuenta("Alice", 100.0)
	bc.CreaCuenta("Bob", 0.0)

	fmt.Println("Saldos iniciales:")
	fmt.Printf("Alice: %.2f\n", bc.ObtieneBalance("Alice"))
	fmt.Printf("Bob: %.2f\n", bc.ObtieneBalance("Bob"))

	// Realiza una transacción
	err := bc.AgregaTransaccion("Alice", "Bob", 50.0)
	if err != nil {
		fmt.Printf("Error en la transacción: %v\n", err)
		return
	}

	fmt.Println("\nSaldos después de la transacción:")
	fmt.Printf("Alice: %.2f\n", bc.ObtieneBalance("Alice"))
	fmt.Printf("Bob: %.2f\n", bc.ObtieneBalance("Bob"))

	// Muestra la blockchain
	blockchainJson, _ := json.MarshalIndent(bc.Bloques, "", "  ")
	fmt.Printf("\nBlockchain: %s\n", string(blockchainJson))
}

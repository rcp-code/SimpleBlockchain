package core

import "fmt"

//Representa una cuenta
type Cuenta struct {
	Saldo            float64
	DireccionPublica string
}

//Representa una transacción
type Transaccion struct {
	Remitente string
	Receptor  string
	Cantidad  float64
}

// Crea una nueva cuenta con saldo inicial
func (bc *Blockchain) CreaCuenta(direccion string, balanceInicial float64) {
	bc.Balances[direccion] = balanceInicial
}

//Verifica y añade una nueva transacción
func (bc *Blockchain) AgregaTransaccion(origen, destino string, cantidad float64) error {
	//Verifica que el remitente existe y tiene saldo suficiente
	if balance, exists := bc.Balances[origen]; !exists || balance < cantidad {
		return fmt.Errorf("saldo insuficiente en la cuenta o esta no existe")
	}

	//Crea nuevo bloque con la transacción
	transacciones := []Transaccion{{Remitente: origen, Receptor: destino, Cantidad: cantidad}}
	bloquePrevio := bc.Bloques[len(bc.Bloques)-1]
	bloqueNuevo := CreaBloque(bloquePrevio.Indice+1, transacciones, bloquePrevio.Hash)
	bc.Bloques = append(bc.Bloques, bloqueNuevo)

	//Actualiza saldos
	bc.Balances[origen] -= cantidad
	bc.Balances[destino] += cantidad

	return nil
}

//Obtiene el saldo de una cuenta
func (bc *Blockchain) ObtieneBalance(direccion string) float64 {
	return bc.Balances[direccion]
}

package core

import (
	"fmt"
)

// Representa una cuenta
type Cuenta struct {
	saldo            float64
	direccionPublica string
}

// Representa una transacción
type Transaccion struct {
	remitente string  `json:"from"`
	receptor  string  `json:"to"`
	cantidad  float64 `json:"amount"`
}

// Crea una nueva cuenta con saldo inicial
func (bc *Blockchain) CreaCuenta(direccion string, balanceInicial float64) {
	bc.balances[direccion] = balanceInicial
}

// Agrega una transacción a la cola de pendientes
func (bc *Blockchain) AgregaTransaccion(trans Transaccion) error {
	if balance, exists := bc.balances[trans.remitente]; !exists || balance < trans.cantidad {
		return fmt.Errorf("saldo insuficiente o cuenta no existente")
	}
	bc.transaccionesPendientes = append(bc.transaccionesPendientes, trans)
	return nil
}

// Crea un nuevo bloque con las transacciones pendientes
func (bc *Blockchain) TransaccionesPendientesDeMinería(direccionMinado string) (*Bloque, error) {
	// Verificar si hay transacciones pendientes
	if len(bc.transaccionesPendientes) == 0 {
		return nil, fmt.Errorf("no hay transacciones pendientes para minar")
	}
	//Crea un nuevo bloque con las transacciones pendientes
	nuevoBloque := bc.CreaBloque()
	//Añadir recompensa de minería como una transacción
	recompensaTransaccion := Transaccion{
		remitente: "Sistema",
		receptor:  direccionMinado,
		cantidad:  float64(bc.recompensaMinería),
	}
	nuevoBloque.transacciones = append(nuevoBloque.transacciones, recompensaTransaccion)
	//Minería del bloque
	fmt.Printf("Minando bloque %d...\n", nuevoBloque.indice)
	nuevoBloque.Minería(bc.dificultadMinería)
	//Procesado de transacciones
	for _, trans := range nuevoBloque.transacciones {
		if trans.remitente != "Sistema" { // No se procesa la transacción de recompensa aquí
			bc.balances[trans.remitente] -= trans.cantidad
			bc.balances[trans.receptor] += trans.cantidad
		}
	}
	//Añadir recompensa al "minero"
	bc.balances[direccionMinado] += float64(bc.recompensaMinería)
	//Añadir bloque y limpiar transacciones pendientes
	bc.bloques = append(bc.bloques, nuevoBloque)
	bc.transaccionesPendientes = []Transaccion{}
	fmt.Printf("Bloque %d minado exitosamente\n", nuevoBloque.indice)
	fmt.Println("Balances actualizados:")
	return nuevoBloque, nil
}

// Obtiene el saldo de una cuenta
func (bc *Blockchain) ObtieneBalance(direccion string) float64 {
	return bc.balances[direccion]
}

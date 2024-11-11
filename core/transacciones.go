package core

import (
	"fmt"
)

// Representa una cuenta
type Cuenta struct {
	Saldo            float64
	DireccionPublica string
}

// Representa una transacción
type Transaccion struct {
	Remitente string  `json:"from"`
	Receptor  string  `json:"to"`
	Cantidad  float64 `json:"amount"`
}

// Crea una nueva cuenta con saldo inicial
func (bc *Blockchain) CreaCuenta(direccion string, balanceInicial float64) {
	bc.Balances[direccion] = balanceInicial
}

// Agrega una transacción a la cola de pendientes
func (bc *Blockchain) AgregaTransaccion(trans Transaccion) error {
	if balance, exists := bc.Balances[trans.Remitente]; !exists || balance < trans.Cantidad {
		return fmt.Errorf("saldo insuficiente o cuenta no existente")
	}
	bc.TransaccionesPendientes = append(bc.TransaccionesPendientes, trans)
	return nil
}

// Crea un nuevo bloque con las transacciones pendientes
func (bc *Blockchain) TransaccionesPendientesDeMinería(direccionMinado string) (*Bloque, error) {
	// Verificar si hay transacciones pendientes
	if len(bc.TransaccionesPendientes) == 0 {
		return nil, fmt.Errorf("no hay transacciones pendientes para minar")
	}
	//Crea un nuevo bloque con las transacciones pendientes
	nuevoBloque := bc.CreaBloque()
	//Añadir recompensa de minería como una transacción
	recompensaTransaccion := Transaccion{
		Remitente: "Sistema",
		Receptor:  direccionMinado,
		Cantidad:  float64(bc.RecompensaMinería),
	}
	nuevoBloque.Transacciones = append(nuevoBloque.Transacciones, recompensaTransaccion)
	//Minería del bloque
	fmt.Printf("Minando bloque %d...\n", nuevoBloque.Indice)
	nuevoBloque.Minería(bc.DificultadMinería)
	//Procesado de transacciones
	for _, trans := range nuevoBloque.Transacciones {
		if trans.Remitente != "Sistema" { // No se procesa la transacción de recompensa aquí
			bc.Balances[trans.Remitente] -= trans.Cantidad
			bc.Balances[trans.Receptor] += trans.Cantidad
		}
	}
	//Añadir recompensa al "minero"
	bc.Balances[direccionMinado] += float64(bc.RecompensaMinería)
	//Añadir bloque y limpiar transacciones pendientes
	bc.Bloques = append(bc.Bloques, nuevoBloque)
	bc.TransaccionesPendientes = []Transaccion{}
	fmt.Printf("Bloque %d minado exitosamente\n", nuevoBloque.Indice)
	fmt.Println("Balances actualizados:")
	return nuevoBloque, nil
}

// Obtiene el saldo de una cuenta
func (bc *Blockchain) ObtieneBalance(direccion string) float64 {
	return bc.Balances[direccion]
}

package core

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"time"
)

// Representa un bloque en la blockchain
type Bloque struct {
	Indice        int
	MarcaTiempo   string
	Transacciones []Transaccion
	HashPrevio    string
	Hash          string
}

// Blockhain: es un slice de tipo bloque
type Blockchain struct {
	Bloques  []*Bloque
	Balances map[string]float64 // Mapa para mantener los saldos de las cuentas
}

// Calcula el hash a partir del bloque anterior
func (b *Bloque) CalculaHash() string {
	//Se concatenan todos los campos del bloque
	bytesTransacciones, _ := json.Marshal(b.Transacciones)
	registro := string(strconv.Itoa(b.Indice) + b.MarcaTiempo + string(bytesTransacciones) + b.HashPrevio)
	h := sha256.New()
	h.Write([]byte(registro))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

// Crea un nuevo bloque usando el hash del bloque anterior
func CreaBloque(indice int, transacciones []Transaccion, hashPrevio string) *Bloque {
	bloque := &Bloque{
		Indice:        indice,
		MarcaTiempo:   time.Now().String(),
		Transacciones: transacciones,
		HashPrevio:    hashPrevio,
	}
	bloque.Hash = bloque.CalculaHash()
	return bloque
}

// Crea el primer bloque de la cadena
func CreaPrimerBloqueBlockchain() *Bloque {
	return CreaBloque(0, []Transaccion{}, "0")
}

// Crea una nueva blockchain con el mapa de saldos inicializado
func CreaBlockchain() *Blockchain {
	blockchain := &Blockchain{
		Bloques:  []*Bloque{CreaPrimerBloqueBlockchain()},
		Balances: make(map[string]float64),
	}
	return blockchain
}

package core

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Representa un bloque en la blockchain
type Bloque struct {
	indice        int           `json:"index"`
	marcaTiempo   string        `json:"timestamp"`
	transacciones []Transaccion `json:"transactions"`
	hashPrevio    string        `json:"previousHash"`
	hash          string        `json:"hash"`
	numRand       int           `json:"nonce"`
	dificultad    int           `json:"difficulty"`
}

// Blockhain: es un slice de tipo bloque
type Blockchain struct {
	bloques                 []*Bloque
	balances                map[string]float64 // Mapa para mantener los saldos de las cuentas
	transaccionesPendientes []Transaccion
	dificultadMinería       int
	recompensaMinería       int
}

// Calcula el hash a partir del bloque anterior
func (b *Bloque) CalculaHash() string {
	//Se concatenan todos los campos del bloque
	bytesTransacciones, _ := json.Marshal(b.transacciones)
	registro := string(strconv.Itoa(b.indice) + b.marcaTiempo + string(bytesTransacciones) + b.hashPrevio + fmt.Sprint(b.numRand))
	h := sha256.New()
	h.Write([]byte(registro))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

// Crea un nuevo bloque usando el hash del bloque anterior
func (bc *Blockchain) CreaBloque() *Bloque {
	bloquePrevio := bc.bloques[len(bc.bloques)-1]
	bloqueNuevo := &Bloque{
		indice:        len(bc.bloques),
		marcaTiempo:   time.Now().String(),
		transacciones: bc.transaccionesPendientes,
		hashPrevio:    bloquePrevio.hash,
		dificultad:    bc.dificultadMinería,
		numRand:       0,
	}
	return bloqueNuevo
}

// Crea nueva blockchain con una configuración inicial
func CreaBlockchain() *Blockchain {
	genesis := &Bloque{
		indice:        0,
		marcaTiempo:   time.Now().GoString(),
		transacciones: []Transaccion{},
		hashPrevio:    "0",
		dificultad:    4,
	}
	genesis.Minería(4)
	return &Blockchain{
		bloques:           []*Bloque{genesis},
		balances:          make(map[string]float64),
		dificultadMinería: 4,
		recompensaMinería: 10.0,
	}
}

func (b *Bloque) Minería(dif int) {
	prefijo := strings.Repeat("0", dif)
	for {
		b.hash = b.CalculaHash()
		if strings.HasPrefix(b.hash, prefijo) {
			//fmt.Printf("¡Bloque minado de manera exitosa! Hash: %s\n", b.Hash)
			break
		}
		b.numRand++
	}
}

// Método para visualizar el balance de cada cuenta
func VisualizaBalances(bc *Blockchain) {
	for clave, valor := range bc.balances {
		fmt.Printf("Dirección: %s\n -Cantidad: %.2f\n", clave, valor)
	}
}

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
	Indice        int           `json:"index"`
	MarcaTiempo   string        `json:"timestamp"`
	Transacciones []Transaccion `json:"transactions"`
	HashPrevio    string        `json:"previousHash"`
	Hash          string        `json:"hash"`
	NumRand       int           `json:"nonce"`
	Dificultad    int           `json:"difficulty"`
}

// Blockhain: es un slice de tipo bloque
type Blockchain struct {
	Bloques                 []*Bloque
	Balances                map[string]float64 // Mapa para mantener los saldos de las cuentas
	TransaccionesPendientes []Transaccion
	DificultadMinería       int
	RecompensaMinería       int
}

// Calcula el hash a partir del bloque anterior
func (b *Bloque) CalculaHash() string {
	//Se concatenan todos los campos del bloque
	bytesTransacciones, _ := json.Marshal(b.Transacciones)
	registro := string(strconv.Itoa(b.Indice) + b.MarcaTiempo + string(bytesTransacciones) + b.HashPrevio + fmt.Sprint(b.NumRand))
	h := sha256.New()
	h.Write([]byte(registro))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

// Crea un nuevo bloque usando el hash del bloque anterior
func (bc *Blockchain) CreaBloque() *Bloque {
	bloquePrevio := bc.Bloques[len(bc.Bloques)-1]
	bloqueNuevo := &Bloque{
		Indice:        len(bc.Bloques),
		MarcaTiempo:   time.Now().String(),
		Transacciones: bc.TransaccionesPendientes,
		HashPrevio:    bloquePrevio.Hash,
		Dificultad:    bc.DificultadMinería,
		NumRand:       0,
	}
	return bloqueNuevo
}

// Crea nueva blockchain con una configuración inicial
func CreaBlockchain() *Blockchain {
	genesis := &Bloque{
		Indice:        0,
		MarcaTiempo:   time.Now().GoString(),
		Transacciones: []Transaccion{},
		HashPrevio:    "0",
		Dificultad:    4,
	}
	genesis.Minería(4)
	return &Blockchain{
		Bloques:           []*Bloque{genesis},
		Balances:          make(map[string]float64),
		DificultadMinería: 4,
		RecompensaMinería: 10.0,
	}
}

func (b *Bloque) Minería(dif int) {
	prefijo := strings.Repeat("0", dif)
	for {
		b.Hash = b.CalculaHash()
		if strings.HasPrefix(b.Hash, prefijo) {
			//fmt.Printf("¡Bloque minado de manera exitosa! Hash: %s\n", b.Hash)
			break
		}
		b.NumRand++
	}
}

// Método para visualizar el balance de cada cuenta
func VisualizaBalances(bc *Blockchain) {
	for clave, valor := range bc.Balances {
		fmt.Printf("Dirección: %s\n -Cantidad: %.2f\n", clave, valor)
	}
}

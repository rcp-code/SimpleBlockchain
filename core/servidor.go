package core

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// API REST handlers

type APIServidor struct {
	bc     *Blockchain
	Router *mux.Router
}

func NuevoServidor(bc *Blockchain) *APIServidor {
	s := &APIServidor{
		bc:     bc,
		Router: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *APIServidor) routes() {
	s.Router.HandleFunc("/blocks", s.handleGetBloques).Methods("GET")
	s.Router.HandleFunc("/mine", s.handleMineriaBloques).Methods("POST")
	s.Router.HandleFunc("/transaction", s.handleNuevaTransaccion).Methods("POST")
	s.Router.HandleFunc("/balance/{address}", s.handleObtieneBalance).Methods("GET")
	s.Router.HandleFunc("/pending", s.handleObtieneTransaccionesPendientes).Methods("GET")
}

func (s *APIServidor) handleGetBloques(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.bc.bloques)
}

func (s *APIServidor) handleMineriaBloques(w http.ResponseWriter, r *http.Request) {
	var direccionMinado struct {
		Direccion string `json:"address"`
	}
	if err := json.NewDecoder(r.Body).Decode(&direccionMinado); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nuevoBloque, err := s.bc.TransaccionesPendientesDeMinería(direccionMinado.Direccion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(nuevoBloque)
}

func (s *APIServidor) handleNuevaTransaccion(w http.ResponseWriter, r *http.Request) {
	var trans Transaccion
	if err := json.NewDecoder(r.Body).Decode(&trans); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := s.bc.AgregaTransaccion(trans); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *APIServidor) handleObtieneBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	direccion := vars["address"]
	balance := s.bc.balances[direccion]
	json.NewEncoder(w).Encode(map[string]float64{"balance": balance})
}

func (s *APIServidor) handleObtieneTransaccionesPendientes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.bc.transaccionesPendientes)
}

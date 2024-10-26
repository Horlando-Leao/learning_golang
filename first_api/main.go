package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Estrutura de exemplo para os dados da API
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// Variável global para armazenar usuários
var usuarios = []Usuario{
	{ID: 1, Nome: "João"},
	{ID: 2, Nome: "Maria"},
}

// Função para lidar com a rota principal (GET /usuarios)
func getUsuarios(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

// Função para lidar com a criação de um usuário (POST /usuarios)
func createUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Erro ao processar dados", http.StatusBadRequest)
		return
	}

	// Define um novo ID para o usuário (incremental, baseado no último ID)
	if len(usuarios) > 0 {
		usuario.ID = usuarios[len(usuarios)-1].ID + 1
	} else {
		usuario.ID = 1
	}

	// Adiciona o novo usuário à lista global de usuários
	usuarios = append(usuarios, usuario)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuario)
}

// Função para lidar com a remoção de um usuário (DELETE /usuarios/:id)
func deleteUsuario(w http.ResponseWriter, r *http.Request) {

}

func handlerHttp(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		getUsuarios(response, request)
	} else if request.Method == http.MethodPost {
		createUsuario(response, request)
	} else if request.Method == http.MethodDelete {
		deleteUsuario(response, request)
	} else {
		http.Error(response, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Define rotas e suas funções de manipulação
	http.HandleFunc("/usuarios", handlerHttp)

	// Inicia o servidor na porta 8080
	fmt.Println("Servidor iniciado em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

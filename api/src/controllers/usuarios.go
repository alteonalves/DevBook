package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}
	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}
	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	userID, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}
	w.Write([]byte(fmt.Sprintf("ID inserido %d", userID)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
}

package controllers

import "net/http"

func CriarUsuario(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Criando usuario!"))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Buscando todos os usuarios!"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Buscando um usuario!"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Atualizando usuario!"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Deletando usuario!"))
}

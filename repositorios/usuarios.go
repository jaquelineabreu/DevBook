package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := u.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	id, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(id), nil
}

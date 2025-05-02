insert into usuarios (nome, nick, email, senha)
values
("usuário1", "usuario_1", "usuario1@gmail.com", "$2a$12$oNsOkEdD8KW7oR6p3TecYeTHCJ0QG0WvR.dp2a3k5HqJlnFKDeqIu"),
("usuário2", "usuario_2", "usuario2@gmail.com", "$2a$12$oNsOkEdD8KW7oR6p3TecYeTHCJ0QG0WvR.dp2a3k5HqJlnFKDeqIu"),
("usuário3", "usuario_3", "usuario3@gmail.com", "$2a$12$oNsOkEdD8KW7oR6p3TecYeTHCJ0QG0WvR.dp2a3k5HqJlnFKDeqIu");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);
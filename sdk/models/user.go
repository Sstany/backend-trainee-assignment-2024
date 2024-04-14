package models

type User struct {
	ID      int    `json:"id"`
	Login   string `json:"login"`
	Hash    string `json:"hash"`
	IsAdmin bool   `json:"is_admin"`
}

type RegisterForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type LoginForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Access struct {
	Token string `json:"token"`
}

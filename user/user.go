package user

import "strconv"

type User struct {
	Name  string
	Login string
	Email string
	Age   int
}

// Initialise le User via un récepteur pointeur (modifie l'instance appelante)
func New(name, login string) *User {
	return &User{
		Name:  name,
		Login: login,
		Email: login + "@example.com",
		Age:   0, // Valeur par défaut
	}
}

// Met à jour le nom
func (u *User) UpdateName(name string) {
	u.Name = name
}

func(u *User) UpdateEmail(email string) {
	u.Email = email
}

func(u *User) UpdateAge(age int){
	u.Age = age
}

// Retourne une représentation "Name -> Login"
func (u *User) GetName() string {
	return u.Name + " -> " + u.Login + " (" + u.Email + ") " + strconv.Itoa(u.Age)
}
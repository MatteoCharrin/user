package main

import (
	"fmt"
	"github.com/fatih/color"
	"jours2/user"
)

// --- TP1: valeur vs pointeur ---
func zeroval(ival int) {
	ival = 0 // copie : ne modifie pas l'original
}

func zeroptr(iptr *int) {
	*iptr = 0 // pointeur : modifie la valeur à l'adresse
}

func tp1(i int) {
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("after zeroval:", i) // i inchangé

	zeroptr(&i)
	fmt.Println("after zeroptr:", i) // i = 0

	fmt.Printf("address of i: %p\n", &i)
}

// --- TP2: struct, constructeur, alias de pointeur ---
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name} // age implicite à 0
	p.age = 42
	return &p
}

func tp2() {
	var bob, alice person

	// Littéraux de struct
	bob = person{"Bob", 39}
	bob.age = 18

	alice = person{
		name: "Alice",
		age:  30,
	}

	// Constructeur qui renvoie un *person
	p := newPerson("Jon")
	q := p           // q et p pointent sur la même struct
	q.age = 33       // modifie la même instance
	p.age = 44       // encore la même -> écrase 33 par 44

	// Affichages
	fmt.Println("bob   :", bob)
	fmt.Println("alice :", alice)
	fmt.Println("p     :", *p) // déréférencé pour voir la valeur
	fmt.Println("q     :", *q) // idem (même contenu que p)
	fmt.Printf("addr p : %p\n", p)
	fmt.Printf("addr q : %p\n", q)
}


// -- tp 3 -- //



func tp3() {

	u := user.New("Sayo", "sayo123")     // Initialise via le récepteur pointeur
	fmt.Println("User:", u.GetName())

	u.UpdateName("Sayo2")
	u.UpdateAge(25)
	u.UpdateEmail("sayoyo@google.fr")
	fmt.Println("User:", u.GetName())
}


func main() {
	color.Red("TP 1")
	// i := 1
	// tp1(i)

	color.Red("TP 2")
	//tp2()
	color.Red("TP 3")
	tp3()
}

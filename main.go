package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Animal struct
type Animal struct {
	Fly  bool
	Legs int
	Name string
}

// NewAnimal is the Animal struct constructor
func NewAnimal(fly bool, legs int, name string) *Animal {
	return &Animal{
		Fly:  fly,
		Legs: legs,
		Name: name,
	}
}

func (a *Animal) tryFly() string {
	if a.Fly {
		return fmt.Sprintf("It flies")
	}
	return fmt.Sprintf("A %s can't fly!", a.Name)
}

// Page struct
type Page struct {
	Title string
	Body  []byte
}

func main() {
	fmt.Println("Hello World")

	var bird *Animal
	bird = NewAnimal(true, 2, "eagle")
	fmt.Printf("%s has %d legs and %s\n", bird.Name, bird.Legs, bird.tryFly())

	var duck *Animal
	duck = NewAnimal(false, 2, "donald")
	fmt.Printf("a %s Can fly? %s\n", duck.Name, duck.tryFly())

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World\n"))
		fmt.Fprintf(res, "Welcome! accesssing route %s", req.URL.Path)
	})
	http.HandleFunc("/view/", viewHandler)

	log.Println("Listening to :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func viewHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("[%s] %s", req.Method, req.URL.Path)
	title := req.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(res, "<h1>%s</h1><p>%s</p>", p.Title, p.Body)
}

func loadPage(p string) (*Page, error) {
	fname := fmt.Sprintf("res/%s.txt", p)
	body, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return &Page{Title: strings.Title(p), Body: body}, nil
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"com.isatdev.golang.intro/utils"
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

// TryFly function will evaluate fly ability of the Anial
// Fly is field of Animal struct with type boolean
func (a *Animal) TryFly() string {
	if a.Fly {
		return fmt.Sprintf("It flies")
	}
	return fmt.Sprintf("A %s can't fly!", a.Name)
}

// Page struct defines a page with simple title and body template
// we should be able to further process its Body
type Page struct {
	Title string
	Body  []byte
}

func main() {
	fmt.Println("Hello World")

	var bird *Animal
	bird = NewAnimal(true, 2, "eagle")
	fmt.Printf("%s has %d legs and %s\n", bird.Name, bird.Legs, bird.TryFly())

	var duck *Animal
	duck = NewAnimal(false, 2, "donald")
	fmt.Printf("a %s Can fly? %s\n", duck.Name, duck.TryFly())

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World\n"))
		fmt.Fprintf(res, "Welcome! accesssing route %s", req.URL.Path)
	})
	// /view/:title
	http.HandleFunc("/view/", viewHandler)
	// /compare/:left/:right
	http.HandleFunc("/compare/", compareHandler)

	log.Println("Listening to :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func compareHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("[%s] %s", req.Method, req.URL.Path)
	params := req.URL.Path[len("/compare/"):]
	p := strings.Split(params, "/") // parse with simple split
	
	var left, right int
	if len(p) > 0 {
		// convert string of param left into integer
		left, _ = strconv.Atoi(p[0])
		// convert string of param right into integer
		right, _ = strconv.Atoi(p[1])
	}
	// perform simple compare from package utils
	eq := utils.DoSimpleCompare(left, right)
	// load the corresponding page
	v, _ := loadPage("compare")

	sEq := "Equal"
	if !eq {
		sEq = "NOT equal"
	}
	fmt.Fprintf(res, "<h1>%s</h1><p>%s</p><p>%d compare to %d is %s</p>",
		v.Title, v.Body, left, right, sEq)
}

func viewHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("[%s] %s", req.Method, req.URL.Path)
	title := req.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(res, "<h1>%s</h1><p>%s</p>", p.Title, p.Body)
}

func loadPage(p string) (*Page, error) {
	fname := fmt.Sprintf("res/%s.txt", p)
	body, err := ioutil.ReadFile(fname) // read file contents into []byte
	if err != nil {
		return nil, err
	}
	// strings.Title(s) will convert lowercase string into 
	// uppercase first letter of space separted word
	return &Page{Title: strings.Title(p), Body: body}, nil
}

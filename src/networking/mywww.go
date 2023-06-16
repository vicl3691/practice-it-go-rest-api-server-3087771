package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Occupation string `json:"occupation"`
}

var data = `{
	"name":"Vlad",
	"occupation":"Architect"
}`

func main() {
	fmt.Println("in mywww")

	rdr := strings.NewReader(data) //simulate a file/socket
	// Decode Person
	dec := json.NewDecoder(rdr)
	var vlad Person
	if err := dec.Decode(&vlad); err != nil {
		log.Fatalf("error: cannot decode - %s", err)
	}

	fmt.Printf("Person is %+v\n", vlad)

	// Encode message to ...
	txt := "Hi " + vlad.Name + ", we understood you are a(n) " + vlad.Occupation + "."
	msg := map[string] interface {}{
		"ok": true,
		"message": txt,
	}
	enc := json.NewEncoder(os.Stdout) // Put message into stdout
	if err := enc.Encode(msg); err != nil {
		log.Fatalf("Encoding failed - %s", err)
	}

}
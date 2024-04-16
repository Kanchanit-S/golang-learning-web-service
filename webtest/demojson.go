package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	Id           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	data, _ := json.Marshal(&employee{101, "Kanchanit", "093-1324654", "kanchnit@gmail.com"})
	fmt.Println("data = ", string(data))

	e := employee{}
	err := json.Unmarshal([]byte(`{"Id":101, "EmployeeName":"Kanchanit","Tel" :"093-1324654","Email" :"kanchnit@gmail.com"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("e =", e.EmployeeName)
}

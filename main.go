package main

import (
	"fmt"

	"github.com/aashishdube1/lld/solid/ocp"
)

func main() {
	fmt.Println("we rollin")
	data := []ocp.Data{
		{"Aashish", 23, "aashish@gmail.com"},
		{"Rahul", 27, "rahul.sharma@gmail.com"},
		{"Priya", 21, "priya.singh@gmail.com"},
		{"Vikram", 35, "vikram.verma@gmail.com"},
		{"Sneha", 29, "sneha.kumar@gmail.com"},
		{"Aman", 19, "aman.raj@gmail.com"},
		{"Neha", 31, "neha.patel@gmail.com"},
		{"Rohit", 26, "rohit.mehta@gmail.com"},
		{"Ananya", 24, "ananya.jain@gmail.com"},
		{"Karan", 42, "karan.mishra@gmail.com"},
		{"Pooja", 38, "pooja.gupta@gmail.com"},
		{"Arjun", 28, "arjun.singh@gmail.com"},
		{"Simran", 22, "simran.khan@gmail.com"},
		{"Manish", 45, "manish.yadav@gmail.com"},
		{"Riya", 18, "riya.verma@gmail.com"},
	}
	ocp.ExtractContent(data, "csv")
}

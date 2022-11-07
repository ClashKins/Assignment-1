package main

import (
	"fmt"
	"os"
	"reflect"
)

type Bio struct {
	ID        	 int
	Nama      	 string
	Alamat    	 string
	Pekerjaan 	 string
	Alasan   	 string
}
func main() {
	s := Bio{
		ID		: 1,
		Nama		: "Stefanus",
		Alamat		: "Depok",
		Pekerjaan	: "Design Interior",
		Alasan		: "Untuk Mendalami Bahasa Pemrograman Go Lebih Dalam",
	}

	values := reflect.ValueOf(s)
	typesOf := values.Type()
	for i := 0; i < values.NumField(); i++ {
		fmt.Printf("%s \t : %v \n", typesOf.Field(i).Name, values.Field(i).Interface())
	}

	var z, sep string
		for i := 1; i < len(os.Args); i++ {
			z += sep + os.Args[i]
			sep = " "
		}
}

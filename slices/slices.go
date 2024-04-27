package main

import (
	"fmt"
	"reflect"
)

func main() {
	// create array, karena dalam [] ada length
	languages	:= [9]string{
		"C", "Lisp", "C++", "Java", "Python",
		"Javascript", "Ruby", "Go", "Rust",
	}

	// slices
	classics := languages[0:3]
	modern := make([]string, 5)
	modern = languages[3:7]
	new := languages[7:9]

	fmt.Printf("classic prog languages: %v\n", classics)
	fmt.Printf("modern: %v\n", modern)
	fmt.Printf("new: %v\n", new)

	fmt.Println("===================")

	allLangs := languages[:] //copy array jadi slices
	fmt.Println(reflect.TypeOf(allLangs).Kind())

	// create slices, karena dalam [] kosong
	frameworks := []string{
		"React", "Vue", "Angular", "Svelte",
		"Laravel", "Django", "Flask", "Fiber",
	}

	jsFrameworks := frameworks[0:5:5]
	frameworks = append(frameworks, "Meteor") // append gini ga bisa di array, bisanya di slices

	fmt.Printf("all fms: %v\n", frameworks)
	fmt.Printf("js fms: %v\n", jsFrameworks)
}
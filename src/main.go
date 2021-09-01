package main

// Comienza con fmt

import "fmt"

func main() {
	// For conditional

	for i := 0; i <= 10; i++{
		// i empieza con 0 y se ejecuta siempre
		// y cuando sea menor a diez,
		// incrementando 1
		fmt.Println(i)
	}

	fmt.Println("\n") // \n = espaciado

	// For While
	counter := 0
	for counter <= 10 {
		fmt.Println(counter)
		counter++ // Incremental del 0 => 10
	}

	// For forever
	counterForever := 0
	for {
		// Se ejecutara por siempre porque no hay condiciones
		fmt.Println(counterForever)
		counterForever++
	}
}

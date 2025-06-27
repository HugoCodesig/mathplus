package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	mathplus "mathplus/pkg"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: mathplus <función> -p <números>")
		os.Exit(1)
	}

	funcName := os.Args[1]

	// Creamos un FlagSet separado para leer -p
	paramSet := flag.NewFlagSet(funcName, flag.ExitOnError)
	pFlag := paramSet.Bool("p", false, "Indica que lo siguiente son parámetros numéricos")

	// Parseamos desde os.Args[2:] → omite el nombre del ejecutable y la función
	paramSet.Parse(os.Args[2:])

	if !*pFlag {
		fmt.Println("Falta el flag -p con los parámetros numéricos")
		os.Exit(1)
	}

	// Lo que queda en paramSet.Args() son los números
	params := paramSet.Args()

	if len(params) < 2 {
		fmt.Println("Error: necesitas al menos dos números con -p")
		os.Exit(1)
	}

	// Convertimos strings a float64
	var nums []float64
	for _, p := range params {
		n, err := strconv.ParseFloat(p, 64)
		if err != nil {
			fmt.Printf("No pude convertir '%s' a número válido\n", p)
			os.Exit(1)
		}
		nums = append(nums, n)
	}

	// Ejecutamos la función correspondiente
	switch funcName {
	case "add":
		fmt.Println(mathplus.Add(nums...))
	case "mul":
		fmt.Println(mathplus.Mul(nums...))
	default:
		fmt.Printf("Función '%s' no reconocida\n", funcName)
	}
}

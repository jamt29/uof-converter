// Command uof-calculator convierte valores entre unidades de temperatura,
// longitud y peso desde la linea de comandos.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jamt29/uof-calculator/converter"
)

func usage() {
	fmt.Println("Uso: uof-calculator <categoria> <valor> <desde> <hasta>")
	fmt.Println()
	fmt.Println("Categorias y unidades soportadas:")
	fmt.Println("  temp     C (Celsius), F (Fahrenheit), K (Kelvin)")
	fmt.Println("  length   m (metro), km (kilometro), mi (milla), ft (pie)")
	fmt.Println("  weight   kg (kilogramo), g (gramo), lb (libra)")
	fmt.Println()
	fmt.Println("Ejemplo: uof-calculator temp 100 C F")
}

func main() {
	if len(os.Args) != 5 {
		usage()
		os.Exit(1)
	}

	category := os.Args[1]
	value, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "valor invalido: %s\n", os.Args[2])
		os.Exit(1)
	}
	from := os.Args[3]
	to := os.Args[4]

	var result float64
	switch category {
	case "temp":
		result, err = converter.ConvertTemperature(value, converter.TempUnit(from), converter.TempUnit(to))
	case "long":
		result, err = converter.ConvertLength(value, converter.LengthUnit(from), converter.LengthUnit(to))
	case "peso":
		result, err = converter.ConvertWeight(value, converter.WeightUnit(from), converter.WeightUnit(to))
	default:
		fmt.Fprintf(os.Stderr, "categoria desconocida: %s\n", category)
		usage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%.4f %s = %.4f %s\n", value, from, result, to)
}

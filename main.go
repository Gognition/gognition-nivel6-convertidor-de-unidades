package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Converter es la interfaz que todos los convertidores deben implementar
type Converter interface {
	Convert(value float64, from, to string) (float64, error)
}

// LengthConverter maneja conversiones de longitud
type LengthConverter struct {
	factors map[string]float64
}

func NewLengthConverter() *LengthConverter {
	return &LengthConverter{
		factors: map[string]float64{
			"m":  1,
			"cm": 100,
			"mm": 1000,
			"km": 0.001,
			"ft": 3.28084,
			"yd": 1.09361,
			"mi": 0.000621371,
		},
	}
}

func (lc *LengthConverter) Convert(value float64, from, to string) (float64, error) {
	fromFactor, fromOk := lc.factors[from]
	toFactor, toOk := lc.factors[to]

	if !fromOk || !toOk {
		return 0, errors.New("unidad de longitud no soportada")
	}

	meters := value / fromFactor
	return meters * toFactor, nil
}

// WeightConverter maneja conversiones de peso
type WeightConverter struct {
	factors map[string]float64
}

func NewWeightConverter() *WeightConverter {
	return &WeightConverter{
		factors: map[string]float64{
			"kg":  1,
			"g":   1000,
			"lb":  2.20462,
			"oz":  35.274,
			"ton": 0.001,
		},
	}
}

func (wc *WeightConverter) Convert(value float64, from, to string) (float64, error) {
	fromFactor, fromOk := wc.factors[from]
	toFactor, toOk := wc.factors[to]

	if !fromOk || !toOk {
		return 0, errors.New("unidad de peso no soportada")
	}

	kg := value / fromFactor
	return kg * toFactor, nil
}

// TemperatureConverter maneja conversiones de temperatura
type TemperatureConverter struct{}

func (tc *TemperatureConverter) Convert(value float64, from, to string) (float64, error) {
	var celsius float64

	switch from {
	case "C":
		celsius = value
	case "F":
		celsius = (value - 32) * 5 / 9
	case "K":
		celsius = value - 273.15
	default:
		return 0, errors.New("unidad de temperatura de origen no soportada")
	}

	switch to {
	case "C":
		return celsius, nil
	case "F":
		return celsius*9/5 + 32, nil
	case "K":
		return celsius + 273.15, nil
	default:
		return 0, errors.New("unidad de temperatura de destino no soportada")
	}
}

// UnitConverter es el convertidor principal que maneja todos los tipos de conversiones
type UnitConverter struct {
	lengthConverter      Converter
	weightConverter      Converter
	temperatureConverter Converter
}

func NewUnitConverter() *UnitConverter {
	return &UnitConverter{
		lengthConverter:      NewLengthConverter(),
		weightConverter:      NewWeightConverter(),
		temperatureConverter: &TemperatureConverter{},
	}
}

func (uc *UnitConverter) Convert(value float64, from, to, unitType string) (float64, error) {
	switch unitType {
	case "length":
		return uc.lengthConverter.Convert(value, from, to)
	case "weight":
		return uc.weightConverter.Convert(value, from, to)
	case "temperature":
		return uc.temperatureConverter.Convert(value, from, to)
	default:
		return 0, errors.New("tipo de unidad no soportado")
	}
}

func main() {
	converter := NewUnitConverter()

	for {
		fmt.Println("\n===== Convertidor de Unidades =====")
		fmt.Println("1. Longitud")
		fmt.Println("2. Peso")
		fmt.Println("3. Temperatura")
		fmt.Println("4. Salir")
		fmt.Print("Seleccione el tipo de conversión (1-4): ")

		choice := getUserInput()

		if choice == "4" {
			fmt.Println("¡Gracias por usar el Convertidor de Unidades! Hasta luego.")
			break
		}

		if choice < "1" || choice > "3" {
			fmt.Println("Opción no válida. Por favor, intente de nuevo.")
			continue
		}

		unitType := map[string]string{
			"1": "length",
			"2": "weight",
			"3": "temperature",
		}[choice]

		fmt.Print("Ingrese el valor a convertir: ")
		value, err := getFloatInput()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Print("Ingrese la unidad de origen: ")
		from := getUserInput()

		fmt.Print("Ingrese la unidad de destino: ")
		to := getUserInput()

		result, err := converter.Convert(value, from, to, unitType)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Resultado: %.2f %s = %.2f %s\n", value, from, result, to)
		}
	}
}

func getUserInput() string {
	var input string
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}

func getFloatInput() (float64, error) {
	input := getUserInput()
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, errors.New("por favor, ingrese un número válido")
	}
	return value, nil
}

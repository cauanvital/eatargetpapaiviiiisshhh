package main

import "fmt"

func main() {

	faturamentoPorEstado := map[string]float32{
		"SP":     67836.43,
		"RJ":     36678.66,
		"MG":     29229.88,
		"ES":     27165.48,
		"Outros": 19849.53,
	}

	var soma float32
	for _, v := range faturamentoPorEstado {
		soma += v
	}

	for k := range faturamentoPorEstado {
		p := faturamentoPorEstado[k] * 100 / soma
		fmt.Printf("%s foi responsável por %f%% do faturamento\n", k, p)
	}

}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type FaturamentoDiario struct {
	Dia   int     `json:"dia"`
	Valor float64 `json:"valor"`
}

func main() {

	dadosFaturamento, err := LoadData("./atividade 3/dados.json")
	if err != nil {
		log.Fatal(err)
	}

	var faturamentosValidos []FaturamentoDiario
	somaFaturamento := 0.0

	for _, fatura := range dadosFaturamento {

		if fatura.Valor > 0 {
			faturamentosValidos = append(faturamentosValidos, fatura)
			somaFaturamento += fatura.Valor
		}

	}

	if len(faturamentosValidos) == 0 {
		fmt.Println("Nenhum dia com faturamento encontrado (desconsiderando os dias com valor zero).")
		return
	}

	menorFaturamento := faturamentosValidos[0]
	for _, fatura := range faturamentosValidos {
		if fatura.Valor < menorFaturamento.Valor {
			menorFaturamento = fatura
		}
	}
	fmt.Printf("Menor valor de faturamento ocorrido em um dia do mês: R$ %.2f (Dia %d)\n", menorFaturamento.Valor, menorFaturamento.Dia)

	maiorFaturamento := faturamentosValidos[0]
	for _, fatura := range faturamentosValidos {
		if fatura.Valor > maiorFaturamento.Valor {
			maiorFaturamento = fatura
		}
	}
	fmt.Printf("Maior valor de faturamento ocorrido em um dia do mês: R$ %.2f (Dia %d)\n", maiorFaturamento.Valor, maiorFaturamento.Dia)

	mediaMensal := 0.0
	if len(faturamentosValidos) > 0 {
		mediaMensal = somaFaturamento / float64(len(faturamentosValidos))
	}
	fmt.Printf("Média mensal de faturamento (considerando apenas dias com faturamento): R$ %.2f\n", mediaMensal)

	diasAcimaDaMedia := 0
	for _, fatura := range faturamentosValidos {
		if fatura.Valor > mediaMensal {
			diasAcimaDaMedia++
		}
	}
	fmt.Printf("Número de dias no mês em que o valor de faturamento diário foi superior à média mensal: %d\n", diasAcimaDaMedia)

}

func LoadData(filePath string) ([]FaturamentoDiario, error) {

	var result []FaturamentoDiario

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil

}

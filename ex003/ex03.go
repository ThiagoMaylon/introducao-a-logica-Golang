package main

import "fmt"

func escolhaServico() float64 {
	var valorServico float64
	var servico string

	fmt.Println("Entre com o tipo de serviço desejado")
	fmt.Println("DIG - Digitalização")
	fmt.Println("ICO - Impressão Colorida")
	fmt.Println("IPB - Impressão Preto e Branco")
	fmt.Println("FOT - Fotocópia")
	fmt.Print(">> ")
	fmt.Scan(&servico)

	switch servico {
	case "dig":
		valorServico += 1.1
	case "ico":
		valorServico += 1.0
	case "ipb":
		valorServico += 0.4
	case "fot":
		valorServico += 0.2
	}

	return valorServico
}

func numPag() float64 {
	var paginas float64

	fmt.Print("\nEntre com o Numero de Páginas: ")
	fmt.Scan(&paginas)

	if paginas >= 10000 {
		fmt.Println("Não aceitamos tantas paginas de uma vez.")
		fmt.Print("\nEntre com o Numero de Páginas: ")
		fmt.Scan(&paginas)
	}

	if paginas < 10 {
		paginas = paginas - (0 * paginas)
	} else if paginas >= 10 && paginas < 100 {
		paginas = paginas - (0.10 * paginas)
	} else if paginas >= 100 && paginas < 1000 {
		paginas = paginas - (0.15 * paginas)
	} else if paginas >= 1000 && paginas < 10000 {
		paginas = paginas - (0.20 * paginas)
	}

	return paginas
}

func servicoExtra() float64 {
	var servico string
	var valorServicoEx float64
	fmt.Println("\nDeseja adionar mais algum serviço?")
	fmt.Println("1 - Encardernação Simples R$ 10,00")
	fmt.Println("2 - Encardernação capa Dura R$ 25,00")
	fmt.Println("0 - Não desejo mais nada")
	fmt.Print(">> ")
	fmt.Scan(&servico)

	if servico == "0" {
		valorServicoEx += 0
	} else if servico == "1" {
		valorServicoEx += 10
	} else if servico == "2" {
		valorServicoEx = 25
	}

	return valorServicoEx
}

func main() {
	servico := escolhaServico()
	paginas := numPag()
	servicoEx := servicoExtra()
	valorTotal := (servico * paginas) + servicoEx
	fmt.Printf("\nValor (R$) %.2f (serviço: %v * paginas: %v + extra(s): %v)", valorTotal, servico, paginas, servicoEx)

}

package main

import "fmt"

func main() {
	var valorProduto float64
	var quantProduto float64
	var valorTotalDesconto float64

	fmt.Println("Seja Bem-vindo a Loja")

	fmt.Print("Entre com o Valor do produto: ")
	fmt.Scan(&valorProduto)
	fmt.Print("Entre com a quantidade do produto: ")
	fmt.Scan(&quantProduto)

	valorTotal := valorProduto * quantProduto

	fmt.Printf("Valor total do produto SEM desconto: %.2f\n", valorTotal)

	if valorTotal < 1000{
		fmt.Println("Esse valor não tem desconto")
	}else if valorTotal >= 1000 && valorTotal < 3000{
		valorTotalDesconto = valorTotal - (0.03 * valorTotal)
		fmt.Printf("Valor total do produto COM desconto: %.2f\n", valorTotalDesconto)
	}else if valorTotal >= 3000 && valorTotal < 5000{
		valorTotalDesconto = valorTotal - (0.05 * valorTotal)
		fmt.Printf("Valor total do produto COM desconto: %.2f\n", valorTotalDesconto)
	}else if valorTotal >= 5000{
		valorTotalDesconto = valorTotal - (0.08 * valorTotal)
		fmt.Printf("Valor total do produto COM desconto: %.2f\n", valorTotalDesconto)
	}
}
package main

import "fmt"

func menu() {
	fmt.Println("               CARDÁPIO")
	fmt.Println("|" + " Tamanho | Cupuaçu (CP) | Açaí (AC) |")
	fmt.Println("|" + "    P    |   R$ 10,00   |  R$ 12,00 |")
	fmt.Println("|" + "    M    |   R$ 15,00   |  R$ 17,00 |")
	fmt.Println("|" + "    G    |   R$ 19,00   |  R$ 21,00 |")

}

func main() {
	fmt.Println("Seja bem-vindo a loja de Gelados")
	menu()

	count := 0
	var valorTotal float64
	for count < 1{
		var sabor string
		var tamanho string
		
		print("Entre com o sabor desejado (AC/CP): ")
		fmt.Scan(&sabor)
		print("Entre com o tamanho desejado (P/M/G): ")
		fmt.Scan(&tamanho)

		// verifica cp
		if sabor == "cp" && tamanho == "p"{
			valorTotal += 10.0
			fmt.Print("Você pediu CUPUAÇU no tamanho P: R$ 10,00")
		}else if sabor == "cp" && tamanho == "m"{
			valorTotal += 15.0
			fmt.Print("Você pediu CUPUAÇU no tamanho M: R$ 15,00")
		}else if sabor == "cp" && tamanho == "g"{
			valorTotal += 19.0
			fmt.Print("Você pediu CUPUAÇU no tamanho G: R$ 19,00")
		}else if sabor == "ac" && tamanho == "p"{
			valorTotal += 12.0
			fmt.Print("Você pediu AÇAÍ no tamanho P: R$ 12,00")
		}else if sabor == "ac" && tamanho == "m"{
			valorTotal += 17.0
			fmt.Print("Você pediu AÇAÍ no tamanho M: R$ 17,00")
		}else if sabor == "ac" && tamanho == "g"{
			valorTotal += 21.0
			fmt.Print("Você pediu AÇAÍ no tamanho G: R$ 21,00")
		}

		var mais string
		fmt.Print("\n\nDeseja mais alguma coisa? (s/qualquer tecla para sair): ")
		fmt.Scan(&mais)

		if mais != "s"{
			
			count = 1
		}
	}
	fmt.Printf("\nValor total a pagar %.2f", valorTotal)

}

package main

import "fmt"

type Livro struct {
	id      int
	nome    string
	autor   string
	editora string
}

var livros = make([]Livro, 0)
var id int

func cadastraLivro(id int) {
	fmt.Println("CADASTRAR LIVRO")
	var (
		nome    string
		autor   string
		editora string
	)

	fmt.Print("nome do livro: ")
	fmt.Scan(&nome)
	fmt.Print("nome do autor do livro: ")
	fmt.Scan(&autor)
	fmt.Print("nome da Editora: ")
	fmt.Scan(&editora)

	livro := Livro{
		id:      id,
		nome:    nome,
		autor:   autor,
		editora: editora,
	}

	livros = append(livros, livro)
}

func consultaLivro() {
	fmt.Println("CONSULTAR LIVRO")
	var consulta string

	fmt.Println("1 - Consultar todos os Livros")
	fmt.Println("2 - Consultar livro por id'")
	fmt.Println("3 - Consultar livro(s) por autor(a)")
	fmt.Print("\n>> ")
	fmt.Scan(&consulta)

	// irá exibir todos os livros da lista
	if consulta == "1" {
		for _, livro := range livros {
			fmt.Printf("\nid: %v", livro.id)
			fmt.Printf("\nnome: %v", livro.nome)
			fmt.Printf("\nautor: %v", livro.autor)
			fmt.Printf("\neditora: %v", livro.editora)
		}
	} else if consulta == "2" {
		var id int
		fmt.Print("Digite o ID: ")
		fmt.Scan(&id)
		for _, livro := range livros {
			if livro.id == id {
				fmt.Printf("\nid: %v", livro.id)
				fmt.Printf("\nnome: %v", livro.nome)
				fmt.Printf("\nautor: %v", livro.autor)
				fmt.Printf("\neditora: %v", livro.editora)
			}
		}
	} else if consulta == "3" {
		var autor string
		fmt.Print("Digite o nome do(a) autor(a): ")
		fmt.Scan(&autor)
		for _, livro := range livros {
			if livro.autor == autor {
				fmt.Printf("\nid: %v", livro.id)
				fmt.Printf("\nnome: %v", livro.nome)
				fmt.Printf("\nautor: %v", livro.autor)
				fmt.Printf("\neditora: %v", livro.editora)
			}
		}
	}

}
func removeLivro() {
	fmt.Println("REMOVER LIVRO")

	var id int

	fmt.Print("Digite o id do livro a ser removido: ")
	fmt.Scan(&id)

	var newIn int
	for in, livro := range livros {

		if livro.id == id {
			newIn = in
		}
	}
	livros = append(livros[:newIn], livros[newIn+1:]...)

	println("livro removido...\n")
}

func main() {
	fmt.Println("Seja bem-vindo ao controle de livros")

	var opcao string
	count := 0

	for count < 1 {
		fmt.Println("\n1 - Cadastrar Livro")
		fmt.Println("2 - Consultar Livro(s)")
		fmt.Println("3 - Remover Livro")
		fmt.Println("4 - sair")
		fmt.Print(">> ")
		fmt.Scan(&opcao)

		if opcao == "1"{
			cadastraLivro(id)
			id++
		}else if opcao == "2"{
			consultaLivro()
		}else if opcao == "3"{
			removeLivro()
		}else if opcao == "4"{
			count = 1
			break
		}

	}

}

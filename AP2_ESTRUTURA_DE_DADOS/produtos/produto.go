package produtos

import (
	"bufio"
	"os"
	"fmt"
	"log"
)



var TotalProdutosJaCadastrados int

func init() {
	linhas, err := ContaLinhas(nomeArquivo)
	if err != nil {
		log.Fatalf("Erro ao contar linhas do arquivo %s: %v", nomeArquivo, err)
	}
	TotalProdutosJaCadastrados = linhas
}

func ContaLinhas(arquivoCSV string) (int, error) {
	arquivo, err := os.Open(arquivoCSV)
	if err != nil {
		return 0, err
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	contadorLinhas := 0

	for scanner.Scan() {
		contadorLinhas++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return contadorLinhas, nil
}

type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Preco     float64
}

func (p *Produto) definirId() {
	p.Id = TotalProdutosJaCadastrados
}

func (p *Produto) Exibir() {
	fmt.Println("\nProduto", p.Id)
	fmt.Println(p.Nome)
	fmt.Println(p.Descricao)
	fmt.Printf("Preço: R$ %.2f\n", p.Preco)
}

func criar(nome, descricao string, preco float64, id int) Produto {
	p := Produto { Nome: nome, Descricao: descricao, Preco: preco }
	if id == -1 {
		p.definirId()
	} else {
		p.Id = id
	}

	return p
}

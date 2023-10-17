package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Struct do Produto
type Produto struct {
	ID        int
	Nome      string
	Descricao string
	Preco     float64
}

var Produtos []Produto // Variável de produtos exportada
var contIDProduto int

// Função para carregar produtos do arquivo CSV
func CarregarProdutosDoCSV(produtos *[]Produto) error {
	// Abre o arquivo CSV de produtos
	arquivo, err := os.Open("dados/produtos.csv")
	if err != nil {
		return err
	}
	defer arquivo.Close()

	// Lê as linhas do arquivo CSV
	leitor := csv.NewReader(arquivo)
	linhas, err := leitor.ReadAll()
	if err != nil {
		return err
	}

	// Limpa a variável Produtos antes de preenchê-la
	*produtos = []Produto{}

	for _, linha := range linhas[1:] {
		id, _ := strconv.Atoi(linha[0])
		nome := linha[1]
		descricao := linha[2]
		preco, _ := strconv.ParseFloat(linha[3], 64)

		produto := Produto{
			ID:        id,
			Nome:      nome,
			Descricao: descricao,
			Preco:     preco,
		}

		*produtos = append(*produtos, produto)
	}

	// Se não houver produtos, inicialize contIDProduto como 1
	if len(*produtos) == 0 {
		contIDProduto = 1
	} else {
		contIDProduto = (*produtos)[len(*produtos)-1].ID + 1
	}

	return nil
}

// Função para salvar produtos no arquivo CSV
func SalvarProdutosNoCSV(produtos []Produto) error {
	// Abre o arquivo CSV de produtos
	arquivo, err := os.Create("dados/produtos.csv")
	if err != nil {
		return err
	}
	defer arquivo.Close()

	// Cria um escritor CSV
	escritor := csv.NewWriter(arquivo)
	defer escritor.Flush()

	// Escreve o cabeçalho no arquivo
	escritor.Write([]string{"ID", "Nome", "Descrição", "Preço"})

	// Escreve os produtos no arquivo
	for _, produto := range produtos {
		linha := []string{
			strconv.Itoa(produto.ID),
			produto.Nome,
			produto.Descricao,
			strconv.FormatFloat(produto.Preco, 'f', -1, 64),
		}
		escritor.Write(linha)
	}

	return nil
}

// Função para adicionar produtos ao array de produtos
func PostProduto(produtos *[]Produto, nome, descricao string, preco float64) {
	produto := Produto{
		ID:        contIDProduto,
		Nome:      nome,
		Descricao: descricao,
		Preco:     preco,
	}

	*produtos = append(*produtos, produto)
	contIDProduto++
}

// Função para deletar um produto do array de produtos com base no seu ID
func DeleteProduto(produtos *[]Produto, idProduto int) error {
	for i, produto := range *produtos {
		if produto.ID == idProduto {
			*produtos = append((*produtos)[:i], (*produtos)[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Produto não encontrado")
}

// Função para encontrar um produto no array pelo seu ID
func FindProdutoByID(produtos *[]Produto, idProduto int) (*Produto, error) {
	for _, produto := range *produtos {
		if produto.ID == idProduto {
			return &produto, nil
		}
	}
	return nil, fmt.Errorf("Produto não encontrado")
}

// Função que retorna todos os produtos cadastrados
func GetAllProdutos(produtos *[]Produto) []Produto {
	return *produtos
}

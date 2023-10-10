package utils

//Struct do Produto
type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Preco     float64
}

//criação do array de produtos
var produtos []Produto
//contador para a criação de ids.
var contIdProduto int = 1


//Função para adicionar Produtos ao array de produtos.
func postProduto(nome string, descricao string, preco float64) {
	produto := Produto{
		Id:        contIdProduto,
		Nome:      nome,
		Descricao: descricao,
		Preco:     preco,
	}

	produtos = append(produtos, produto)
	contIdProduto++
}

//Função para deletar um produto do array de produtos a partir de seu Id.
func deleteProduto(idProduto int) {
	for i, produto := range produtos {
		if produto.Id == idProduto {
			produtos = append(produtos[:i], produtos[i+1:]...)
		}
	}
}

//Função de busca de um produto dentro do array pelo seu id. Retorna os dados do produto.
func findProdutoById(idProduto int) *Produto {

	for _, produto := range produtos {
		if produto.Id == idProduto {
			return &produto
		}

	}

	return nil
}

//Função que retorna todos os produtos cadastrados.
func getAllProdutos() []Produto {
	return produtos
}

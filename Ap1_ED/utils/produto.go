package utils

type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Preco     float64
}

var produtos []Produto
var contIdProduto int = 1

func PostProduto(nome string, descricao string, preco float64) {
	produto := Produto{
		Id:        contIdProduto,
		Nome:      nome,
		Descricao: descricao,
		Preco:     preco,
	}

	produtos = append(produtos, produto)
	contIdProduto++
}

func DeleteProduto(idProduto int) {
	for i, produto := range produtos {
		if produto.Id == idProduto {
			produtos = append(produtos[:i], produtos[i+1:]...)
		}
	}
}

func FindProdutoById(idProduto int) *Produto {

	for _, produto := range produtos {
		if produto.Id == idProduto {
			return &produto
		}

	}

	return nil
}

func GetAll() []Produto {
	return produtos
}

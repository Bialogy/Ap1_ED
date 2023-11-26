package produtos

import (
	m "mcronalds/metricas"
	"strings"
)

const maxProdutos = 50

type NoProduto struct {
	Produto
	Prox *NoProduto
}

var ListaProdutos *NoProduto
var totalProdutos = 0


func tentarCriar(nome, descricao string, preco float64, id int) *NoProduto {
	if id != -1 {
		_, idProcurado := BuscarId(id)
		if idProcurado != -1 {
			return nil
		}
	}

	return &NoProduto{Produto: criar(nome, descricao, preco, id), Prox: nil}
}

/*
Adiciona um produto com nome, descrição e preço à lista de produtos.
Adiciona o produto primeiro espaço vazio da lista.
Caso já exista um produto com o mesmo id, não adiciona e retorna -3.
Caso já exista um produto com o mesmo nome, não adiciona e retorna erro -2.
Retorna -1 caso a lista esteja cheia, ou o número de produtos cadastrados em caso de sucesso.
*/
func AdicionarUnico(nome, descricao string, preco float64, id int) int {
	if totalProdutos == maxProdutos {
		return -1 // Overflow
	}

	if ListaProdutos == nil {
		ListaProdutos = tentarCriar(nome, descricao, preco, id)
		if ListaProdutos == nil {
			return -3
		}
		totalProdutos++
		m.M.SomaProdutosCadastrados(1)
		return totalProdutos
	}

	noAtual := ListaProdutos
	for noAtual != nil {
		if noAtual.Nome == nome {
			return -2
		}
		noAtual = noAtual.Prox
	}

	noAtual = ListaProdutos
	for noAtual.Prox != nil {
		noAtual = noAtual.Prox
	}

	noAtual.Prox = tentarCriar(nome, descricao, preco, id)
	if noAtual.Prox == nil {
		return -3
	}

	totalProdutos++
	m.M.SomaProdutosCadastrados(1)
	return totalProdutos
}

/*
Localiza um produto a partir do seu id.
Retorna o produto encontrado e a sua posição na lista, em caso de sucesso.
Retorna um produto vazio e -1 em caso de erro.
*/
func BuscarId(id int) (*NoProduto, int) {
	noAtual := ListaProdutos
	indice := 0

	for noAtual != nil {
		if noAtual.Id == id {
			return noAtual, indice
		}
		noAtual = noAtual.Prox
		indice++
	}

	return nil, -1
}

/*
Localiza produtos que iniciem com a string passada.
Retorna um slice com todos os produtos encontrados, e o tamanho do slice.
*/
func BuscarNome(comecaCom string) ([]Produto, int) {
	var produtosEncontrados []Produto
	noAtual := ListaProdutos

	for noAtual != nil {
		if strings.HasPrefix(noAtual.Nome, comecaCom) {
			produtosEncontrados = append(produtosEncontrados, noAtual.Produto)
		}
		noAtual = noAtual.Prox
	}

	return produtosEncontrados, len(produtosEncontrados)
}

/*
Exibe todos os produtos cadastrados.
*/
func Exibir() {
	noAtual := ListaProdutos

	for noAtual != nil {
		noAtual.Exibir()
		noAtual = noAtual.Prox
	}
}

func ExibirPorNome() {

	primeiroNo := ListaProdutos
	var ListaProdutosAnt *NoProduto

	for i := 0; i < totalProdutos-1; i++ {
		ListaProdutos = primeiroNo
		ListaProdutosAnt = nil

		for j := 0; j < totalProdutos-i-1; j++ {
			if ListaProdutos.Prox != nil && strings.Compare(ListaProdutos.Nome, ListaProdutos.Prox.Nome) > 0 {
				temp := ListaProdutos.Prox
				ListaProdutos.Prox = ListaProdutos.Prox.Prox
				temp.Prox = ListaProdutos
				if j == 0 {
					primeiroNo = temp
				} else {
					ListaProdutosAnt.Prox = temp
				}
			} else {
				ListaProdutosAnt = ListaProdutos
				ListaProdutos = ListaProdutos.Prox
			}
		}
	}

	ListaProdutos = primeiroNo

	noAtual := ListaProdutos
	for noAtual != nil {
		noAtual.Exibir()
		noAtual = noAtual.Prox
	}
}

/*
Remove um produto da lista a partir do seu id.
Retorna -2 caso não haja produtos na lista.
Retorna -1 caso não haja um produto com o id passado, ou 0 em caso de sucesso.
*/
func Excluir(id int) int {
	if ListaProdutos == nil {
		return -2
	}

	if ListaProdutos.Id == id {
		ListaProdutos = ListaProdutos.Prox
		totalProdutos--
		m.M.SomaProdutosCadastrados(-1)
		return 0
	}

	noAnt := ListaProdutos
	noAtual := ListaProdutos.Prox

	for noAtual != nil {
		if noAtual.Id == id {
			noAnt.Prox = noAtual.Prox
			totalProdutos--
			m.M.SomaProdutosCadastrados(-1)
			return 0
		}
		noAnt = noAtual
		noAtual = noAtual.Prox
	}

	return -1
}

/*
Atualiza um produto na lista a partir do seu id.
Retorna -2 caso não haja produtos na lista.
Retorna -1 caso não haja um produto com o id passado, ou 0 em caso de sucesso.
*/
func Atualizar(id int, novoPreco float64) int {
	noAtual := ListaProdutos

	for noAtual != nil {
		if noAtual.Id == id {
			noAtual.Preco = novoPreco
			return 0
		}
		noAtual = noAtual.Prox
	}

	return -1
}
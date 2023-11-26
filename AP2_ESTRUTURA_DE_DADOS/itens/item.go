package itens

import (
	p "mcronalds/produtos"
)

type Item struct {
	Produto *p.NoProduto
	Quant   int
}

func (q *Item) CalcularPrecoParcial() float64 {
	if q.Produto == nil || *q.Produto == (p.NoProduto{}) {
		return 0.0
	}
	return q.Produto.Produto.Preco * float64(q.Quant)
}

/*
Retorna um Item com as informações solicitadas.
Se o id não existir para um produto, retorna um Item vazio.
*/
func Criar(id int, quant int) Item {
	produto, _ := p.BuscarId(id)
	if produto == nil || *produto == (p.NoProduto{}) {
		return Item{}
	}

	return Item{Produto: produto, Quant: quant}
}

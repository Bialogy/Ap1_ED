package utils

import (
	"time"
	"fmt"
)

type Pedido struct {
	Id               int
	Delivery         bool
	idProduto        int
	PrecoTotalPedido float64
	HoraPedido       time.Time
	Concluido        bool
}

var pedidos []Pedido
var contIdPedido int = 1

func CalculaValorTotalPedido(produtos []Produto) float64 {
	var precoTotalPedido float 64
	for _, produto := range produtos {
		precoTotalPedido += produto.Preco
	}
	return precoTotalPedido
}

func PostPedido(delivery bool, idProduto int) error {

	if len(idProduto) < 1 || len (idProduto) > 10 {
		return fmt.Println "Por favor insira de 1 até 10 produtos."
	}

	pedido := Pedido {
	Id:          contIdPedido,
	Delivery:    delivery,
	idProduto:   idProduto,
	HoraPedido:  time.Now
	}

	order.PrecoTotalPedido = CalculaValorTotalPedido(produtos []Produto)

	pedidos = append(pedidos, pedido)

	contIdPedido++

}

func DeletePedido(idPedido int) {

}

func ConcluirPedido() {
	for i, pedido := range pedidos {
		if !pedido.Concluido {
			pedidos[i].Concluido = true
			fmt.Printf("Pedido concluído.")
			return
		}
	}
	fmt.Println("Não há pedido a ser concluído.")
}

func FindPedidoById(idPedido int) *Pedido {

}

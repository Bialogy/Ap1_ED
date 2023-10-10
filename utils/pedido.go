package utils

import (
	"time"
	"fmt"
)

//Struct do Pedido
type Pedido struct {
	Id               int
	Delivery         bool
	idProduto        int
	PrecoTotalPedido float64
	HoraPedido       time.Time
	Encerrado        bool
}

//criação do array de pedidos
var pedidos []Pedido
//contador para a criação dos ids.
var contIdPedido int = 1

//Função para somar o preco de todos os produtos de um pedido
func calculaValorTotalPedido(produtos []Produto) float64 {
	var precoTotalPedido float 64
	for _, produto := range produtos {
		precoTotalPedido += produto.Preco
	}
	return precoTotalPedido
}

//Função para adicionar um pedido
func postPedido(delivery bool, idProduto int) error {

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

//Função para deletar um pedido.
func deletePedido(idPedido int) {

}

//Função para marcar um pedido como concluído ou não
func encerrarPedido() {
	for i, pedido := range pedidos {
		if !pedido.Concluido {
			pedidos[i].Concluido = true
			fmt.Printf("Pedido concluído.")
			return
		}
	}
	fmt.Println("Este pedido não existe.")
}

//Função para encontrar um pedido pelo seu id
func findPedidoById(idPedido int) *Pedido {

}

//Função que retorna todos os pedidos.
func getAllPedidos() []Pedido {
	return pedidos
}

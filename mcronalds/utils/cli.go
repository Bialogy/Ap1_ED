package utils

import (
	"fmt"
	"time"
)

var pedidosEncerrados int
var faturamentoTotal float64
var tempoTotalExpedicao time.Duration

// Exibe métricas do sistema
func ExibeMetricas() {
	produtosCadastrados := len(*produtos)
	pedidosEncerrados, pedidosEmAndamento := CalculaStatusPedido()
	faturamentoTotal := CalculaFaturamento()
	tempoTotalExpedicao := CalculaTotalExpedição()

	fmt.Printf("Número de produtos cadastrados: %d\n", produtosCadastrados)
	fmt.Printf("Número total de pedidos encerrados: %d\n", pedidosEncerrados)
	fmt.Printf("Número de pedidos em andamento: %d\n", pedidosEmAndamento)
	fmt.Printf("Faturamento total até o momento: R$%.2f\n", faturamentoTotal)

	// Calcula o tempo médio de expedição
	tempoMedioExpedicao := int(tempoTotalExpedicao.Seconds() / float64(pedidosEncerrados) / 60)
	fmt.Printf("Tempo médio de expedição (minutos): %d\n", tempoMedioExpedicao)
}

func CalculaTotalExpedição() time.Duration {
	pedidos := GetAllPedidos()
	for _, pedido := range pedidos {

		if pedido.Encerrado {
			tempoTotalExpedicao += pedido.TempoExpedicao
		}
	}
	return tempoTotalExpedicao
}

// Calcula o faturamento
func CalculaFaturamento() float64 {
	pedidos := GetAllPedidos()
	faturamentoTotal := 0.0

	for _, pedido := range pedidos {
		faturamentoTotal += pedido.PrecoTotalPedido
	}

	return faturamentoTotal

}

// Calcula o número de pedidos encerrados e em andamento
func CalculaStatusPedido() (int, int) {
	pedidos := GetAllPedidos()
	pedidosEncerrados := 0
	pedidosEmAndamento := 0

	for _, pedido := range pedidos {
		if pedido.Encerrado {
			pedidosEncerrados++
		} else {
			pedidosEmAndamento++
		}
	}

	return pedidosEncerrados, pedidosEmAndamento
}

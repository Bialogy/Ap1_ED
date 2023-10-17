package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

// Struct do Pedido
type Pedido struct {
	ID               int
	Delivery         bool
	IDProdutos       []int
	PrecoTotalPedido float64
	HoraPedido       time.Time
	Encerrado        bool
	TempoExpedicao   time.Duration
}

// Criação do ponteiro para o array de produtos
var produtos *[]Produto

// Criação do array de pedidos
var pedidos []Pedido
var contIDPedido int = 1

// Função para configurar o ponteiro para o array de produtos
func SetProdutos(p *[]Produto) {
	produtos = p
}

// Função para calcular o valor total do pedido com base em IDs de produtos
func CalculaValorTotalPedido(produtosQuantidade map[int]int) float64 {
	var precoTotalPedido float64
	for idProduto, quantidade := range produtosQuantidade {
		produto, err := FindProdutoByID(produtos, idProduto)
		if err != nil {
			fmt.Println("Erro ao calcular o valor do pedido:", err)
			return 0
		}
		precoTotalPedido += produto.Preco * float64(quantidade)
	}
	return precoTotalPedido
}

// Função para adicionar um pedido
func PostPedido(delivery bool, produtosQuantidade map[int]int) (int, error) {
	pedidosEmAndamento := 0
	for _, pedido := range pedidos {
		if !pedido.Encerrado {
			pedidosEmAndamento++
		}
	}

	// Verifica o limite de 1000 pedidos em andamento
	if pedidosEmAndamento >= 1000 {
		return -1, fmt.Errorf("O limite de 1000 pedidos em andamento foi atingido. Não é possível adicionar mais pedidos.")
	}

	if len(produtosQuantidade) > 10 {
		return -1, fmt.Errorf("O total de produtos no pedido excede o limite de 10.")
	}

	pedido := Pedido{
		ID:         contIDPedido,
		Delivery:   delivery,
		HoraPedido: time.Now(),
		Encerrado:  false,
	}

	// Cálculo do preço total do pedido com base nas quantidades de produtos
	pedido.PrecoTotalPedido = CalculaValorTotalPedido(produtosQuantidade)

	// Adicionar taxa de entrega se for um pedido de delivery
	if delivery {
		pedido.PrecoTotalPedido += 10.0
	}

	pedidos = append(pedidos, pedido)
	contIDPedido++

	return pedido.ID, nil
}

// Função para deletar um pedido do array de pedidos com base no seu ID
func DeletePedido(idPedido int) error {
	for i, pedido := range pedidos {
		if pedido.ID == idPedido {
			if pedido.Encerrado {
				return fmt.Errorf("Não é possível excluir um pedido encerrado.")
			}
			pedidos = append(pedidos[:i], pedidos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Pedido não encontrado")
}

// Função para encontrar um pedido no array pelo seu ID
func FindPedidoByID(idPedido int) (*Pedido, error) {
	for _, pedido := range pedidos {
		if pedido.ID == idPedido {
			return &pedido, nil
		}
	}
	return nil, fmt.Errorf("Pedido não encontrado")
}

// Função que retorna todos os pedidos cadastrados
func GetAllPedidos() []Pedido {
	return pedidos
}

// Função para contar quantos pedidos estão na frente do ID especificado
func PedidosNaFrente(idPedido int) int {
	count := 0
	for _, pedido := range pedidos {
		if pedido.ID == idPedido {
			break
		}
		if !pedido.Encerrado {
			count++
		}
	}
	return count
}

// Função para expedir um pedido e adicionar ao faturamento
func ExpedirPedido() (*Pedido, error) {
	for i, pedido := range pedidos {
		if !pedido.Encerrado {
			// Calcula o tempo de expedição
			tempoExpedicao := time.Since(pedido.HoraPedido)
			pedidos[i].TempoExpedicao = tempoExpedicao

			// Marca o pedido como encerrado
			pedidos[i].Encerrado = true

			return &pedido, nil
		}
	}

	return nil, fmt.Errorf("Nenhum pedido em aberto para ser expedido.")
}

// Função para salvar todos os pedidos em um arquivo CSV
func SalvarPedidosNoCSV() error {
	var arquivoCSV *os.File
	var err error

	_, err = os.Stat("dados/pedidos.csv")

	if err == nil {
		arquivoCSV, err = os.OpenFile("dados/pedidos.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
	} else {
		arquivoCSV, err = os.Create("dados/pedidos.csv")
		if err != nil {
			return err
		}

		cabecalho := []string{"ID", "Delivery", "PrecoTotalPedido", "HoraPedido"}
		escritorCSV := csv.NewWriter(arquivoCSV)
		escritorCSV.Write(cabecalho)
		escritorCSV.Flush()
	}

	escritorCSV := csv.NewWriter(arquivoCSV)
	escritorCSV.Comma = ';'

	// Escrever apenas o último pedido
	ultimoPedido := pedidos[len(pedidos)-1]
	strID := fmt.Sprintf("%d", ultimoPedido.ID)
	strDelivery := fmt.Sprintf("%t", ultimoPedido.Delivery)
	strPrecoTotalPedido := fmt.Sprintf("%.2f", ultimoPedido.PrecoTotalPedido)
	strHoraPedido := ultimoPedido.HoraPedido.Format("2006-01-02 15:04:05")
	linha := []string{strID, strDelivery, strPrecoTotalPedido, strHoraPedido}
	escritorCSV.Write(linha)
	escritorCSV.Flush()

	return nil
}

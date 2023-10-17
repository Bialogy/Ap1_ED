package main

import (
	"fmt"
	"mcronalds/utils"
	"os"
	"strings"
	"bufio"
)

func main() {
	var produtos []utils.Produto

	if err := utils.CarregarProdutosDoCSV(&produtos); err != nil {
		fmt.Println("Erro ao carregar produtos:", err)
		return
	}

	// Define o ponteiro de produtos para que o menu do cliente possa acessá-los
	utils.SetProdutos(&produtos)

	clienteMenu(&produtos)
}

func clienteMenu(produtos *[]utils.Produto) {
	fmt.Println("Bem-vindo ao McRonald's")
	for {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1. Cliente")
		fmt.Println("2. Funcionário")
		fmt.Println("3. Encerrar programa")
		fmt.Print("Opção: ")

		var escolha string
		fmt.Scanln(&escolha)

		switch escolha {
		case "1":
			menuCliente(produtos)
		case "2":
			menuFuncionario(produtos)
		case "3":
			// Salva os produtos no arquivo CSV antes de encerrar o programa
			utils.SalvarProdutosNoCSV(*produtos)
			fmt.Println("Programa encerrado.")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}

func menuCliente(produtos *[]utils.Produto) {
	fmt.Println("Bem-vindo ao McRonald's - clientes")
	for {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1. Ver produtos")
		fmt.Println("2. Cadastrar pedido")
		fmt.Println("3. Deletar pedido")
		fmt.Println("4. Achar pedido")
		fmt.Println("5. Verificar posição na fila")
		fmt.Println("6. Voltar para o menu anterior")
		fmt.Print("Opção: ")

		var opcao string
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			// Opção: Ver produtos
			fmt.Println("Produtos disponíveis:")
			for _, produto := range *produtos {
				fmt.Printf("ID: %d, Nome: %s, Descrição: %s, Preço: R$%.2f\n", produto.ID, produto.Nome, produto.Descricao, produto.Preco)
			}

		case "2":
			// Opção: Cadastrar pedido
			var deliveryInput string
			for {
				fmt.Print("Pedido para delivery? (S/N): ")
				fmt.Scanln(&deliveryInput)
				if strings.EqualFold(deliveryInput, "s") || strings.EqualFold(deliveryInput, "n") {
					break
				}
			}

			delivery := strings.EqualFold(deliveryInput, "s")

			fmt.Print("Quantos produtos diferentes deseja pedir? ")
			var numProdutos int
			fmt.Scanln(&numProdutos)

			if numProdutos > 10 {
				fmt.Println("O limite de 10 produtos diferentes foi excedido.")
				return
			}

			produtosQuantidade := make(map[int]int)

			for i := 0; i < numProdutos; i++ {
				fmt.Printf("Informe o ID do produto %d: ", i+1)
				var produtoID int
				fmt.Scanln(&produtoID)

				fmt.Printf("Informe a quantidade do produto %d: ", i+1)
				var quantidade int
				fmt.Scanln(&quantidade)

				produtosQuantidade[produtoID] = quantidade
			}

			idDoPedido, err := utils.PostPedido(delivery, produtosQuantidade)
			if err != nil {
				fmt.Println("Erro ao cadastrar pedido:", err)
			} else {
				fmt.Printf("Pedido cadastrado com sucesso! ID do Pedido: %d\n", idDoPedido)

				// Salva os pedidos no arquivo CSV toda vez que um novo pedido é criado
				if err := utils.SalvarPedidosNoCSV(); err != nil {
					fmt.Println("Erro ao salvar os pedidos no arquivo CSV:", err)
				}
			}

		case "3":
			// Opção: Deletar pedido
			fmt.Print("Informe o ID do pedido que deseja deletar: ")
			var pedidoID int
			fmt.Scanln(&pedidoID)
			err := utils.DeletePedido(pedidoID)
			if err != nil {
				fmt.Println("Erro ao deletar pedido:", err)
			} else {
				fmt.Println("Pedido deletado com sucesso!")
			}

		case "4":
			// Opção: Achar pedido
			fmt.Print("Informe o ID do pedido que deseja encontrar: ")
			var pedidoID int
			fmt.Scanln(&pedidoID)
			pedido, err := utils.FindPedidoByID(pedidoID)

			if pedido.Encerrado {
				fmt.Println("Pedido finalizado!")
				return
			}

			if err != nil {
				fmt.Println("Pedido não encontrado:", err)
			} else {
				fmt.Printf("Pedido encontrado:\nID: %d\nDelivery: %v\nValor Total: R$%.2f\n", pedido.ID, pedido.Delivery, pedido.PrecoTotalPedido)
			}

		case "5":
			// Opção: Verificar posição na fila
			fmt.Print("Informe o ID do seu pedido: ")
			var pedidoID int
			fmt.Scanln(&pedidoID)

			pedido, err := utils.FindPedidoByID(pedidoID)

			if pedido.Encerrado {
				fmt.Println("Pedido finalizado!")
				return
			} else if !pedido.Encerrado {
				fmt.Println("Pedido em andamento.")
			} else {
				fmt.Println("Pedido não encontrado:", err)
			}

			if len(utils.GetAllPedidos()) == 0 {
				fmt.Println("Não há pedidos cadastrados.")
			} else {
				pedidosNaFrente := utils.PedidosNaFrente(pedidoID)
				if pedidosNaFrente < 0 {
					fmt.Println("Erro ao verificar posição na fila.")
				} else if pedidosNaFrente == 0 {
					fmt.Println("Seu pedido é o próximo a ser expedido!")
				} else {
					fmt.Printf("Há %d pedidos na frente do seu pedido.\n", pedidosNaFrente)
				}
			}

		case "6":
			// Opção: Voltar para o menu anterior
			return

		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}

func menuFuncionario(produtos *[]utils.Produto) {
	fmt.Println("Bem-vindo ao McRonald's - funcionários")
	for {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1. Procurar produto")
		fmt.Println("2. Procurar pedido")
		fmt.Println("3. Listar produtos")
		fmt.Println("4. Listar pedidos")
		fmt.Println("5. Exibir métricas")
		fmt.Println("6. Voltar para o menu anterior")
		fmt.Print("Opção: ")

		var opcaoFuncionario string
		fmt.Scanln(&opcaoFuncionario)

		switch opcaoFuncionario {
		case "1":
			// Opção: Procurar produto
			fmt.Print("Informe o ID do produto que deseja procurar: ")
			var produtoID int
			fmt.Scanln(&produtoID)

			produto, err := utils.FindProdutoByID(produtos, produtoID)
			if err != nil {
				fmt.Println("Produto não encontrado:", err)
			} else {
				fmt.Printf("Produto encontrado - ID: %d, Nome: %s, Descrição: %s, Preço: R$%.2f\n", produto.ID, produto.Nome, produto.Descricao, produto.Preco)
			}

		case "2":
			// Opção: Procurar pedido
			fmt.Print("Informe o ID do pedido que deseja procurar: ")
			var pedidoID int
			fmt.Scanln(&pedidoID)

			pedido, err := utils.FindPedidoByID(pedidoID)
			if err != nil {
				fmt.Println("Pedido não encontrado:", err)
			} else {
				fmt.Printf("Pedido encontrado - ID: %d, Delivery: %v, Valor Total: R$%.2f\n", pedido.ID, pedido.Delivery, pedido.PrecoTotalPedido)
			}

		case "3":
			// Opção: Listar produtos
			fmt.Println("Produtos disponíveis:")
			for _, produto := range *produtos {
				fmt.Printf("ID: %d, Nome: %s, Descrição: %s, Preço: R$%.2f\n", produto.ID, produto.Nome, produto.Descricao, produto.Preco)
			}

			// Submenu de gerenciamento de produtos
			fmt.Println("Escolha uma opção:")
			fmt.Println("1. Cadastrar novo produto")
			fmt.Println("2. Deletar produto")
			fmt.Println("3. Voltar para o menu anterior")
			fmt.Print("Opção: ")

			var opcaoProduto string
			fmt.Scanln(&opcaoProduto)

			switch opcaoProduto {
			case "1":
				// Opção: Cadastrar novo produto
				reader := bufio.NewReader(os.Stdin)

                fmt.Print("Nome do novo produto: ")
                novoNome, _  := reader.ReadString('\n')
                novoNome = strings.TrimSpace(novoNome)

                fmt.Print("Descrição do novo produto: ")
                novaDescricao, _ := reader.ReadString('\n')
                novaDescricao = strings.TrimSpace(novaDescricao)

                fmt.Print("Preço do novo produto (R$): ")
                var novoPreco float64
                fmt.Scanln(&novoPreco)
				

				utils.PostProduto(produtos, novoNome, novaDescricao, novoPreco)
				fmt.Println("Novo produto cadastrado com sucesso!")

			case "2":
				// Opção: Deletar produto
				fmt.Print("Informe o ID do produto que deseja deletar: ")
				var produtoID int
				fmt.Scanln(&produtoID)

				err := utils.DeleteProduto(produtos, produtoID)
				if err != nil {
					fmt.Println("Erro ao deletar o produto:", err)
				} else {
					fmt.Println("Produto deletado com sucesso!")
				}

			case "3":
				// Opção: Voltar para o menu anterior
				return

			default:
				fmt.Println("Opção inválida. Tente novamente.")
			}
		case "4":
			// Opção: Listar pedidos
			pedidos := utils.GetAllPedidos()
			if len(pedidos) == 0 {
				fmt.Println("Não há pedidos em andamento.")
			} else {
				fmt.Println("Pedidos em andamento:")
				for _, pedido := range pedidos {
					if !pedido.Encerrado {
						fmt.Printf("ID: %d, Delivery: %v, Valor Total: R$%.2f\n", pedido.ID, pedido.Delivery, pedido.PrecoTotalPedido)
					}
				}

				// Submenu de gerenciamento de pedidos
				fmt.Println("Escolha uma opção:")
				fmt.Println("1. Expedir pedido")
				fmt.Println("2. Voltar para o menu anterior")
				fmt.Print("Opção: ")

				var opcaoPedido string
				fmt.Scanln(&opcaoPedido)

				switch opcaoPedido {
				case "1":
					// Opção: Expedir pedido
					pedido, err := utils.ExpedirPedido()
					if err != nil {
						fmt.Println("Erro ao expedir pedido:", err)
					} else {
						fmt.Printf("Pedido ID %d expedido com sucesso!\n", pedido.ID)
					}

				case "2":
					// Opção: Voltar para o menu anterior
					return

				default:
					fmt.Println("Opção inválida. Tente novamente.")
				}
			}
		case "5":
			// Opção: Exibir as métricas do sistema
			utils.ExibeMetricas()

		case "6":
			// Opção: Voltar para o menu anterior
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")

		}
	}
}

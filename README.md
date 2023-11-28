# Ap1 Estrutura de Dados 2023.2

## **Grupo**:
+ **Victor Furtado**
+ **Ana Beatriz Lessa**
+ **João Pedro Alencar**

## Primeiros Passos

**Opções do Menu:**
    - Ao executar o programa, você será apresentado a um menu de opções. Digite o número correspondente para realizar a ação desejada.

## Opções do Menu

1. **Cadastrar produto:**
    - Adiciona um novo produto ao sistema. Siga as instruções para inserir os detalhes do produto.

2. **Remover produto:**
    - Remove um produto do sistema. Informe o ID do produto a ser removido.

3. **Atualizar produto:**
    - Atualiza os detalhes de um produto. Informe o ID do produto a ser atualizado e forneça as novas informações.

4. **Buscar produto por ID:**
    - Busca um produto pelo seu ID. Informe o ID do produto a ser pesquisado.

5. **Buscar produto por nome:**
    - Busca produtos pelo nome ou pelo início do nome. Informe o nome ou parte do nome do produto.

6. **Exibir todos os produtos ordenados por ID:**
    - Exibe todos os produtos ordenados pelos seus IDs.

7. **Exibir todos os produtos ordenados por nome:**
    - Exibe todos os produtos ordenados pelos seus nomes.

8. **Adicionar pedido:**
    - Adiciona um novo pedido. Você pode especificar se o pedido é para entrega. Informe os IDs dos produtos e as quantidades para o pedido.

9. **Expedir pedido:**
    - Marca um pedido como expedido. Informe o ID do pedido a ser expedido.

10. **Exibir métricas do sistema:**
    - Exibe métricas do sistema, incluindo o número total de produtos, pedidos, etc.

20. **Exibir todos os pedidos em andamento:**
    - Exibe todos os pedidos que estão atualmente em andamento.

21. **Cadastrar produtos em lote:**
    - Permite adicionar vários produtos em lote. Siga as instruções para adicionar produtos até decidir parar.

100. **Sair do programa:**
    - Encerra o programa. Esta opção permite encerrar o programa de forma controlada.

## Exemplos

### Cadastrar produto:
```plaintext
1
Nome do produto: batata
Descrição: descrição do batata
Preço do produto (em R$): 5.00
```

### Buscar produto por nome:
```plaintext
5
Informe o nome do produto ou o início do nome: batata
```
Saída:
```plaintext
Produto encontrado:
ID: 1
Nome: batata
Descrição: descrição do batata
Preço: R$ 5.00
```

### Adicionar pedido:
```plaintext
8
O pedido é para entrega (s/n)? n
1 2
2 1
0 0
```

### Expedir pedido:
```plaintext
9
Informe o ID do pedido a ser expedido: 1
```

### Exibir métricas do sistema:
```plaintext
10
```
Saída:
```plaintext
Métricas do Sistema:
Total de produtos: 2
Total de pedidos: 1
```

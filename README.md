# Ap1 Estrutura de Dados 2023.2

## **Grupo**:
+ **Gustavo Braga**
+ **Ana Beatriz Lessa**
+ **João Pedro Alencar**

### Cli McRonalds

O arquivo CliMcRonalds.exe se encontra dentro da pasta mcronalds.

+ O progama inicia com uma interface podendo entrar como cliente, funcionário ou encerrar programa.
+ Opção 1, **cliente**: ver produtos, cadastrar pedido,  deletar pedido, achar pedido, verificar posição na fila e voltar para o menu anterior. 
+ Opção 2, **funcionário**: procurar produto, procurar pedido, listar produtos, listar pedidos, voltar para o menu anterior. 
+ Opção 3, Encerrar o programa.

**Cliente**:
+ Opção 1: todos os produtos disponíveis para pedir serão listados.
    + cada produto tem id, nome, descrição e preço.
+ Opção 2: 
    + o cliente deve indicar se o produto é ou não um delivery através de um "s" ou "n". Caso escolha que sim, um valor de 10.00 será adicionado ao valor total do pedido.
    + o cliente podera escolher quantos produtos diferentes ele vai pedir, o produto que ele vai pedir pelo id e a quantidade de cada produto.
+ Opção 3: o cliente poderá deletar um pedido que já realizou.
+ Opção 4: pede para o cliente escolher o id do pedido e depois mostra o pedido. 
+ Opção 5: pede para o cliente escolher o id do pedido que deseja verificar a posição na fila.
    + Mostra se o pedido está expedido, se é o próximo da fila ou sua posição na fila.
+ Opção 6: Retorna para o menu inicial.

**Funcionário**:
+ Opção 1: o funcionário deve entrar o id do produto que deseja que o cli exiba.
+ Opção 2: o funcionário deve entrar o id do pedido que deseja que o cli exiba.
+ Opção 3: todos os produtos disponíveis para pedir serão listados.
    + Opção 1: cadastrar um novo produto.
        + deve-se inserir nome (string), descrição(string) e preço(float64).
    + Opção 2: deletar um produto pelo seu id
    + Opção 3: retornar
+ Opção 4: mostra os pedidos existentes
    + existe a opção de expedir pedido, a qual vai marcar o primeiro pedido da fila como concluído.
+ Opção 5: mostra:
    + número de produtos cadastrados, número total de pedidos encerrados , número de pedidos em andamento, faturamento total até o momento e tempo médio de expedição.


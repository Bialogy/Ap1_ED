# Requisitos

## Requisitos Funcionais

1. O programa deve incluir a opção de atualizar um produto, modificando o seu preço. O id, o nome e a descrição devem ser mantidos.
2. O programa deve incluir uma nova métrica, chamada ticket médio, que calcula o valor médio de cada pedido (total faturado até o momento, dividido pelo número de pedidos já encerrados).
3. O programa deve incluir uma segunda forma de exibir os produtos, ordenados por nome, e não por id.

## Requisitos Não Funcionais

1. A estrutura de dados que forma a lista de produtos deve ser uma transformada em uma lista simplesmente encadeada. Todas as operações sobre essa lista (adicionar, buscar, excluir, atualizar, exibir etc.) devem ser refatorados para considerar uma lista encadeada.
2. Não há restrição do algoritmo de ordenação utilizado para a exibição dos produtos ordenados pelo nome (pode ser o bubblesort, por exemplo).

## Bugs Identificados e Cenários de Teste

### Bug 1

1. Selecione a opção 6 (adicionar pedido).
2. Insira valores quaisquer para um pedido (p.ex., s, depois 1 1 e depois 0 0).
3. Selecione a opção 7 (expedir pedido).
4. Selecione a opção 7 (expedir pedido).

**Comportamento Esperado:** O programa não deve entrar em um estado de panic e deve resolver a operação corretamente, sem falhas.

### Bug 2

1. Abra o programa, utilizando o arquivo de dados (dados.csv) fornecido.
2. Selecione a opção 5 (exibir os produtos).
3. Verifique os produtos criados.
4. Selecione a opção 1 (cadastrar novo produto).
5. Insira quaisquer valores para nome, descrição e preço.
6. Selecione a opção 5 (exibir os produtos).

**Comportamento Esperado:** O programa deve listar todos os produtos, considerando corretamente os ids dos produtos pré-carregados.

# Critérios de Aceitação

Para o sistema ser aceito pelos clientes, as seguintes condições devem ser satisfeitas:

1. O sistema precisa ser desenvolvido na linguagem Go.
2. O projeto deve ser entregue no seu formato de código fonte, bem como um executável (.exe), construído a partir do fonte.
3. O sistema deve ser entregue em um repositório privado no GitHub, adicionando um dos sócios (@victor0machado) como colaborador.
4. O repositório possui um arquivo readme.md, descrevendo as pessoas da equipe e como utilizar o programa (comandos aceitos, exemplos de respostas, etc.).
5. O não atendimento de algum dos critérios de aceitação pode acarretar em penalização na solução. É recomendado utilizar o mesmo repositório já criado para a fase anterior do projeto.

# Critérios de Avaliação

O sistema será avaliado pelos clientes segundo critérios objetivos e subjetivos, quais sejam:

1. Testes Funcionais (80%): os testes realizados pelos clientes não apresentam falhas. Ou seja, os requisitos funcionais detalhados acima estão sendo atendidos completamente.
2. Avaliação estática do código (20%): o código é eficiente e bem escrito, está organizado, não apresenta redundâncias e atende às boas práticas de programação.

# Demais Informações

- **Prazo do Projeto:** até 27/11/2023, sem possibilidade de adiamento do prazo.
- **Submissão do Link do GitHub:** através do estudante.ibmec.br.
- **Equipes:** podem ser formadas por no mínimo 3 pessoas e no máximo 5.
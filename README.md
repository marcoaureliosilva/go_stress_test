# Sistema de Teste de Carga em Go

Este é um sistema de linha de comando (CLI) desenvolvido em Go para realizar testes de carga em serviços web. O objetivo do projeto é permitir que os usuários possam testar a performance de uma URL específica ao realizar um número determinado de requisições, com a capacidade de definir o nível de concorrência.

## Funcionalidades

- Realiza múltiplas requisições HTTP para uma URL especificada.
- Permite configurar o número total de requests e a quantidade de chamadas simultâneas.
- Gera um relatório com:
  - Tempo total gasto na execução.
  - Quantidade total de requests realizados.
  - Quantidade de requests com status HTTP 200.
  - Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Pré-requisitos

Antes de executar o projeto, certifique-se de que você possui:

- [Go](https://golang.org/dl/) instalado na sua máquina.
- [Docker](https://www.docker.com/products/docker-desktop) instalado na sua máquina.

Execute a aplicação usando Docker:
docker run carga-cli --url=http://exemplo.com --requests=1000 --concurrency=10
Novamente, substitua http://exemplo.com pela URL que deseja testar.

Observações
Certifique-se de que as URLs que você está testando estão acessíveis e que você tem permissão para realizar testes de carga.
Ajuste o número de requisições e a concorrência de acordo com o que for apropriado para o sistema em teste e sua própria rede.
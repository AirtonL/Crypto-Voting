# Klever-Crypto

## Tecnologias útilizadas

* [Golang](https://go.dev/) - Linguagem de programação criada pela Google de código livre
* [gRPC](https://grpc.io/) - Estrutura de código aberto e alto desempenho criada pelo Google
* [MongoDB](https://github.com/mongodb) - Banco de dados não relacional

## Sobre
- O gRPC segue amplamente a semântica HTTP sobre HTTP/2 e, assim permite que você use o streaming full-duplex, possibilitando a comunicação entre diferentes sistemas via conexão de rede.

- Utilizei gRPC com Go para desenvolver um serviço de crypto moedas. Onde o client faz a requisição para o servidor podendo criar, deletar, dar upvote ou downvote e buscar pelo id ou buscar por todas as cryptos.

## Pré requisitos
- Go 1.18.4
- MongoDB 5.0.9 Community
- Protoc 3.6.1

## Como executar o projeto localmente

- Clone o projeto
- Entre na pasta raiz do projeto
- Instale as dependencias necessárias do go.mod
- Rode na pasta raiz os comandos:
```bash
Para iniciar o server: go run server/crypto_grpc.go
Para iniciar o client: go run client/main.go
Certifique-se de estar com o mongoDB up
```
# EndPoints

### POST http://localhost:8080/crypto
#### Request
```bash
{
  "name": "Bitcoin"
}
```
#### Response
```bash
{
  "crypto": {
    "id": "62e1599d556b57cb3d0943a9",
    "name": "Bitcoin"
  }
}
```


### PUT http://localhost:8080/upvote/:id
#### Request
```bash
NoBody
```
#### Response
```bash
{
  "upvote": true
}
```


### PUT http://localhost:8080/downvote/:id
#### Request
```bash
NoBody
```
#### Response
```bash
{
  "downvote": true
}
```


### DELETE http://localhost:8080/crypto/:id
#### Request
```bash
NoBody
```
#### Response
```bash
{
  "delete": true
}
```


### GET http://localhost:8080/crypto
#### Request
```bash
NoBody
```
#### Response
```bash
{
  "crypto": {
    "id": "62e1587c556b57cb3d0943a6",
    "name": "Bitcoin1BTC",
    "upvote": 8,
    "downvote": 3,
    "totalscore": 5
  }
}{
  "crypto": {
    "id": "62e15885556b57cb3d0943a7",
    "name": "Ethereum2ETH",
    "upvote": 6,
    "downvote": 3,
    "totalscore": 3
  }
}{
  "crypto": {
    "id": "62e1588b556b57cb3d0943a8",
    "name": "Tether3USDT",
    "upvote": 3,
    "downvote": 2,
    "totalscore": 1
  }
}
```

## IMPORTANTE
- Todas referencias para criação estão no arquivo references na pasta raiz do projeto

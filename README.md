# Projeto de Demonstração Pub/Sub em Golang

Este é um projeto de demonstração para ilustrar o uso do padrão Publisher/Subscriber (Pub/Sub) utilizando a linguagem Go. O objetivo deste projeto é fornecer uma base prática e fácil de entender para a implementação de sistemas de mensagens assíncronas usando o padrão Pub/Sub em Go.

## Funcionalidades

- **Publisher**: Envia mensagens para um tópico específico.
- **Subscriber**: Recebe mensagens de um tópico ao qual está inscrito.
- **Tópicos**: Canais onde as mensagens são publicadas e dos quais os subscribers recebem mensagens.

## Como Executar

1. **Clone o Repositório:**

    ```sh
    git clone https://github.com/aubermardegan/pubsub.git
    cd pubsub
    ```

2. **Instale as Dependências:**

    Certifique-se de ter o Go instalado e execute:

    ```sh
    go mod tidy
    ```

3. **Execute o Projeto:**

    Para iniciar o publicador e o assinante, execute:

    ```sh
    go run main.go
    ```

## Exemplos de Uso

### Publicador

Para enviar uma mensagem a um tópico:

```go
publisher.Publish("nome_do_tópico", "sua mensagem")
```

### Assinante

Para receber mensagens de um tópico:

```go
sub := publisher.Subscribe("nome_do_tópico")
go sub.Listen()
```

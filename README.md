# go_pro_udemy
    
Curso da Udemy sobre Desenvolvimento Web com Go - Do Zero ao Deploy

![GOlang](/github/banner.png)

# Sumário

1. [Introdução](#introdução)
2. [Configurações do Projeto](#configurações-do-projeto)
3. [Mux](/notes/mux.md#mux)
    * [Manipulando Requisições HTTP no Go](/notes/mux.md#manipulando-requisições-http-no-go)
4. [Organização de Pastas](/notes/structure.md#organização-de-pastas)
    * [Importância de uma Estrutura de Pastas](/notes/structure.md#importância-de-uma-estrutura-de-pastas)
    * [Como Estruturar Pastas de um Projeto](/notes/structure.md#como-estruturar-pastas-de-um-projeto)
    * [Componentes Importantes](/notes/structure.md#componentes-importantes)
5. [SSR vs CSR](/notes/side_render.md#server-side-render-vs-client-side-render)
    * [SSR (Server-Side Rendering)](/notes/side_render.md#ssr-server-side-rendering)
    * [CSR (Client-Side Rendering)](/notes/side_render.md#csr-client-side-rendering)
6. [Templates](/notes/template.md#templates)
    * [O Pacote html/template](/notes/template.md#o-pacote-htmltemplate)
    * [Reuso de Templates com template](/notes/template.md#reuso-de-templates-com-template)
    * [Renderização de Dados Dinâmicos](/notes/template.md#renderização-de-dados-dinâmicos)
    * [Funções Personalizadas nos Templates](/notes/template.md#funções-personalizadas-nos-templates)
    * [Segurança ao Renderizar HTML](/notes/template.md#segurança-ao-renderizar-html)
    * [Desempenho e Cache de Templates](/notes/template.md#desempenho-e-cache-de-templates)
    * [Aplicações com SSR em Golang](/notes/template.md#aplicações-com-ssr-em-golang)
7. [Configurações](/notes/config.md#configurações)
    * [Variáveis de Ambiente](/notes/config.md#variáveis-de-ambiente)
    * [Arquivo .env](/notes/config.md#arquivo-env)
    * [Arquivos de Configuração JSON](/notes/config.md#arquivos-de-configuração-json)
    * [Arquivos .properties](/notes/config.md#arquivos-properties)
    * [Gerenciamento de Configurações por Ambiente (Development, Staging, Production)](/notes/config.md#gerenciamento-de-configurações-por-ambiente-development-staging-production)
    * [Gerenciamento Avançado com Viper](/notes/config.md#gerenciamento-avançado-com-viper)
8. [Logs](/notes/logs.md#logs)
    * [Importância dos Logs](/notes/logs.md#importância-dos-logs)
    * [Tipos de Logs](/notes/logs.md#tipos-de-logs)
    * [Ferramentas de Logging no Go](/notes/logs.md#ferramentas-de-logging-no-go)
        1. [Pacote Padrão log](/notes/logs.md#pacote-padrão-log)
        2. [Pacote slog (Go 1.21)](/notes/logs.md#pacote-slog-go-121)
        3. [Pacotes de Logging Populares](/notes/logs.md#pacotes-de-logging-populares)
    * [Boas Práticas de Logging](/notes/logs.md#boas-práticas-de-logging)
    * [Logging em Ambientes de Produção e Desenvolvimento](/notes/logs.md#logging-em-ambientes-de-produção-e-desenvolvimento)
    * [Integração com Ferramentas de Monitoramento e Observabilidade](/notes/logs.md#integração-com-ferramentas-de-monitoramento-e-observabilidade)
9. [Observações](#observações)
10. [Referências](#referências)

## Introdução

Ao usar Golang (Go) como backend ou API, uma das suas maiores vantagens é a simplicidade e o desempenho elevado que a linguagem proporciona. A biblioteca padrão de Go já oferece todas as ferramentas necessárias para criar servidores HTTP, permitindo manipular rotas, métodos, middlewares e interações com banco de dados de maneira direta e eficiente.

Golang é frequentemente usado como backend devido ao seu desempenho, concorrência eficiente (com goroutines) e suporte nativo a rede. Em APIs RESTful, você geralmente lida com requisições HTTP e responde com JSON. Para fazer isso, o Go oferece uma abordagem simples usando o pacote net/http.

Um exemplo básico de uma API em Go poderia ser:

```go
package main

import (
    "encoding/json"
    "net/http"
)

type Response struct {
    Message string `json:"message"`
}

func main() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        response := Response{Message: "Hello, World!"}
        json.NewEncoder(w).Encode(response)
    })
    http.ListenAndServe(":8080", nil)
}
```

* Estamos criando uma rota /hello que responde com um JSON simples.
* Usamos http.HandleFunc, que aceita uma função que recebe o http.ResponseWriter para escrever a resposta e o *http.Request para lidar com a requisição.

## Configurações do Projeto

* Go version 1.23.1
* Make and Choco
* Docker

* DataBase:  

```bash 
    docker compose up
```

* Server: 

```bash 
    make server
```

ou

```bash 
    go run cmd/api/main.go
```

## Observações

* **Nome dos arquivos:** Não utilizar {nome}.{sub}.go e sim {nome}_{sub}.go

## Referências

- [Udemy](https://www.udemy.com/course/desenvolvimento-web-com-go-do-zero-ao-deploy/)
- [Chat GPT](https://chat.openai.com/)




* Cookies
* Sessions
* CSRF
* SameSite
* emails

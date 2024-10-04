# go_pro_udemy
    
Curso da Udemy sobre Desenvolvimento Web com Go - Do Zero ao Deploy

![GOlang](/github/banner.png)

# Sumário

1. [Introdução](#introdução)
2. [Mux](/notes/mux.md#mux)
    * [Manipulando Requisições HTTP no Go](/notes/mux.md#manipulando-requisições-http-no-go)
3. [Organização de Pastas](/notes/structure.md#organização-de-pastas)
    * [Importância de uma Estrutura de Pastas](/notes/structure.md#importância-de-uma-estrutura-de-pastas)
    * [Como Estruturar Pastas de um Projeto](/notes/structure.md#como-estruturar-pastas-de-um-projeto)
    * [Componentes Importantes](/notes/structure.md#componentes-importantes)
4. [SSR vs CSR](/notes/side_render.md#server-side-render-vs-client-side-render)
    * [SSR (Server-Side Rendering)](/notes/side_render.md#ssr-server-side-rendering)
    * [CSR (Client-Side Rendering)](/notes/side_render.md#csr-client-side-rendering)
    * [Templates](/notes/side_render.md#templates)

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
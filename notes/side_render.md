## Server Side Render vs Client Side Render

No desenvolvimento de aplicações web, existem dois conceitos-chave para renderização de páginas: SSR (Server-Side Rendering) e CSR (Client-Side Rendering). Esses conceitos estão relacionados à maneira como o conteúdo de uma página web é carregado e exibido ao usuário. No contexto de Golang, ambos os métodos podem ser implementados dependendo do tipo de aplicação que você está criando, seja uma aplicação estática, uma SPA (Single Page Application), ou uma aplicação web tradicional.

### SSR (Server-Side Rendering)

Server-Side Rendering refere-se ao processo de renderizar o HTML completo no lado do servidor antes de enviá-lo ao navegador. Ou seja, a página é montada no servidor e entregue completamente ao cliente. Isso significa que o navegador apenas recebe o HTML já renderizado e o exibe ao usuário.

**Como funciona o SSR:**

* O cliente (navegador) faz uma requisição HTTP ao servidor.
* O servidor processa essa requisição, renderiza a página HTML (possivelmente com base em dados dinâmicos) e envia o conteúdo já pronto ao navegador.
* O navegador exibe o HTML renderizado imediatamente.

**Implementando SSR em Golang:**

No Golang, a SSR pode ser feita utilizando templates HTML e o pacote html/template da biblioteca padrão. Isso permite gerar páginas dinâmicas diretamente no servidor.

Exemplo básico de SSR com templates em Go:

```go
package main

import (
    "html/template"
    "net/http"
)

type PageData struct {
    Title   string
    Message string
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := PageData{
            Title:   "SSR with Golang",
            Message: "Hello, this page is rendered on the server side!",
        }

        tmpl, _ := template.ParseFiles("template.html")
        tmpl.Execute(w, data)
    })

    http.ListenAndServe(":8080", nil)
}
```

Arquivo template.html:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
</head>
<body>
    <h1>{{.Message}}</h1>
</body>
</html>
```

Nesse exemplo, o conteúdo HTML é gerado no servidor e enviado completo para o navegador, garantindo que a página seja renderizada antes de ser enviada ao cliente.

**Casos de uso do SSR:**

* Aplicações que precisam de SEO aprimorado.
* Websites que demandam um carregamento inicial rápido.
* Aplicações com conteúdo dinâmico, mas com estrutura de página constante (por exemplo, blogs, sites de notícias).

### CSR (Client-Side Rendering)

Client-Side Rendering refere-se ao processo de carregar uma página básica no navegador e, em seguida, construir e renderizar o conteúdo dinamicamente no lado do cliente usando JavaScript. Com o CSR, o HTML inicial enviado pelo servidor é geralmente apenas uma estrutura mínima, e o conteúdo real é carregado e exibido no navegador por meio de requisições assíncronas (ex.: fetch ou XMLHttpRequest).

**Como funciona o CSR:**

* O cliente (navegador) faz uma requisição HTTP ao servidor.
* O servidor responde com um arquivo HTML básico e os arquivos JavaScript.
* O navegador carrega o JavaScript, que faz requisições adicionais (geralmente APIs) para buscar dados, renderizando o conteúdo dinâmico no cliente.

**Implementando CSR com Golang:**

Para implementações CSR em Golang, o backend geralmente serve como uma API REST que fornece dados para uma aplicação frontend separada (geralmente construída em frameworks JavaScript como React, Vue, ou Svelte). O Go, nesse caso, não está envolvido na renderização de HTML diretamente, mas sim no fornecimento de dados para o frontend que faz a renderização.

Aqui está um exemplo básico de Go servindo uma API que pode ser usada por um frontend para CSR:

```go
package main

import (
    "encoding/json"
    "net/http"
)

type APIResponse struct {
    Message string `json:"message"`
}

func main() {
    http.HandleFunc("/api/message", func(w http.ResponseWriter, r *http.Request) {
        response := APIResponse{Message: "Hello from Go API!"}
        json.NewEncoder(w).Encode(response)
    })

    http.ListenAndServe(":8080", nil)
}
```

No lado do cliente (ex.: com React), a requisição à API poderia ser feita assim:

```javascript
import React, { useEffect, useState } from "react";

function App() {
  const [message, setMessage] = useState("");

  useEffect(() => {
    fetch("/api/message")
      .then((res) => res.json())
      .then((data) => setMessage(data.message));
  }, []);

  return <h1>{message}</h1>;
}

export default App;
```

Aqui, o servidor Go apenas entrega dados JSON, e a aplicação React faz a renderização do conteúdo dinâmico com base nesses dados.

**Casos de uso do CSR:**

* Single Page Applications (SPA): Aplicações web que precisam de interações rápidas e não desejam recarregar toda a página.
* Dashboards e aplicações interativas que dependem de atualizações frequentes e rápidas do conteúdo.
* Cenários em que o SEO não é uma prioridade ou pode ser gerenciado de outras formas (ex.: aplicações internas, painéis administrativos).
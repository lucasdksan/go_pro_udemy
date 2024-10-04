## Templates

Usar HTML no Golang envolve a construção de aplicações web que renderizam páginas no lado do servidor (SSR - Server-Side Rendering). Golang oferece suporte nativo para o processamento e renderização de templates HTML por meio do pacote padrão html/template. Este pacote é robusto e permite a criação de templates dinâmicos, a integração de dados e o reuso de componentes, mantendo a segurança contra ataques de injeção de código.

### O Pacote html/template

O pacote html/template fornece uma maneira segura e eficiente de gerar HTML dinâmico em Golang. Ele suporta a utilização de variáveis, condicionais, laços, funções personalizadas, e uma série de mecanismos para a construção de páginas complexas.

A principal vantagem desse pacote é que ele automaticamente faz a escapagem de caracteres especiais, como &, <, >, e " ao renderizar conteúdo dinâmico, prevenindo ataques de Cross-Site Scripting (XSS).

```go
package main

import (
    "html/template"
    "log"
    "net/http"
)

type PageData struct {
    Title   string
    Message string
}

func main() {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := PageData{
            Title:   "Welcome to Golang",
            Message: "Hello, this is a dynamic message.",
        }

        tmpl.Execute(w, data)
    })

    log.Println("Server running at :8080")
    http.ListenAndServe(":8080", nil)
}
```

Arquivo index.html:

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

Neste exemplo, o template index.html utiliza placeholders ({{.Title}} e {{.Message}}) que são preenchidos com os dados dinâmicos fornecidos no handler. O método template.Execute() cuida da injeção de dados no template HTML.

### Reuso de Templates com template

O pacote html/template também suporta a reutilização de templates e fragmentos HTML por meio de blocos e definições de templates parciais. Isso facilita a criação de layouts reutilizáveis, como cabeçalhos, rodapés e layouts de páginas comuns.

```bash
my-project/
├── templates/
│   ├── layout.html       # Template de layout base
│   ├── index.html        # Conteúdo principal da página
│   └── partials/
│       ├── header.html   # Cabeçalho compartilhado
│       └── footer.html   # Rodapé compartilhado
```

Arquivo layout.html (layout base com placeholders):

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
</head>
<body>
    <header>{{template "header.html" .}}</header>

    <main>
        {{block "content" .}} {{end}}
    </main>

    <footer>{{template "footer.html" .}}</footer>
</body>
</html>
```

Arquivo index.html (conteúdo específico da página):

```html
{{define "content"}}
    <h1>{{.Message}}</h1>
    <p>This is the homepage.</p>
{{end}}
```

Handler no Go:

```go
package main

import (
    "html/template"
    "log"
    "net/http"
)

type PageData struct {
    Title   string
    Message string
}

func main() {
    tmpl := template.Must(template.ParseFiles(
        "templates/layout.html",
        "templates/index.html",
        "templates/partials/header.html",
        "templates/partials/footer.html",
    ))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := PageData{
            Title:   "Home",
            Message: "Welcome to the homepage!",
        }
        tmpl.ExecuteTemplate(w, "layout.html", data)
    })

    log.Println("Server running at :8080")
    http.ListenAndServe(":8080", nil)
}
```

* {{template "header.html" .}}: Inclui o conteúdo do template parcial header.html.
* {{block "content" .}} {{end}}: Define um bloco onde o conteúdo específico da página (index.html, neste caso) será injetado.
* ExecuteTemplate(): Renderiza o template base layout.html, injetando os dados e conteúdo das outras partes do template.

### Renderização de Dados Dinâmicos

A renderização dinâmica permite injetar dados de forma flexível nos templates. Você pode passar praticamente qualquer dado, desde variáveis simples a objetos e listas, e manipulá-los diretamente no template com expressões condicionais e laços.

```html
{{define "content"}}
<h1>{{.Title}}</h1>

{{if .Items}}
<ul>
    {{range .Items}}
    <li>{{.}}</li>
    {{end}}
</ul>
{{else}}
<p>No items available.</p>
{{end}}
{{end}}
```

Handler no Go:

```go
package main

import (
    "html/template"
    "log"
    "net/http"
)

type PageData struct {
    Title string
    Items []string
}

func main() {
    tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/list.html"))

    http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
        data := PageData{
            Title: "Item List",
            Items: []string{"Item 1", "Item 2", "Item 3"},
        }
        tmpl.ExecuteTemplate(w, "layout.html", data)
    })

    log.Println("Server running at :8080")
    http.ListenAndServe(":8080", nil)
}
```

* range é usado para iterar sobre uma lista de itens.
* if e else fornecem controle condicional para exibir conteúdo dinamicamente.

### Funções Personalizadas nos Templates

Golang permite que você registre funções personalizadas para serem usadas dentro dos templates, o que expande bastante as capacidades de formatação e manipulação de dados.

```go
package main

import (
    "html/template"
    "log"
    "net/http"
    "strings"
)

type PageData struct {
    Title string
}

func main() {
    tmpl := template.New("index.html").Funcs(template.FuncMap{
        "toUpper": strings.ToUpper,
    })
    tmpl, _ = tmpl.ParseFiles("templates/index.html")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := PageData{
            Title: "hello world",
        }
        tmpl.Execute(w, data)
    })

    log.Println("Server running at :8080")
    http.ListenAndServe(":8080", nil)
}
```

Arquivo index.html:

```html
<h1>{{.Title | toUpper}}</h1>
```

Aqui, a função toUpper converte o valor da variável Title para letras maiúsculas antes de renderizá-lo no HTML.

### Segurança ao Renderizar HTML

O pacote html/template faz automaticamente a escapagem de dados dinâmicos, impedindo que strings perigosas (como código JavaScript ou tags HTML) sejam executadas no navegador do usuário. Isso previne ataques de injeção de script (XSS).

```go
data := PageData{
    Title:   "Hello",
    Message: "<script>alert('XSS')</script>",
}
```

O Go automaticamente converte o conteúdo para uma string segura:

```bash
&lt;script&gt;alert(&#39;XSS&#39;)&lt;/script&gt;
```

Isso garante que o conteúdo seja exibido como texto, sem ser executado como código malicioso.

### Desempenho e Cache de Templates

Para melhorar o desempenho em produção, você pode pré-compilar e armazenar os templates em cache, evitando o custo de carregá-los e analisá-los a cada requisição. O template.Must() pode ser usado para carregar os templates uma vez no início do servidor.

### Aplicações com SSR em Golang

Golang é muito eficiente para SSR (Server-Side Rendering), ideal para aplicações que precisam servir HTML dinâmico diretamente do servidor. Isso é especialmente útil em cenários onde SEO é uma prioridade ou em sites onde o conteúdo precisa ser entregue rapidamente ao usuário, sem a necessidade de uma aplicação frontend separada, como em React ou Vue.
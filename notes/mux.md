## Mux

O mux (ou multiplexador) é uma peça fundamental no roteamento de requisições HTTP. Ele é responsável por mapear diferentes padrões de URL para seus respectivos handlers (manipuladores). O pacote padrão de Go (net/http) fornece um mux básico chamado http.ServeMux, mas também existem outras bibliotecas populares como gorilla/mux que oferecem mais funcionalidades.

O mux básico (http.ServeMux) permite mapear rotas e padrões de URL para handlers de uma maneira simples e eficiente.

```go
mux := http.NewServeMux()

mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
})

mux.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Goodbye, World!"))
})

http.ListenAndServe(":8080", mux)
```

* Estamos usando o ServeMux para roteamento de duas rotas: /hello e /goodbye.
* Cada rota tem sua própria função que responde a requisições.

### Manipulando Requisições HTTP no Go

Existem três formas principais de manipular requisições HTTP no Go: Handle, HandleFunc e Handler. Esses métodos são usados para associar padrões de URL a funções que tratam as requisições.

1. Handle

Handle associa uma rota a um handler que implementa a interface http.Handler. Um http.Handler é qualquer tipo que tenha o método ServeHTTP com a seguinte assinatura:

```go
type Handler interface {
    ServeHTTP(w http.ResponseWriter, r *http.Request)
}
```

Isso permite flexibilidade para criar tipos customizados para lidar com requisições.

```go
type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Custom Handler"))
}

http.Handle("/custom", MyHandler{})
```

Aqui, a struct MyHandler implementa a interface http.Handler e pode ser passada diretamente para http.Handle.

2. HandleFunc

HandleFunc é uma conveniência que permite associar diretamente uma função a uma rota. A função precisa ter a assinatura func(http.ResponseWriter, *http.Request):

```go
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, HandleFunc!"))
})
```

Isso é útil para casos onde você não precisa de um handler complexo e pode lidar com a lógica diretamente na função.

3. Handler

O Handler é uma interface que define a base para todo o sistema de tratamento de requisições no Go. Qualquer coisa que implemente a interface http.Handler pode ser registrada para manipular uma rota.

Por exemplo, além de http.HandleFunc e http.Handle, você pode criar um handler completo implementando ServeHTTP em uma struct customizada para encapsular mais lógica.

```go
type LoggingHandler struct {
    next http.Handler
}

func (h LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    log.Printf("Request: %s %s", r.Method, r.URL.Path)
    h.next.ServeHTTP(w, r)
}

func main() {
    myHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello with Logging"))
    })

    loggingHandler := LoggingHandler{next: myHandler}
    http.Handle("/log", loggingHandler)
    http.ListenAndServe(":8080", nil)
}
```
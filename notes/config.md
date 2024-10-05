## Configurações 

Configurar aplicações web em projetos Golang é uma tarefa essencial para garantir que o código seja flexível, seguro e escalável. Para isso, utilizam-se variáveis de ambiente, arquivos de configuração como .env, .json ou .properties, e outras abordagens que permitem externalizar configurações de maneira eficiente. Vamos abordar cada uma dessas estratégias e como elas podem ser implementadas em Golang.

### Variáveis de Ambiente

As variáveis de ambiente são uma maneira comum e segura de armazenar configurações sensíveis, como credenciais de banco de dados, chaves de API, e informações de configuração. Elas são especialmente úteis porque não precisam ser incluídas no código fonte, o que melhora a segurança e facilita a configuração em diferentes ambientes (desenvolvimento, teste, produção).

**Como ler variáveis de ambiente no Go**

O pacote padrão os oferece funções para trabalhar com variáveis de ambiente. O método os.Getenv() permite recuperar valores de variáveis de ambiente, enquanto os.LookupEnv() também retorna um segundo valor indicando se a variável foi definida ou não.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Obtém a variável de ambiente "PORT"
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"  // Valor padrão se a variável não estiver definida
    }

    fmt.Println("Server running on port:", port)
}
```

Usos Comuns:

* Configuração do banco de dados: DB_HOST, DB_USER, DB_PASSWORD.

* Configurações do servidor web: PORT, HOST.

* Chaves de API: API_KEY, SECRET_KEY.

### Arquivo .env

O arquivo .env é uma maneira popular de armazenar variáveis de ambiente em formato texto, geralmente utilizado em ambientes de desenvolvimento local. Esse arquivo contém pares de chave-valor que podem ser carregados automaticamente no ambiente de execução da aplicação.

Para trabalhar com arquivos .env no Go, uma das bibliotecas mais utilizadas é o github.com/joho/godotenv, que facilita o carregamento de variáveis de ambiente diretamente de um arquivo .env.

Exemplo de um arquivo .env:

```env
PORT=8080
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=secret
API_KEY=your-api-key-here
```

Como usar .env no Go:

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
)

func main() {
    // Carrega as variáveis do arquivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Usa as variáveis carregadas
    port := os.Getenv("PORT")
    fmt.Println("Server running on port:", port)
}
```

Aqui, o arquivo .env é carregado no ambiente de execução, e as variáveis podem ser acessadas com os.Getenv().

### Arquivos de Configuração JSON

Outra abordagem comum para armazenar configurações é usar arquivos JSON. Eles são fáceis de ler e escrever, além de serem amplamente utilizados. Em projetos Go, você pode utilizar o pacote padrão encoding/json para ler e interpretar arquivos JSON.

Exemplo de um arquivo config.json:

```json
{
    "port": "8080",
    "db": {
        "host": "localhost",
        "user": "root",
        "password": "secret"
    },
    "api_key": "your-api-key-here"
}
```

Carregando configurações JSON no Go:

```go
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

type Config struct {
    Port   string `json:"port"`
    DB     struct {
        Host     string `json:"host"`
        User     string `json:"user"`
        Password string `json:"password"`
    } `json:"db"`
    APIKey string `json:"api_key"`
}

func main() {
    // Abre o arquivo config.json
    file, err := os.Open("config.json")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Lê o conteúdo do arquivo
    byteValue, _ := ioutil.ReadAll(file)

    // Carrega o conteúdo em uma struct
    var config Config
    json.Unmarshal(byteValue, &config)

    fmt.Println("Server running on port:", config.Port)
    fmt.Println("Database User:", config.DB.User)
}
```

### Arquivos .properties

Os arquivos .properties são comuns em linguagens como Java, mas também podem ser usados no Go para armazenar configurações simples em formato de pares chave-valor.

Exemplo de um arquivo config.properties:

```properties
port=8080
db.host=localhost
db.user=root
db.password=secret
api.key=your-api-key-here
```

Como carregar arquivos .properties no Go:

Embora não haja suporte nativo para arquivos .properties, você pode usar pacotes como github.com/magiconair/properties para lidar com esse formato.

```go
package main

import (
    "fmt"
    "log"

    "github.com/magiconair/properties"
)

func main() {
    // Carrega o arquivo config.properties
    p := properties.MustLoadFile("config.properties", properties.UTF8)

    // Usa as propriedades carregadas
    port := p.MustGetString("port")
    dbUser := p.MustGetString("db.user")

    fmt.Println("Server running on port:", port)
    fmt.Println("Database User:", dbUser)
}
```

### Gerenciamento de Configurações por Ambiente (Development, Staging, Production)

Ao desenvolver aplicações que serão executadas em diferentes ambientes (desenvolvimento, homologação, produção), é importante configurar o projeto para que ele carregue as configurações adequadas para cada ambiente.

**Estratégias para múltiplos ambientes:**

* Arquivos .env múltiplos: Ter arquivos .env específicos para cada ambiente, como .env.development, .env.production.

No código Go, você pode carregar o arquivo correto com base em uma variável de ambiente:

```env
env := os.Getenv("GO_ENV") // production, development, etc.
godotenv.Load(fmt.Sprintf(".env.%s", env))
```

* Flags de linha de comando: Usar flags para definir configurações em tempo de execução. Isso é útil para ambientes de produção, onde você pode passar diferentes parâmetros diretamente na linha de comando ou script de deploy.

Exemplo usando o pacote flag:

```go
package main

import (
    "flag"
    "fmt"
)

func main() {
    port := flag.String("port", "8080", "Port to run the server on")
    flag.Parse()

    fmt.Println("Server running on port:", *port)
}
```

### Gerenciamento Avançado com Viper

Uma solução mais avançada e flexível para gerenciamento de configurações é o uso da biblioteca Viper (github.com/spf13/viper), que suporta diferentes formatos de arquivos (JSON, YAML, TOML), variáveis de ambiente, e flags de linha de comando.

Exemplo básico com Viper:

```go
package main

import (
    "fmt"
    "log"

    "github.com/spf13/viper"
)

func main() {
    // Configura o nome do arquivo de configuração (sem extensão)
    viper.SetConfigName("config")

    // Configura o caminho onde o arquivo está
    viper.AddConfigPath(".")

    // Tenta ler o arquivo de configuração
    if err := viper.ReadInConfig(); err != nil {
        log.Fatal("Error reading config file:", err)
    }

    // Recupera as configurações
    port := viper.GetString("port")
    dbUser := viper.GetString("db.user")

    fmt.Println("Server running on port:", port)
    fmt.Println("Database User:", dbUser)
}
```

Esse exemplo mostra como o Viper pode ser usado para carregar configurações de um arquivo JSON, YAML, ou outro formato suportado, além de integrar variáveis de ambiente e linhas de comando.
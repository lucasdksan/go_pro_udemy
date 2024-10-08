## Logs

A utilização de logs em um backend em Go é essencial para a observabilidade, manutenção e depuração da aplicação. Logs ajudam os desenvolvedores e operadores a entender o que está acontecendo na aplicação durante a execução, fornecendo informações cruciais sobre o fluxo, comportamento, e, especialmente, erros ou anomalias.

### Importância dos Logs

* **Depuração:** Durante o desenvolvimento, os logs são uma das formas mais diretas de entender o que a aplicação está fazendo, identificando falhas ou fluxos inesperados.

* **Auditoria e Monitoramento:** Em produção, logs servem como um registro das atividades, permitindo auditorias para entender ações realizadas pelos usuários e sistemas, além de monitorar o comportamento da aplicação.

* **Segurança:** Logs podem registrar atividades suspeitas ou anômalas, ajudando a detectar e investigar tentativas de ataque ou falhas de segurança.

* **Diagnóstico de Performance:** É possível identificar gargalos de performance ou comportamento inadequado através de análises de logs de tempo de resposta ou execuções anormalmente longas.

### Tipos de Logs

* **Logs de Depuração (Debug logs):** Usados durante o desenvolvimento e para encontrar bugs. Detalham cada ação importante da aplicação.

* **Logs Informativos (Info logs):** Usados para relatar eventos normais que ocorrem no sistema. Exemplos incluem inicialização de serviços, requisições recebidas, entre outros.

* **Logs de Aviso (Warning logs):** Indicativos de situações anômalas que não são necessariamente erros, mas que requerem atenção.

* **Logs de Erro (Error logs):** Reportam erros que precisam ser resolvidos. São extremamente valiosos para diagnósticos.

* **Logs de Fatal (Fatal logs):** Logs que indicam uma falha crítica, onde a aplicação não pode continuar executando.

### Ferramentas de Logging no Go

#### Pacote Padrão log

O pacote padrão log do Go oferece funcionalidades básicas de logging e é fácil de usar. Ele escreve mensagens diretamente para a saída padrão (stdout) ou para arquivos, sendo simples, mas limitado quando comparado a soluções mais robustas.

Exemplo básico de uso:

```go
package main

import (
    "log"
)

func main() {
    log.Println("This is an info log")
    log.Fatal("This is a fatal log") // Termina a execução da aplicação
}
```

Limitações do pacote log:

* Não oferece níveis de log (como debug, warning, error).
* Formatação limitada (ex.: logs sempre em texto simples).
* Falta de suporte nativo para ambientes estruturados, como JSON.

#### Pacote slog (Go 1.21)

Com a versão 1.21, o Go introduziu o pacote slog, oferecendo um logger mais poderoso, com suporte a logs estruturados e maior flexibilidade para registro e filtragem de logs.

Exemplo de uso:

```go
package main

import (
    "os"
    "golang.org/x/exp/slog"
)

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout)) // Logs em JSON

    logger.Info("This is an info log", "user", "Alice", "action", "login")
    logger.Error("This is an error log", "error", "Invalid credentials")
}
```

Vantagens:

* Suporte a logs estruturados (como JSON).
* Flexibilidade para adicionar campos contextuais (ex.: informações de usuários, transações, etc.).
* Níveis de log claros (info, error, etc.).
* Suporte a filtros e handlers para rotear ou redirecionar logs conforme necessário.

#### Pacotes de Logging Populares

* **logrus**

logrus é uma biblioteca popular de logging para Go que oferece uma API rica, com suporte a níveis de log, formatação e hooks.

Exemplo de uso:

```go
package main

import (
    log "github.com/sirupsen/logrus"
)

func main() {
    log.SetFormatter(&log.JSONFormatter{}) // Logs em formato JSON

    log.WithFields(log.Fields{
        "user": "Bob",
        "action": "purchase",
    }).Info("User action")
}
```

Vantagens:

* Fácil de usar e amplamente adotado.
* Suporte a logs estruturados.
* Hooks personalizáveis para integração com sistemas externos (ex.: enviar logs para serviços remotos).

* **zap**

zap é uma biblioteca de logging altamente performática, projetada para alta performance e baixa latência. É especialmente útil em sistemas de produção de grande escala, onde a performance é crítica.

Exemplo de uso:

```go
package main

import (
    "go.uber.org/zap"
)

func main() {
    logger, _ := zap.NewProduction()
    defer logger.Sync() // Flush de logs pendentes

    logger.Info("This is an info log", zap.String("user", "Alice"), zap.String("action", "login"))
}
```

Vantagens:

* Alto desempenho com otimizações como logging zerado em alocações.
* Suporte a logs estruturados.
* Amplamente usado em sistemas de produção de alta escala.

### Boas Práticas de Logging

* **Níveis de Log Adequados:** Sempre use o nível de log correto. Logs de depuração (debug) devem ser usados durante o desenvolvimento, enquanto logs de erro (error) devem ser usados para capturar falhas.

* **Logs Estruturados:** Sempre que possível, prefira logs estruturados (ex.: JSON) sobre logs de texto simples. Isso facilita a análise automática e o envio para sistemas como ELK (Elastic Stack) ou Prometheus.

* **Contexto e Detalhes:** Sempre inclua o máximo de contexto relevante no log. Por exemplo, ao registrar uma falha de login, inclua informações como o ID do usuário, o timestamp, e a ação realizada.

* **Evitar Logs Excessivos em Produção:** Em ambientes de produção, evite logar tudo. Use configurações de nível de log (ex.: INFO ou ERROR) e desative logs de depuração (DEBUG).

* **Logs Assíncronos:** Se a performance for crítica, considere o uso de loggers assíncronos para reduzir o impacto da escrita de logs no desempenho.

* **Rotação de Logs:** Utilize mecanismos de rotação de logs (log rotation) para evitar o crescimento excessivo de arquivos de log em disco. Ferramentas como o logrotate em sistemas Unix ajudam nisso.

### Logging em Ambientes de Produção e Desenvolvimento

* **Desenvolvimento:** Em ambientes de desenvolvimento, use níveis de log mais verbosos (ex.: DEBUG ou INFO), e prefira logs em formato legível por humanos.

* **Produção:** Em produção, use níveis de log menos verbosos (ex.: ERROR e WARN), e prefira logs estruturados, como JSON. Integre logs a sistemas de monitoramento (como o ELK Stack, Loki, etc.) para coleta, análise e visualização de logs.

### Integração com Ferramentas de Monitoramento e Observabilidade

Para ambientes de produção, você pode integrar os logs com ferramentas como:

* **ELK Stack (Elastic, Logstash, Kibana):** Para centralizar logs, realizar buscas avançadas e visualizações.

* **Grafana Loki:** Um sistema de gerenciamento de logs que integra-se facilmente com o Grafana, permitindo monitoramento e observação em tempo real.

* **Prometheus e Alertmanager:** Embora Prometheus não seja um sistema de logging, ele pode se integrar com logs de sistemas para gerar alertas baseados em métricas associadas aos logs.
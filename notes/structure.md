## Organização de Pastas

A estrutura de pastas em um projeto de software é fundamental para manter a organização, escalabilidade e manutenção do código. Um bom layout de pastas ajuda a separar responsabilidades, melhorar a legibilidade e facilita o trabalho em equipe. Além disso, uma estrutura consistente torna o projeto mais acessível para novos desenvolvedores, ajuda na navegação entre arquivos, e facilita a aplicação de padrões de design e boas práticas.

### Importância de uma Estrutura de Pastas

* **Organização e Escalabilidade:** Conforme o projeto cresce, mais arquivos e funcionalidades são adicionados. Ter uma estrutura de pastas bem definida ajuda a controlar essa complexidade, agrupando arquivos de maneira lógica e previsível.

* **Facilidade de Manutenção:** Um código bem organizado facilita a manutenção ao longo do tempo. Quando há uma divisão clara de responsabilidades entre pastas, é mais fácil localizar onde fazer mudanças ou adicionar novas funcionalidades sem causar problemas em outras partes do sistema.

* **Colaboração:** Em equipes de desenvolvimento, uma estrutura clara facilita a colaboração, pois todos sabem onde encontrar e adicionar novas funcionalidades. Isso reduz a curva de aprendizado para novos membros do time e previne conflitos desnecessários.

* **Reutilização e Modificação:** Uma boa estrutura permite que componentes e módulos sejam reutilizados em diferentes partes do sistema, além de isolar mudanças. Isso também contribui para uma implementação mais fácil de testes.

### Como Estruturar Pastas de um Projeto

A estrutura de pastas pode variar dependendo da linguagem, do framework e das convenções do projeto, mas algumas diretrizes são comuns a muitos projetos:

1. Divisão por Funcionalidade vs. Divisão por Camadas:

    * **Funcionalidade (Feature-based):** Agrupa arquivos relacionados a uma funcionalidade específica. Muito usada em pr ojetos com microsserviços, API e grandes sistemas.

    * **Camadas (Layer-based):** Separa os arquivos com base nas camadas do sistema (ex.: controladores, modelos, repositórios). Essa é comum em projetos tradicionais baseados em MVC ou arquitetura em camadas.

2. Pasta Raiz do Projeto: A pasta raiz deve conter apenas os arquivos essenciais para o projeto, como arquivos de configuração (por exemplo, README.md, go.mod, package.json, Dockerfile) e diretórios principais. Evite encher a pasta raiz com arquivos de código, agrupando-os em subdiretórios.

### Componentes Importantes

1. **cmd/:** Contém os entrypoints da aplicação. Por exemplo, em projetos Go, aqui ficam os pacotes que têm o main() e rodam diferentes componentes, como o servidor HTTP.

2. **internal/:** Usado para definir pacotes que são internos ao projeto e que não devem ser acessíveis por outros pacotes externos. Cada funcionalidade (ou domínio) deve ter sua própria pasta, com subdiretórios para organizar a lógica de negócio, adapters (integrações externas), e testes.

3. **pkg/:** Para pacotes que podem ser reutilizados em diferentes partes do projeto. Como a pasta internal/, mas com o objetivo de ser exportada para outros projetos.

4. **vendor/:** Dependências de aplicativos (gerenciadas manualmente ou por sua ferramenta de gerenciamento de dependências favorita, como o novo recurso integrado Go Modules). O comando go mod vendor criará o diretório /vendor para você. Note que você pode precisar adicionar a flag -mod=vendor ao seu comando go build se você não estiver usando Go 1.14 onde ele está ativado por padrão.

5. **api/:** Define as rotas e handlers da API. Em projetos mais complexos, pode incluir subdiretórios para agrupar controllers, middlewares e schemas de request/response.

6. **web/:** Componentes específicos de aplicativos da Web: ativos estáticos da Web, modelos do lado do servidor e SPAs.

7. **config/:** Contém arquivos de configuração, como .env, arquivos YAML, ou JSON que controlam os parâmetros do ambiente (ex.: banco de dados, variáveis de ambiente).

8. **init/:** Configurações de inicialização do sistema (systemd, upstart, sysv) e gerenciador/supervisor de processos (runit, supervisord).

9. **scripts/:** Scripts que automatizam tarefas importantes, como inicialização de banco de dados, geração de código, migrações, etc.

10. **build/:** Empacotamento e integração contínua.

Coloque suas configurações de pacote e scripts em nuvem (AMI), contêiner (Docker), sistema operacional (deb, rpm, pkg) no diretório /build/package.

Coloque suas configurações e scripts de CI (travis, circle, drone) no diretório /build/ci. Observe que algumas das ferramentas de CI (por exemplo, Travis CI) são muito exigentes quanto à localização de seus arquivos de configuração. Tente colocar os arquivos de configuração no diretório /build/ci vinculando-os ao local onde as ferramentas de CI os esperam (quando possível).

11. **deployments/:** IaaS, PaaS, configurações e modelos de implantação de orquestração de sistema e contêiner (docker-compose, kubernetes / helm, mesos, terraform, bosh). Observe que em alguns repositórios (especialmente em aplicativos implantados com kubernetes), esse diretório é denominado /deploy.

12. **tests/:** Aplicações de testes externos adicionais e dados de teste. Sinta-se à vontade para estruturar o diretório /test da maneira que quiser. Para projetos maiores, faz sentido ter um subdiretório de dados. Por exemplo, você pode ter /test/data ou /test/testdata se precisar que o Go ignore o que está naquele diretório.

13. **docs/:** Armazena documentação gerada, como especificações OpenAPI/Swagger, manuais de uso ou de desenvolvimento.

14. **tools/:** Ferramentas de suporte para este projeto. Observe que essas ferramentas podem importar código dos diretórios /pkg e /internal.

15. **examples/:** Exemplos para seus aplicativos e / ou bibliotecas públicas.

16. **third_party/:** Ferramentas auxiliares externas, código bifurcado e outros utilitários de terceiros (por exemplo, Swagger UI).
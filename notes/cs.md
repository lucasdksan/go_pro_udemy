## Cookies

Cookies são pequenos arquivos de texto armazenados no navegador do usuário por um site web. Eles são usados para manter informações de estado entre requisições HTTP, que são por natureza stateless. Com cookies, um site pode lembrar as ações e preferências de um usuário (como login, idioma, itens em um carrinho de compras) ao longo de diferentes sessões ou páginas.

### Tipos de Cookies:

* **Session Cookies:** Armazenados temporariamente e apagados quando o navegador é fechado.
* **Persistent Cookies:** Têm uma data de expiração definida, permanecendo no navegador até que expirem ou sejam manualmente deletados pelo usuário.
* **Secure Cookies:** Só são enviados por conexões seguras (HTTPS), garantindo maior proteção.
* **HttpOnly Cookies: Não podem ser acessados via JavaScript, protegendo contra ataques como XSS (Cross-Site Scripting).
* **SameSite Cookies:** Controlam se os cookies são enviados ou não com requisições cross-site, importante para mitigar ataques CSRF (Cross-Site Request Forgery).

## Sessions

Sessões (sessions) são um mecanismo usado pelos servidores para armazenar informações sobre o estado de um usuário durante a interação com um site. Ao contrário dos cookies, que são armazenados no cliente, as sessões geralmente são mantidas no servidor. Uma identificação única da sessão (geralmente um session ID) é enviada ao cliente, que a armazena em um cookie ou a inclui na URL.

### Características das Sessões:

* **Stateful:** As sessões mantêm o estado entre múltiplas requisições HTTP. O estado de um usuário pode ser mantido mesmo que as requisições sejam independentes.

* **Escopo:** Normalmente, uma sessão dura enquanto o navegador estiver aberto, ou até o tempo de inatividade configurado ser alcançado.

* **Segurança:** Sessões são frequentemente protegidas por mecanismos como HttpOnly e Secure Cookies, bem como por tokens de segurança (por exemplo, JWT em sistemas modernos).

Sessões são amplamente usadas para login de usuários. Quando o usuário autentica, uma sessão é criada no servidor e um session ID é armazenado no navegador, permitindo que o servidor reconheça aquele usuário em requisições subsequentes.

## CSRF (Cross-Site Request Forgery)

CSRF é um tipo de ataque onde um usuário autenticado é induzido a executar uma ação indesejada em uma aplicação web sem seu consentimento. O atacante usa a sessão ou autenticação existente da vítima para enviar requisições maliciosas ao servidor.

### Como funciona o CSRF:

1. O usuário faz login em um site (ex: banco online) e obtém um cookie de sessão.

2. O atacante engana o usuário, fazendo-o clicar em um link ou carregar um script malicioso que faz uma requisição ao site legítimo.

3. O site legítimo processa a requisição acreditando que foi enviada pelo usuário autenticado.

### Proteção contra CSRF:

* Tokens Anti-CSRF (CSRF Tokens): Um token único gerado pelo servidor e incluído em formulários ou cabeçalhos de requisições. O servidor valida o token para garantir que a requisição foi originada da aplicação correta.

* SameSite Cookies: Configurar cookies com SameSite=Lax ou SameSite=Strict para impedir que cookies sejam enviados com requisições cross-origin, limitando o risco de ataques CSRF.

## SameSite

O atributo SameSite de cookies controla se um cookie pode ser enviado com requisições cross-site. Isso é uma camada adicional de proteção para evitar que sites terceiros possam realizar requisições não autorizadas utilizando cookies válidos de outra origem.

### Valores de SameSite:

* **Strict:** O cookie não será enviado em nenhuma requisição cross-origin. Isso oferece a maior proteção, mas pode prejudicar a usabilidade (ex: login não funcionará ao abrir links de outros sites).

* **Lax:** O cookie só é enviado em requisições de navegação de nível superior (ex: links), mas não em requisições automáticas como POST ou carregamento de imagens de outros domínios.

* **None:** O cookie é enviado em qualquer requisição, incluindo cross-site. Para ser seguro, este valor precisa ser usado com Secure, ou seja, apenas em conexões HTTPS.

### Uso em Segurança:

O uso do atributo SameSite é altamente recomendado como parte da defesa contra CSRF. Ao impedir que cookies sejam enviados automaticamente em requisições de outros sites, ele ajuda a limitar ataques onde um atacante tenta forjar uma requisição autenticada.
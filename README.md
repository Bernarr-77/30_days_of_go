# 🚀 30 Dias de Go: Do Zero até Backend Completo

**Objetivo Final:** Um backend Go completo, funcional, estruturado e pronto pra usar.

**Filosofia:** Projeto-driven. Você constrói coisas reais, não faz exercícios chatos. **Sem código entregue — você escreve tudo.**

---

## ⚡ Semana 1: Fundamentos + Primeira CLI Tool
**Meta:** Entender sintaxe, tipos, funções. Sair com uma CLI tool funcional.

### Dia 1: Setup e "Hello, World"
- Instale Go (https://golang.org/dl)
- Configure variáveis de ambiente (`$GOPATH`, `$GOROOT`)
- Crie um arquivo `main.go` que imprime uma mensagem
- Entenda os blocos: `package main`, `import`, `func main()`
- Execute com `go run main.go`

**Desafio:** O programa deve rodar sem erros.

**Entrega:** Arquivo `main.go` no GitHub.

---

### Dia 2: Tipos, Variáveis e Operadores
- Aprenda os tipos primitivos de Go: inteiros, floats, strings, booleanos
- Declaração de variáveis: sintaxe `var` vs inferência de tipo (`:=`)
- Operadores aritméticos, comparação e lógicos
- Leia entrada do usuário via `bufio.Scanner` ou `fmt.Scan()`
- **Projeto:** Construa uma calculadora que lê dois números do terminal e realiza operações básicas (soma, subtração, multiplicação, divisão)

**Desafio:** A calculadora deve validar entrada (ex: divisão por zero)

**Entrega:** `calculator.go` funcional.

---

### Dia 3: Funções e Controle de Fluxo
- Defina funções com parâmetros e retorno
- Explore retorno múltiplo (funções que retornam dados + erro)
- Entenda `if/else` e `for loops`
- Tratamento básico de erros: padrão `if err != nil`
- **Projeto:** Crie um jogo de adivinhação onde:
  - Programa gera um número secreto (1-100)
  - Usuário tenta adivinhar em um loop
  - A cada tentativa, diga "muito alto" ou "muito baixo"
  - Conte quantas tentativas levou

**Desafio:** Adicione limite de tentativas (ex: máximo 10). Valide entrada.

**Entrega:** `game.go` com lógica completa.

---

### Dia 4: Arrays, Slices e Maps
- Entenda a diferença: arrays têm tamanho fixo, slices são dinâmicos
- Operações em slices: `append()`, `len()`, `cap()`, indexação
- Maps (hash tables): inserção, busca, iteração
- Loop com `range` pra arrays, slices e maps
- **Projeto:** Construa um sistema de registro de contatos:
  - Armazene em um map: nome → número de telefone
  - Menu com opções: adicionar, remover, buscar, listar contatos
  - Tudo em memória (sem banco ainda)

**Desafio:** Valide números telefônicos (deve ter 11 dígitos). Impeça duplicação de nomes.

**Entrega:** `contacts.go` com CRUD funcionando.

---

### Dia 5: Structs e Métodos
- Defina estruturas (structs) pra representar dados (ex: `Person`, `Product`)
- Crie métodos associados à struct (funções que "pertencem" a um tipo)
- Entenda ponteiros: `*Type` e `&variable`
- **Projeto:** Refatore o sistema de contatos:
  - Use uma struct `Contact` com campos: ID, Name, Email, Phone
  - Crie métodos na struct: `Display()`, `Validate()`
  - Mantenha lista de contatos como `[]Contact`

**Desafio:** O método `Validate()` deve retornar erro se dados inválidos.

**Entrega:** `contact.go` com struct e métodos.

---

### Dia 6: Interfaces e Composição
- Defina interfaces: contratos que tipos devem seguir
- Implemente interfaces implicitamente (sem palavra-chave `implements`)
- Entenda polimorfismo via interfaces
- **Projeto:** Crie um sistema de logging com múltiplas saídas:
  - Defina uma interface `Logger` com método `Log(msg string) error`
  - Implemente dois tipos: `ConsoleLogger` (imprime no stdout) e `FileLogger` (escreve em arquivo)
  - Use interfaces pra fazer código agnóstico de onde os logs vão

**Desafio:** `FileLogger` deve criar o arquivo se não existir. Adicione níveis de log (Info, Warn, Error).

**Entrega:** `logger.go` com interface e implementações.

---

### Dia 7: Testes Unitários + Review
- Aprenda pacote `testing` nativo de Go
- Conveção: testes em arquivos `_test.go`
- Escreva funções `TestNomeDoTeste(t *testing.T)`
- Use `t.Errorf()` pra falhar testes
- **Projeto:** Escreva testes pra tudo que fez nos dias 1-6:
  - Teste a calculadora com casos normais e edge cases
  - Teste o jogo de adivinhação
  - Teste validação de contatos
  - Teste métodos da struct Contact
  - Teste ambos os loggers

**Desafio:** Cobertura mínima de 70% dos seus dias anteriores.

**Entrega do Dia 7:**
- Pasta `tests/` com todos os testes
- Arquivo `SEMANA_1_REVIEW.md` explicando o que aprendeu
- Commit no GitHub com tag `semana-1-completa`

---

## 🔧 Semana 2: HTTP, Routing e Handlers
**Meta:** Construir um servidor HTTP que processa requisições reais.

### Dia 8: HTTP Server Básico
- Aprenda pacote `net/http`
- Entenda handlers: funções que respondem a requisições HTTP
- Registre handlers em rotas específicas
- Inicie um servidor escutando em uma porta
- Teste com `curl` do terminal
- **Projeto:** Crie um servidor HTTP simples que:
  - Responda em `/` com uma mensagem "Bem-vindo"
  - Responda em `/greet/:name` com "Olá, {name}"
  - Responda em `/api/status` com um JSON `{"status": "ok"}`

**Desafio:** Extraia parâmetros da URL corretamente. Trate rotas inválidas (404).

**Entrega:** `server.go` rodando em `localhost:8080`.

---

### Dia 9: Roteamento e Métodos HTTP
- Diferencie `GET`, `POST`, `PUT`, `DELETE`
- Valide método HTTP dentro do handler
- Entenda diferença entre path e query params
- **Projeto:** Refatore o servidor anterior pra uma API de tarefas (TODO):
  - `GET /todos` → retorna lista de todas as tarefas
  - `POST /todos` → cria nova tarefa
  - `GET /todos/:id` → retorna tarefa específica
  - `DELETE /todos/:id` → deleta tarefa
  - Armazene TODOs em memória (slice de structs)

**Desafio:** Valide entrada no POST. Retorne erro 404 se todo não existe.

**Entrega:** `todo_api.go` com 4+ endpoints.

---

### Dia 10: JSON e Parsing
- Entenda serialização/desserialização de JSON
- Use struct tags pra mapear campos JSON
- Parse JSON de requisições (`json.Decoder`)
- Retorne respostas em JSON (`json.Encoder`)
- **Projeto:** Refatore o TODO:
  - Parse bodies JSON em `POST /todos`
  - Retorne todas as respostas como JSON válido
  - Use struct tags pra customizar nomes de campos JSON

**Desafio:** Valide que JSON é válido antes de processar. Retorne erros em JSON estruturado.

**Entrega:** API respondendo e aceitando JSON real.

---

### Dia 11: Middleware e Error Handling
- Entenda padrão de middleware em Go
- Crie middleware pra logging de requisições
- Estruture respostas de erro consistentemente
- **Projeto:** Adicione à API TODO:
  - Middleware que loga cada requisição (método, path, timestamp)
  - Função auxiliar pra retornar erros em JSON estruturado
  - Valide dados de entrada (TODO não pode ter nome vazio)

**Desafio:** Middleware deve ser reutilizável. Trate diferentes tipos de erro (validação, não encontrado, etc).

**Entrega:** API com logging e erros estruturados.

---

### Dia 12: Goroutines e Channels (Introdução)
- Entenda goroutines: lightweight threads gerenciadas por Go
- Crie goroutines com `go func()`
- Use channels pra comunicação entre goroutines
- **Projeto:** Simule processamento pesado:
  - Adicione um endpoint que simula trabalho longo (sleep 2-5 segundos)
  - Não bloqueie a resposta: processe em background com goroutine
  - Use channel pra notificar quando está pronto

**Desafio:** Servidor continua aceitando requisições durante processamento longo.

**Entrega:** Servidor processando requisições concorrentemente.

---

### Dia 13: Banco de Dados — SQLite Básico
- Aprenda pacote `database/sql`
- Instale driver SQLite: `github.com/mattn/go-sqlite3`
- Abra conexão, execute queries (Query, Exec, QueryRow)
- Implemente CRUD básico
- **Projeto:** Migre TODO de memória para SQLite:
  - Crie tabela `todos` com campos: id, name, description, done, created_at
  - Reescreva handlers pra ler/escrever no banco
  - Implemente CREATE, READ, UPDATE, DELETE

**Desafio:** Trate erros de banco (conexão falha, query inválida). Crie índices nas colunas apropriadas.

**Entrega:** API persistindo dados em banco SQLite.

---

### Dia 14: Estrutura de Projeto + Review
- Organize código em pacotes Go: `handlers/`, `models/`, `database/`, `main.go`
- Cada pacote tem responsabilidade clara
- Crie `go.mod` e mantenha dependências organizadas
- **Projeto:** Refatore TODO API em estrutura profissional:
  - `handlers/` → funções de handler
  - `models/` → structs de dados
  - `database/` → lógica de banco
  - `main.go` → setup do servidor

**Desafio:** Código deve estar pronto pra crescer sem bagunça.

**Entrega do Dia 14:**
- Código organizado em pacotes
- Testes pra handlers principais
- `SEMANA_2_REVIEW.md` documentando arquitetura
- README com instruções de rodagem
- Commit com tag `semana-2-completa`

---

## 📦 Semana 3-4: Backend Completo e Produção
**Meta:** Construir backend production-ready.

### Dia 15: Autenticação JWT Básica
- Entenda JWT: token baseado em assinatura digital
- Instale lib de JWT: `github.com/golang-jwt/jwt/v5`
- Crie claims customizados
- Gere tokens no login
- Valide tokens em endpoints protegidos
- **Projeto:** Adicione autenticação ao TODO:
  - Crie endpoint `POST /login` que recebe email/senha
  - Gere JWT token com expiração de 24h
  - Crie middleware `RequireAuth` que valida token
  - Proteja endpoints com middleware (ex: só usuário autenticado pode criar TODO)

**Desafio:** Token inválido ou expirado retorna 401. Inclua userID no token.

**Entrega:** Login + autenticação funcionando.

---

### Dia 16: Validação de Entrada
- Instale lib de validação: `github.com/go-playground/validator/v10`
- Defina struct tags pra regras de validação
- Valide structs antes de usar
- **Projeto:** Valide todos os inputs da API:
  - TODO.Name obrigatório, min 3 caracteres
  - Email válido no login
  - Password min 6 caracteres
  - Retorne erros de validação claros

**Desafio:** Mensagens de erro descrevem qual campo falhou e por quê.

**Entrega:** API rejeitando dados inválidos com clareza.

---

### Dia 17: Logging Profissional
- Use `log/slog` (Go 1.21+) ou lib alternativa
- Implemente logs estruturados com níveis: Debug, Info, Warn, Error
- Inclua contexto relevante (IDs, timestamps)
- **Projeto:** Substitua todos os logs da app:
  - Log cada requisição HTTP
  - Log erros de banco
  - Log autenticação/autorização
  - Estruture dados nos logs

**Desafio:** Logs incluem trace ID pra rastrear request end-to-end.

**Entrega:** App com logging estruturado.

---

### Dia 18: Testes de Integração
- Teste fluxos que envolvem HTTP + banco
- Use `httptest.NewServer()` pra servidor de teste
- Crie helpers pra setup/teardown de banco de teste
- **Projeto:** Escreva testes de integração:
  - Teste criar TODO (POST)
  - Teste buscar TODO (GET)
  - Teste atualizar TODO (PUT/PATCH)
  - Teste deletar TODO (DELETE)
  - Teste autenticação (token inválido retorna 401)

**Desafio:** Cada teste é independente (não depende de outro). Banco de teste é limpo entre testes.

**Entrega:** Suite de testes de integração cobrindo fluxos críticos.

---

### Dia 19: Variáveis de Ambiente e Config
- Use pacote `os` pra ler variáveis de ambiente
- Crie arquivo `.env` com configurações
- Construa struct `Config` que carrega todas as settings
- **Projeto:** Parametrize sua app:
  - PORT (porta do servidor)
  - DB_PATH (caminho do SQLite)
  - JWT_SECRET (chave de assinatura)
  - LOG_LEVEL (nível de logging)

**Desafio:** App roda sem `.env` com valores padrão sensatos.

**Entrega:** Config parametrizado, `.env` no `.gitignore`.

---

### Dia 20: Deployment Local
- Compile Go pra binário: `go build -o app`
- Teste binário funcionando localmente
- Entenda diferença entre `go run` vs `go build`
- **Projeto:** Compileizado sua TODO API:
  - Build sem erros
  - Teste binário executável
  - Documente como rodar em produção

**Desafio:** Binário é pequeno e rápido (Go compila pra binary nativo).

**Entrega:** Binário executável funcional.

---

### Dia 21: Extensão — Upload de Arquivo (Bônus)
- Entenda `multipart/form-data`
- Processe uploads: `r.FormFile()`
- Salve arquivos em disco
- **Projeto:** Permita anexar arquivo a TODO:
  - `POST /todos/:id/attachment` aceita file upload
  - Salve em `/uploads/`
  - Retorne URL do arquivo

**Desafio:** Valide tipo de arquivo. Limite tamanho.

**Entrega:** Upload funcionando.

---

### Dia 22: Paginação e Filtros
- Extraia query params: `r.URL.Query().Get()`
- Implemente LIMIT + OFFSET em SQL
- Retorne metadados (total, página atual, etc)
- **Projeto:** Adicione paginação ao GET /todos:
  - Query params: `?page=1&limit=10`
  - Query params: `?done=true` pra filtrar por status
  - Resposta inclui: dados, total de items, página atual

**Desafio:** Valores padrão sensatos (ex: limit máximo = 100).

**Entrega:** API paginada e filtrável.

---

### Dia 23: Tratamento de Transações
- Entenda conceito ACID
- Use `BeginTx()` pra iniciar transação
- Commit/Rollback automático
- **Projeto:** Crie operação atômica:
  - `POST /lists` cria "lista de afazeres" que contém múltiplos TODOs
  - Tudo deve ser inserido atomicamente (tudo ou nada)
  - Se falhar no meio, rollback automático

**Desafio:** Transação inclui timeout e context.

**Entrega:** Operações atômicas funcionando.

---

### Dia 24: Cache Básico
- Implemente cache simples em memória
- Use `sync.Map` pra thread-safety
- Invalide cache quando dados mudam
- **Projeto:** Adicione cache ao TODO:
  - Cache resultado de `GET /todos` por 30 segundos
  - Invalide cache quando criar/atualizar TODO
  - Monitore hit rate (opcional)

**Desafio:** Cache pode ser desativado via config.

**Entrega:** Cache funcionando, reduzindo queries.

---

### Dia 25: Documentação com Swagger
- Instale `github.com/swaggo/swag`
- Adicione anotações nos handlers pra gerar OpenAPI
- Gere documentação automática
- **Projeto:** Documente toda a API:
  - Cada endpoint tem comentários descrevendo: resumo, parâmetros, respostas
  - Swagger UI em `/swagger`
  - Docs podem ser gerados automaticamente

**Desafio:** Documentação está sempre sincronizada com código.

**Entrega:** Swagger UI rodando com docs completos.

---

### Dia 26: Docker e Containerização
- Crie `Dockerfile` pra sua app
- Multi-stage build (compile em um stage, rode em outro)
- Crie `docker-compose.yml` com app + SQLite
- **Projeto:** Containerize sua TODO API:
  - Build image Docker
  - App rodando em container
  - Docker-compose pra facilitar setup

**Desafio:** Image é pequeno (< 50MB). App em container fala com banco.

**Entrega:** App rodando em Docker.

---

### Dia 27: CI/CD Básico (GitHub Actions)
- Crie workflow `.github/workflows/test.yml`
- Roda testes automaticamente em cada push
- **Projeto:** Configure GitHub Actions:
  - Instale dependências
  - Roda `go test ./...`
  - Build e valida binário
  - Reporta cobertura de testes

**Desafio:** Testes rodam automaticamente. Build falha se testes falham.

**Entrega:** CI/CD rodando no GitHub.

---

### Dia 28: Tratamento de Erros Avançado
- Crie tipos de erro customizados
- Use `errors.Is()`, `errors.As()`
- Wrapping: `fmt.Errorf("context: %w", err)`
- **Projeto:** Estruture erros da app:
  - `NotFoundError`, `ValidationError`, `UnauthorizedError`
  - Handlers retornam tipo de erro apropriado
  - Cliente recebe código HTTP e mensagem estruturada

**Desafio:** Stack traces úteis. Erros descritivos.

**Entrega:** Sistema de erro robusto.

---

### Dia 29: Rate Limiting
- Instale `golang.org/x/time/rate`
- Implemente rate limiter por IP
- **Projeto:** Proteja sua API:
  - Limite 100 requisições por minuto por IP
  - Retorne 429 se limite excedido
  - Retorne headers informando limit/remaining

**Desafio:** Rate limit não afeta autenticação (login sempre allowed).

**Entrega:** API com proteção contra abuse.

---

### Dia 30: Projeto Final + Deploy
**Meta:** Publicar seu backend Go online.

- Escolha plataforma: Railway, Render, DigitalOcean, ou VPS
- Deploy TODO API com banco
- Configure domínio (opcional)
- Teste endpoints públicos com `curl`
- Documente processo de deploy

**Entrega Final do Dia 30:**
- Backend online e funcional
- GitHub repo com estrutura completa
- README passo-a-passo (setup, testes, deploy)
- `PROJETO_FINAL.md` explicando:
  - Arquitetura (pacotes, responsabilidades)
  - Decisões técnicas
  - Trade-offs

---

## 📊 Checklist de Completion

Ao final dos 30 dias, você deve ter:

- ✅ Sintaxe Go fluente (tipos, funções, structs, interfaces)
- ✅ HTTP server robusto com roteamento correto
- ✅ Banco de dados (SQLite com CRUD)
- ✅ Autenticação com JWT
- ✅ Validação de entrada
- ✅ Testes unitários + integração (cobertura > 70%)
- ✅ Logging estruturado
- ✅ Documentação (Swagger)
- ✅ Docker
- ✅ CI/CD (GitHub Actions)
- ✅ Deployed em produção (online)

**Resultado:** Um backend Go production-ready que você construiu do zero.

---

## 🛠️ Setup Inicial (Antes do Dia 1)

- Instale Go 1.21+
- Crie pasta workspace: `mkdir 30-days-of-go && cd 30-days-of-go`
- Inicie repo Git: `git init`
- Inicie módulo Go: `go mod init github.com/seu-user/30-days-of-go`
- Crie arquivo `main.go` vazio
- Primeiro commit: "Initial commit"

Conforme avança, você vai adicionar dependências com `go get`.

---

## 💡 Dicas Finais

1. **Comita diariamente** — Isso força disciplina e cria histórico
2. **Teste constantemente** — `go run` frequente, `go test` antes de commit
3. **Leia Go idiomático** — effective-go.golang.org é essencial
4. **Refatore sem medo** — Go é fácil refatorar, compile rápido
5. **Comunidade** — Go Reddit, Go Discord, Gophers Slack

---

**Você consegue. Vamo lá.**

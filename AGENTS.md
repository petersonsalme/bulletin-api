# AGENTS.md — bulletin-api

Diretrizes para agentes de IA trabalhando no repositório `bulletin-api`.

## Contexto do Projeto

API em Go para publicação e listagem de avisos/boletins (Bulletins), com persistência em PostgreSQL e deploy em Docker/Kubernetes.

- Para o esquema de banco de dados, documentação de endpoints e exemplos de chamadas, consulte o [README.md](README.md).

## Estrutura do Repositório

```
bulletin-api/
├── api/                # Ponto de entrada da aplicação (main.go)
├── pkg/                # Pacotes Go (handlers, modelos, conexão com banco)
├── migrations/         # Scripts SQL de migração de banco de dados
├── docker/             # Dockerfiles e scripts de containerização
├── docker-compose.yml  # Orquestração local de serviços (API + PostgreSQL)
├── k8s/                # Manifestos Kubernetes
├── requests.http       # Exemplos de requisições HTTP para testes manuais
├── Makefile            # Comandos de build e execução
└── README.md           # Documentação da API e instruções de setup
```

## Regras e Comandos

- Prefix shell commands with `rtk` (see `.agents/rules/rtk.md`).

1. **Build e Testes:**
   - Baixar dependências: `go mod download`
   - Compilar binário: `go build -o ./bin/bulletin-api api/main.go` (ou `make build-api`)
   - Rodar testes: `go test ./...`
   - Formatação: `gofmt -s -w .`
2. **Ambiente Local:**
   - Iniciar serviços via Docker: `docker-compose up -d` (ou `make run`)
3. **Padrões de Desenvolvimento:**
   - Manter consistência com as migrações SQL em `migrations/` ao alterar modelos.
   - Commits seguindo Conventional Commits (`feat:`, `fix:`, `refactor:`, `test:`).
   - Comunicação com o usuário em Português do Brasil; código e commits em Inglês.

@.agents/rules/rtk.md

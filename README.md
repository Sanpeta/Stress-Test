# Desafio-Tecnico-Stress-Test

Este projeto fornece uma ferramenta de linha de comando (CLI) para realizar testes de carga e estresse em um serviço web. Utiliza Go e Cobra CLI para configurar e executar os testes, além de gerar relatórios com os resultados.

## Funcionalidades

- **Teste de Carga**: Realiza um número fixo de requisições simultâneas e distribui as requisições de acordo com o nível de concorrência.
- **Teste de Estresse**: Aumenta gradualmente a carga até atingir a concorrência máxima definida e gera um relatório de desempenho.
- **Relatório Detalhado**: Inclui tempo total gasto, quantidade de requisições realizadas, e distribuição de códigos de status HTTP.

## Requisitos

- Docker (para executar a aplicação no contêiner)

## Construindo a Imagem Docker

Para construir a imagem Docker para o projeto, execute o seguinte comando no diretório raiz do projeto:

```bash
docker build -t loadtester .
```

## Executando o Teste de Carga

Para executar um teste de carga, use o comando abaixo:

```bash
docker run loadtester loadtest --url=http://example.com --requests=1000 --concurrency=10
```

**Parâmetros:**

- `--url`: URL do serviço a ser testado (obrigatório)
- `--requests`: Número total de requisições (obrigatório)
- `--concurrency`: Número de chamadas simultâneas (obrigatório)

## Executando o Teste de Estresse

Para executar um teste de estresse, use o comando abaixo:

```bash
docker run loadtester stresstest --url=http://example.com --initial-concurrency=10 --max-concurrency=100 --increment=10
```

**Parâmetros:**

- `--url`: URL do serviço a ser testado (obrigatório)
- `--initial-concurrency`: Número inicial de chamadas simultâneas
- `--max-concurrency`: Número máximo de chamadas simultâneas
- `--increment`: Incremento na concorrência a cada passo

## Relatório

Após a execução, um relatório será exibido no terminal com informações sobre o tempo total, quantidade de requisições, e a distribuição de códigos de status HTTP.

## Exemplo de Uso com Docker

Para construir e executar o contêiner Docker com um teste de carga, você pode usar:

```bash
docker build -t loadtester .
docker run loadtester loadtest --url=http://example.com --requests=1000 --concurrency=10
```

Para um teste de estresse:

```bash
docker run loadtester stresstest --url=http://example.com --initial-concurrency=10 --max-concurrency=100 --increment=10
```

## Contribuindo

Sinta-se à vontade para contribuir com o projeto. Envie suas sugestões, correções ou melhorias através de pull requests.

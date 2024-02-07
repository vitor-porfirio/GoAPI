# API em Golang - README

## Visão Geral

Este projeto fornece uma estrutura simples para uma API em Golang, usando Docker e MySQL.

## Pré-requisitos

Certifique-se de ter o Docker instalado em sua máquina antes de prosseguir.

## Configuração e Execução

1. **Clone este repositório:**
    ```bash
    git clone https://github.com/vitor-porfirio/GoAPI.git
    cd GoAPI
    ```

2. **Inicie os contêineres Docker usando o Docker Compose:**
    ```bash
    docker-compose up -d
    ```

3. **Acesse o bash do contêiner MySQL:**
    ```bash
    docker-compose exec mysql bash
    ```

4. **No bash do MySQL, execute o seguinte comando para acessar o banco de dados:**
    ```bash
    mysql -uroot -p imersao17
    ```
    - **Senha:** `3711521`

    Agora você está dentro do banco de dados.

5. **Execute o script para gerar as tabelas. O script está localizado no arquivo `db.sql`.**

## Testando a API

Para testar a API, utilize o arquivo `test.http`. Este arquivo contém requisições HTTP pré-configuradas que você pode usar diretamente no seu ambiente de desenvolvimento.

Certifique-se de ter uma extensão ou ferramenta que suporte a execução de arquivos `.http`. Caso esteja usando o Visual Studio Code, a extensão "REST Client" é uma opção popular.

Execute as requisições no arquivo `test.http` para verificar se a API está funcionando corretamente.

## Contribuindo

Se você encontrar problemas ou tiver sugestões de melhoria, sinta-se à vontade para abrir uma issue ou enviar um pull request. 🚀

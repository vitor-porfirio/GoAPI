<!DOCTYPE html>
<html>

<head>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            color: #333;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
        }

        h1,
        h2,
        h3 {
            color: #008080;
        }

        code {
            background-color: #f4f4f4;
            border: 1px solid #ddd;
            border-radius: 4px;
            font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
            padding: 2px 6px;
        }

        pre code {
            display: block;
            padding: 10px;
            overflow-x: auto;
            max-width: 100%;
        }

        pre {
            background-color: #f8f8f8;
            border: 1px solid #ddd;
            border-radius: 4px;
            font-size: 14px;
            line-height: 1.4;
            overflow: auto;
            padding: 10px;
        }

        pre, code {
            margin-bottom: 20px;
        }

        ul {
            list-style: disc;
            padding-left: 20px;
        }

        li {
            margin-bottom: 5px;
        }
    </style>
</head>

<body>

    <h1>API em Golang - README</h1>

    <h2>Visão Geral</h2>

    <p>Bem-vindo à API minimalista em Golang! Este projeto fornece uma estrutura simples para uma API usando Docker e MySQL.</p>

    <h2>Pré-requisitos</h2>

    <p>Certifique-se de ter o Docker instalado em sua máquina antes de prosseguir.</p>

    <h2>Configuração e Execução</h2>

    <ol>
        <li>
            <strong>Clone este repositório:</strong>
            <pre><code>git clone https://github.com/seu-usuario/sua-api-golang.git
cd sua-api-golang</code></pre>
        </li>
        <li>
            <strong>Inicie os contêineres Docker usando o Docker Compose:</strong>
            <pre><code>docker-compose up -d</code></pre>
        </li>
        <li>
            <strong>Acesse o bash do contêiner MySQL:</strong>
            <pre><code>docker-compose exec mysql bash</code></pre>
        </li>
        <li>
            <strong>No bash do MySQL, execute o seguinte comando para acessar o banco de dados:</strong>
            <pre><code>mysql -uroot -p imersao17</code></pre>
            <ul>
                <li><strong>Senha:</strong> <code>3711521</code></li>
            </ul>
        </li>
    </ol>

    <h2>Testando a API</h2>

    <p>Para testar a API, utilize o arquivo <code>test.http</code>. Este arquivo contém requisições HTTP pré-configuradas que você pode usar diretamente no seu ambiente de desenvolvimento.</p>

    <p>Certifique-se de ter uma extensão ou ferramenta que suporte a execução de arquivos <code>.http</code>. Caso esteja usando o Visual Studio Code, a extensão "REST Client" é uma opção popular.</p>

    <p>Execute as requisições no arquivo <code>test.http</code> para verificar se a API está funcionando corretamente.</p>

    <h2>Contribuindo</h2>

    <p>Se você encontrar problemas ou tiver sugestões de melhoria, sinta-se à vontade para abrir uma issue ou enviar um pull request.</p>

    <p>Divirta-se codificando! 🚀</p>

</body>

</html>

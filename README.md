<h1>API REST em Go</h1>

<p>Este é um exemplo de implementação de uma API REST em Go. A API oferece endpoints para realizar operações CRUD em recursos específicos.</p>

<h2>Configuração do Ambiente</h2>

<p>Antes de começar a usar a API, é necessário configurar o ambiente de desenvolvimento:</p>

<ol>
    <li><strong>Instalação do Go</strong>: Certifique-se de ter o Go instalado. Você pode encontrar instruções de instalação em <a href="https://golang.org/doc/install">golang.org</a>.</li>
    <li><strong>Instalação das Dependências</strong></li>
</ol>

<pre><code>go mod tidy</code></pre>

<p>3. <strong>Execução da API</strong>: Para iniciar a API, execute o seguinte comando:</p>

<pre><code>go run main.go</code></pre>

<p>Isso iniciará o servidor da API na porta padrão <code>8080</code>.</p>

<h2>Endpoints</h2>

<p>A API oferece os seguintes endpoints:</p>

<ul>
    <li><strong>GET /api/v1/users</strong>: Retorna uma lista de todos os usuários.</li>
    <li><strong>GET /api/v1/users/{id}</strong>: Retorna os detalhes de um usuário específico.</li>
    <li><strong>POST /api/v1/users</strong>: Cria um novo usuário.</li>
    <li><strong>PUT /api/v1/users/{id}</strong>: Atualiza os detalhes de um usuário existente.</li>
    <li><strong>DELETE /api/v1/users/{id}</strong>: Exclui um usuário existente.</li>
</ul>

<h2>Estrutura do Projeto</h2>

<ul>
    <li><strong>main.go</strong>: O ponto de entrada da aplicação.</li>
    <li><strong>handlers</strong>: Contém os manipuladores (handlers) para cada rota da API.</li>
    <li><strong>models</strong>: Define os modelos de dados da aplicação.</li>
    <li><strong>router</strong>: Configuração do roteador da API.</li>
</ul>

<h2>Contribuição</h2>

<p>Sinta-se à vontade para contribuir com melhorias para este projeto! Você pode abrir um pull request ou reportar problemas através das issues.</p>


</body>
</html>
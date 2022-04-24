

## Exemplo Utilizando GraphQL 

   
    

## Configuração do ambiente de desenvolvimento

    - Utilizado Visual Studio Code 
        - Plugin Remote Container (https://code.visualstudio.com/docs/remote/containers) 
    
    - go version go1.17.8 linux/amd64
    - Depencias do go está no arquivo go.mod
  
## Comandos

   
## Historico de comandos

    - go mod init github.com/rafaelyokota/fc-graphql
    - go get github.com/99designs/gqlgen
    - go run github.com/99designs/gqlgen init //cria modelo de exemplo
    - alterar  schema.graphql 
    - no arquivo graph/schema.resolvers.go [https://www.howtographql.com/graphql-go/1-getting-started/]
        no final do arquivo deletar ou comentar linhas que conter os "metodos":
            - CreateTodo
            - Todos

    - gqlgen generate


## Notas
    - Arquivo de resolver.go 
        - EX: resolver.go para categoria requisitada pelo client, caso precise buscar dados para complementar a categoria, em uma outra api, servico, banco, depois retornar ao que cliente requisitou (insert ou update ou delete)
        - resolver.go é o cara que faz injeção de depencia
        - A regra de negocio pode ficar nesse arquivo

    - No arquivo de 'schema.resolvers.go'
        - é como se fosse "controller" de uma API, não tem "regra de negocio" nesse arquivo
        - utilizar resolver.go (injeção de dependencias)
        - é o arquivo onde você implementa o "select"(ou utilizando repostory)  por exemplo
        - tem types de "mutation" e "query" que faze gerenciamento do Resolver
    -  A separação da models, e criar o relacionamento "manualmente", somente vai existir o relacionamento quando necessario(sob demanda!), faz os relacionamento entre as models conforme precisa, se o cliente requisitar somente dados de uma model não tem porque deixar estar relacionado sempre com a outra model.
    - TOMAR MUITO CUIDADO relacionamentos N+1, para cada registro vai de um curso por exemplo vai relacionar toda vez que tiver outros cursos, ou em outras palavras para cada curso vai ter seu relacionamento, vai gerar efeito cascata(meio que loop infinito), degradando a performance.
        - Quando for implementar crie debugs, logs muito bom, pra poder auditar, e não entrar no problema no N+1
        - Leia sobre https://gqlgen.com/reference/dataloaders/
    - DataLoaders tenta otimizar essas querys N+1, cacheando alguns dados de relacionamento, otimizando algoritmo nessas situações
    - Algo que vai ajudar é criar politicas para perfomance, limites de relacionamento, tempos de respotas...
    - Pesquisar sobre boas praticas,  para alinhar com a performance!
# Referencias 

    - https://code.visualstudio.com/docs/remote/create-dev-container
    - https://gqlgen.com/
    - https://github.com/99designs/gqlgen



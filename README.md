```shell
go run main.go
```

Gerar o executável:

```shell
go build main.go
```

Executável para windows:

```shell
GOOS=windows go build main.go
```

Começar o projeto: 

```shell
go mod init github.com/xandreafonso/gogo
```

Executar todos os testes: 

```shell
go test ./...
```

Baixar um pacote:

```shell
go get github.com/stretchr/testify/assert
```

Ou que podemos usar depois de já ter usado a importação no arquivo:

```shell
go mod tidy
```

Para baixar o módulo do sqlite3:

```shell
apt-get install gcc
CGO_ENABLED=1 go get github.com/mattn/go-sqlite3
```

Criar tabela no sqlite3. Primeiro instalar:

```shell
apt install sqlite3
```

Depois cria o banco ao mesmo tempo que abre o prompt para executar SQL no mesmo:

```shell
sqlite3 db.sqlite3
```

Vai abri o prompt do sqlite3 e é só criar a tabela:

```sql
create table ordr (id varchar(255) not null, price float not null, tax float not null, final_price float not null, primary key (id));
```

Gerar o executável do projeto:

```shell
go build
```

Gerar o container Docker do projeto:

```shell
docker build -t xandreafonso/gogo:latest .
```

Analisar se uma imagem tem vulnerabilidades:

```shell
docker scout quickview
```

Rodar o container docker:

```shell
docker run -p 8888:8888 xandreafonso/gogo:latest
```

Subir imagem do docker para o docker hub:

```shell
docker push xandreafonso/gogo:latest
```

TODO. Ainda não consegui gerar a imagem Docker. Na hora de baixar dependências, dá erro. Erros:

```
#0 1055.6       golang.org/x/text/secure/bidirule: golang.org/x/text@v0.13.0: read "https:/proxy.golang.org/@v/v0.13.0.zip": stream error: stream ID 105; INTERNAL_ERROR; received from peer

#0 1055.6       golang.org/x/text/unicode/bidi: golang.org/x/text@v0.13.0: read "https:/proxy.golang.org/@v/v0.13.0.zip": stream error: stream ID 105; INTERNAL_ERROR; received from peer

#0 1055.6       golang.org/x/text/unicode/norm: golang.org/x/text@v0.13.0: read "https:/proxy.golang.org/@v/v0.13.0.zip": stream error: stream ID 105; INTERNAL_ERROR; received from peer

Dockerfile:4
--------------------
   2 |     WORKDIR /app
   3 |     COPY . .
   4 | >>> RUN go mod tidy
   5 |     RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd/api/main.go
   6 |     CMD ["./api"]
--------------------
ERROR: failed to solve: process "/bin/sh -c go mod tidy" did not complete successfully: exit code: 1
```
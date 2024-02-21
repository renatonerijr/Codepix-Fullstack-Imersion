# CodePIX

Repositório feito com base nas aulas da Imersão Full Stack && Full Cycle, codepix é um repositório em GO que utiliza a arquitetura hexagonal e comunica via GRPC.

# Como rodar

Para fazer a build da API GO
```
docker compose up -d --build
```

Para rodar o servidor GRPC
```
docker compose exec app bash
```

Após isso rodar com
```
go run grpc --port 5501
```

Em outra aba é possivel testar o server grpc com
```
evans -r repl
```
# poc-registry

Exemplo de implementação de um service registry + proxy.

## registry

* Contém uma base de dados para registrar os serviços remotos.
* Contém um proxy para redirecionar as chamadas para os serviços concretos (first-process e second-process)

## first-process

* Simula um processo qualquer.
* No start-up, esse serviço se registra em "registry".

## second-process

* Simula um outro processo qualquer.
* No start-up, esse serviço se registra em "registry".

## Execução

1) ```
    cd registry
    go run cmd/main.go
   ```
2) ```
    cd first-process
    go run cmd/main.go
   ```
3) ```
    cd second-process
    go run cmd/main.go
   ```
______________

As chamadas devem ser direcionadas a "registry". Idealmente first-process e second-process ficam "escondidas".

```
curl --location --request POST 'http://localhost:8080/commands' \
--header 'x-process-name: first-process' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "123",
    "process": "xpto"
}'
```

```
curl --location --request POST 'http://localhost:8080/commands' \
--header 'x-process-name: second-process' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "321",
    "process": "xpto2"
}'
```

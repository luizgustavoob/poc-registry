# poc-registry

Exemplo de implementação de um service registry + proxy.

## registry

* Contém uma base de dados para registrar os serviços remotos.
* Contém um proxy para redirecionar as chamadas para os serviços concretos (hu-assembly e shipment-injection)

## hu-assembly

* Backlog para manipular o processo "hu-assembly".
* No start-up, esse serviço se registra em "registry".

## shipment-injection

* Backlog para manipular o processo "shipment-injection".
* No start-up, esse serviço se registra em "registry".

## Execução

1) ```
    cd registry
    go run cmd/main.go
   ```
2) ```
    cd hu-assembly
    go run cmd/main.go
   ```
3) ```
    cd shipment-injection
    go run cmd/main.go
   ```
______________

As chamadas devem ser direcionadas a "registry". Idealmente hu-assembly e shipment-injection ficam "escondidas".

```
curl --location --request POST 'http://localhost:8080/commands' \
--header 'x-process-name: shipment-injection' \
--header 'Content-Type: application/json' \
--data-raw '{
    "type": "create",
    "target": {
        "type": "target",
        "id": "id"
    }
}'
```

```
curl --location --request POST 'http://localhost:8080/commands' \
--header 'x-process-name: hu-assembly' \
--header 'Content-Type: application/json' \
--data-raw '{
    "type": "create",
    "target": {
        "type": "target",
        "id": "id"
    }
}'
```

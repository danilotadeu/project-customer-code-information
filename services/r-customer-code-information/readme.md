# r-customer-code-information #

- Serviço construído para consumir uma api do serviço r-customer-code-information-provider.

## Como rodar a aplicação ##

- Para rodar a aplicação localmente, utilizar somente o comando "make run-watch", 
o serviço estará disponível na porta :3000

## CURL para efetuar a chamada no serviço ##

```bash
$ curl --location --request POST 'http://127.0.0.1:3000/v1/r-customer-code-information' \
--header 'clientId: 123' \
--header 'messageId: 123456' \
--header 'Authorization: Bearer omcauckqoiy6jyzgsxu6gi7sxh' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userId": "F8036589",
    "customer": {
        "id": "116084470",
        "syncFlag": false
    }
}'
```

Response:
```
{
    "customerCode": "8882828282"
}
```
## Tests ##

- Para rodar os testes, rodar o comando abaixo:

```bash
go test -timeout 30s -run ^Test_apiImpl_CodeInformationHandler$ github.com/danilotadeu/r-customer-code-information/api/codeinformation
```

## Provider ## 

- Para rodar o mock do provider, acessar o path:

```
$ cd /mocks/provider
```

e rodar o comando:

```
$ make run-watch
```

o mock provider estará disponível na porta :4000.
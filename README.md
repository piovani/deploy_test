# Deploy Test

## Gerar um Build

```
GOOS=linux GOARCH=amd64 go build nome-arquivo.go 
```

## Gerar um serviço

1. Criar o arquivo de serviço

```
touch /lib/systemd/system/goapp.service
```

2. Configurar o serviço
```
[Unit]
Description=goapp

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/home/golang/nome-arquivo.go

[Install]
WantedBy=multi-user.target
```

3. Iniciar o serviço
```
service goapp start
```

4. Abilitar o serviço na inicialização do sistema
```
service goapp enable

```

5. Status do servíço
```
service goapp status
```
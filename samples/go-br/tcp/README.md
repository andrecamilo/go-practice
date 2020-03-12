# tcp Trabalhando com o protocolo tcp

Go através do pacote net provê interfaces de acesso a I/O, incluindo a pilha TCP/IP, UDP, resolução por nome de domínio e UNIX Sockets.

Neste exemplo, iremos criar um cliente TCP/IP e um servidor TCP/IP.

Servidor.

O protocolo de comunicação que nosso servidor deverá trabalhar será este:
- 1) Ouvir a interface tcp na porta 8081
- 2) Aceitar conexões
- 3) Dentro de um loop infinito, ouvir as mensagens a serem transmitidas pelo cliente
- 4) Escrever no terminal estas mensagens
- 5) Devolver a mensagem recebida ao cliente (eco)

```go
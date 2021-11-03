# K2 Publisher-Subscriber

Your task is to create 2 services that will accept WebSocket
connection using `github.com/gorilla/websocket` package.

Both of them should expose:
* _GET_ `/healthz` with _hostname_, _app name_ and _current time_
* _WS_ `/ws` that will handle all WebSockets connections

First service will listen to all incoming messages and publish
them in a queue (you can use `github.com/wagslane/go-rabbitmq` or any other),
don't have to publish anything in response.

Second will listen to incoming messages from queue and publish them
to all connected WS clients.

There is simple HTML client provided in `tools/client/index.html` that by default is
trying to connect to `ws://localhost:8080/ws` and `ws://localhost:8090/ws`.

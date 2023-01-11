### Simple gRPC binary

This can be used to simulate traffic. The dynamic interval settings can simulate different network congestions in distributed systems.

---

### Config

|Property | Info
|-|-|
|Listenport | server listening port
|Interval | interval of client requests to the server
|Target | server IP/URL for the client to connect to
|Message | client message to send
|Answer | server message to respond
|Runtype | binary behaviour (client or server)

---

### Docker Image

There is a docker image of v1.0.2-alpha on https://hub.docker.com/r/hyrsh/simplegrpc

```
docker pull hyrsh/simplegrpc:1.0
```
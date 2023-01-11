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

### Usage

Run the binary without config and a default template gets created and the program exits so that you can adjust the template to your needs. If a valid config is found, it will just start.

The "-config" flag can specify what config to load (useful for systemd; not working in docker images --> there you have to overwrite the existing config in the default path).

Client startup:

```
./simplegrpc_linux_x86_64 -config myclientcfg.yml
```

Server startup:

```
./simplegrpc_linux_x86_64 -config myservercfg.yml
```

---

### Kubernetes

Deploy a client with:

```
kubectl apply -f https://raw.githubusercontent.com/hyrsh/simplegrpc/main/kubernetes/simplegrpc-client-deploy.yml
```

Deploy a server with:

```
kubectl apply -f https://raw.githubusercontent.com/hyrsh/simplegrpc/main/kubernetes/simplegrpc-server-deploy.yml
```

---

### Docker Image

There is a docker image of v1.0.2-alpha on https://hub.docker.com/r/hyrsh/simplegrpc

```
docker pull hyrsh/simplegrpc:1.0
```
### Simple gRPC binary

This is just a tool to send hello messages and receive server responses over gRPC.

You can adjust the client intervall in a typical time.Duration string (e.g. "1s", "1ms", "1m" ...)

### Docker Image

There is a docker image of v1.0.2-alpha on https://hub.docker.com/r/hyrsh/simplegrpc

```
docker pull hyrsh/simplegrpc:1.0
```
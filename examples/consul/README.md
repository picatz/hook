# Consul Dev Agent Configuration

This directory only contains the Consul agent dev server configuration found in `server.hcl`.

Refer the the `Makefile`s in the `examples/` directory to see the `consul` commands used.

## Example Usage

Start a development agent:

```console
$ consul agent -dev -config-dir="./consul/"
...
```

Register an example "web" service that is connect-enabled, but without custom WASM filters in this case:

```hcl
// service.hcl
service {
  name = "web"
  port = 8080
  connect { sidecar_service {} }
}
```

Register the "web" service:

```console
$ consul services register service.hcl
Registered service: web
```

Start a simple web service and run an Envoy sidecar for the service:

```console
$ python3 -m http.server 8080 &
$ consul connect envoy -sidecar-for web
...
```

Then in another terminal, use the development proxy to connect to the service through the service mesh on `localhost:9090` masquerading as an `example` service:

```console
$ consul connect proxy -service example -upstream web:9090
...
```

Then in another terminal, we can curl the proxy on `localhost:9090` which will tunnel through the mesh .All on `localhost` for developmenet purposes, of course. While you can `curl` the service directly on `localhost:8080`, we're ignoring that for this example, as this wouldn't be possible outside the single host. That's where service mesh gets more interesting, but outside the scope of this example.

```console
$ curl localhost:9090
...
```

To summarize:

* We started a `consul` dev agent.
* Registered a simple connect-enabled service.
* Ran a sidecar proxy using Envoy for the connect-enabled service.
* Ran a development proxy to test the connect-enable service.
* Tested the service connectivity worked through the mesh using `curl` using a local proxy endpoint on `localhost:9090`.

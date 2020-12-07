# 🏴‍☠️ Hook

Proxy [WASM](https://webassembly.org/) filter SDK with a bit of [fairy dust](https://youtu.be/aYNNwKluGpc?t=58) ✨🧚‍♀️

## Examples

To help you get started, here are a few examples to use as reference:

* [Header](./examples/header) is a filter that sets the `wasm: enabled` HTTP header on all requests/responses.
* [Security Headers](./examples/security_headers) is a filter that adds multiple security headers on all HTTP responses.
* [Replace Response](./examples/replace_response) is a filter that replaces the HTTP response body (and `content-length` header) for all HTTP responses using a custom `http.Context` object.
* [Sniff](./examples/sniff) is a filter that logs all HTTP request/response headers and bodies.

## Why

There is already a [similar SDK framework for Go](https://github.com/tetratelabs/proxy-wasm-go-sdk) maintained by @mathetake, and I totally suggest using that SDK. They are doing fantastic work across the Envoy, WASM, and TinyGo community. This is currently a personal project to experiment with different SDK patterns focused on breaking up the required pieces into logical packages, and enabling a potentially more idiomatic SDK. It's another option for the community.

I've also started playing with enabling WASM filters in [HashiCorp Consul](https://www.consul.io/) using expierimental changes on the [`wasm-filters`](https://github.com/hashicorp/consul/compare/master...wasm-filters) branch. Having a ton of fun learning how exactly this all works at a lower-level, and exposing clean APIs at a higher-level in Consul and Go.

The Consul service config to enable WASM filters for an Envoy sidecar proxy looks like this:

```hcl
protocol = "http"
wasm_filters = [
    {
        name       = "security_headers"
        local_file = "path/to/filters/security_headers.wasm"
    },
    {
        name       = "replace_response"
        local_file = "path/to/filters/replace_response.wasm"
    }
]
```

<details>
  <summary>ℹ️ Full configuration example (click to expand)</summary>

  Sample configuration using `wasm_filters` for Consul to configure Envoy sidecar
  proxies with the `/tmp/filters/replace_response.wasm` WASM filter, assuming this
  binary file is appropriatley compiled and is avaiable on the local filesystem of
  the sidecar host.

  ```hcl
  service {
    name = "web"
    port = 8080
    connect {
        sidecar_service {
            proxy {
                config {
                    protocol = "http"
                    wasm_filters = [
                        {
                            name       = "replace_response"
                            local_file = "/tmp/filters/replace_response.wasm"
                        }
                    ]
                }
            }
        }
  }
  ```

  👩🏽‍💻 Learn more about Consul [here](https://learn.hashicorp.com/tutorials/consul/service-mesh?in=consul/gs-consul-service-mesh).
</details>

## SDK Usage

Simple example to set the `wasm: enabled` HTTP header on all requests/responses through the proxy. This is obviously a trivial example in order to demonstrate the how to start "hooking" into the the proxy.

```golang
package main

import (
    "github.com/picatz/hook/pkg/call/http"
    "github.com/picatz/hook/pkg/types/action"
)

func main() {
    http.OnRequestHeaders(func(int, bool) action.Type {
        http.SetRequestHeader("wasm", "enabled")
        return action.Continue
    })

    http.OnResponseHeaders(func(int, bool) action.Type {
        http.SetResponseHeader("wasm", "enabled")
        return action.Continue
    })
}
```

<details>
  <summary>ℹ️ Example details (click to expand)</summary>

* `github.com/picatz/hook/pkg/call/http` provides functions to interact with HTTP requests and responses.
  * `http.OnRequestHeaders` is a function to hook into the proxy. From here you can inspect, set, delete, read, and further handle header-based authn/authz tasks.
    * You can optionally use the two arguments provided to the function including the `maxSize` (`int` type) and `endOfStream` (`bool` type). You do not _need_ to name these types at all if you do not plan to use them. That is a subtle, often unused, feature of the Go language.
    * To continue processing the request after your logic, return the `action.Continue` type to signal the HTTP stream is ready to be handled again by the proxy.
  * `http.OnResponseHeaders` is a function to hook into the the response headers, very similiar to `http.OnRequestHeaders`, but for the response side of the upstream service. You can do essentially the exact same things, but applying the logic to the "other side" of the proxy connection.
  * `http.SetRequestHeader` is a function to set an HTTP _request_ header using a given key and value.
  * `http.SetResponseHeader` is a function to set an HTTP _response_ header using a given key and value.
* `github.com/picatz/hook/pkg/types/action` provides types to signal the proxy to continue/stop processing the next steps of the request/response stream.
  * `action.Continue` is a common type used to signal to the proxy to continue handling the connection.

</details>

### Build with TinyGo

Compile using [`tinygo`](https://tinygo.org/getting-started/) to bulild `example.wasm`:

```console
$ tinygo build -o example.wasm -scheduler=none -target=wasi -wasm-abi=generic main.go
```

### Configure Envoy Proxy

Add the filter config to the [Envoy](https://github.com/envoyproxy/envoy) `http_filters` section using a version with WASM support:

```yaml
# envoy.yaml
name: envoy.filters.http.wasm
typed_config:
  "@type": type.googleapis.com/udpa.type.v1.TypedStruct
  type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
  value:
    config:
      name: "example"
      root_id: "example"
      vm_config:
        vm_id: "example"
        runtime: "envoy.wasm.runtime.v8"
        code:
          local:
            filename: "/full/path/to/example.wasm"
        allow_precompiled: true
```

<details>
  <summary>ℹ️ Full configuration example (click to expand)</summary>

  ```yaml
  static_resources:
  listeners:
    - name: main
      address:
        socket_address:
          address: 0.0.0.0
          port_value: 18000
      filter_chains:
        - filters:
            - name: envoy.http_connection_manager
              typed_config:
                "@type": type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager
                stat_prefix: ingress_http
                codec_type: auto
                route_config:
                  name: local_route
                  virtual_hosts:
                    - name: local_service
                      domains:
                        - "*"
                      routes:
                        - match:
                            prefix: "/"
                          route:
                            cluster: web_service
                http_filters:
                  - name: envoy.filters.http.wasm
                    typed_config:
                      "@type": type.googleapis.com/udpa.type.v1.TypedStruct
                      type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
                      value:
                        config:
                          name: "req_body_replace"
                          root_id: "req_body_replace"
                          vm_config:
                            vm_id: "req_body_replace"
                            runtime: "envoy.wasm.runtime.v8"
                            code:
                              local:
                                filename: "./examples/header/header.wasm"
                            allow_precompiled: true
                  - name: envoy.filters.http.router
                    typed_config: {}

    - name: staticreply
      address:
        socket_address:
          address: 127.0.0.1
          port_value: 8099
      filter_chains:
        - filters:
            - name: envoy.http_connection_manager
              typed_config:
                "@type": type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager
                stat_prefix: ingress_http
                codec_type: auto
                route_config:
                  name: local_route
                  virtual_hosts:
                    - name: local_service
                      domains:
                        - "*"
                      routes:
                        - match:
                            prefix: "/"
                          direct_response:
                            status: 200
                            body:
                              inline_string: "example body\n"
                http_filters:
                  - name: envoy.filters.http.router
                    typed_config: {}

  clusters:
    - name: web_service
      connect_timeout: 0.25s
      type: STATIC
      lb_policy: ROUND_ROBIN
      load_assignment:
        cluster_name: mock_service
        endpoints:
          - lb_endpoints:
              - endpoint:
                  address:
                    socket_address:
                      address: 127.0.0.1
                      port_value: 8099

  admin:
    access_log_path: "/dev/null"
    address:
      socket_address:
        address: 0.0.0.0
        port_value: 8001
  ```

</details>

### Run with Docker

Run an the `envoyproxy/envoy-dev:latest` container with the required config:

```console
$ mkdir -p /tmp/envoy/filters        # setup temp dir to share with container
$ cp envoy.yaml  /tmp/envoy          # copy full envoy config into dir
$ cp example.wasm /tmp/envoy/filters # copy wasm binary into dir
$ docker run --entrypoint='/usr/local/bin/envoy' \
    -p 18000:18000 -p 8099:8099 \
    -w /tmp/envoy \
    -v /tmp/envoy:/tmp/envoy \
    envoyproxy/envoy-dev:latest \
    -c /tmp/envoy/envoy.yaml --concurrency 2 --log-level info --log-format '%v'
```

```console
$ curl http://localhost:18000 -v
> GET / HTTP/1.1
> Host: localhost:18000
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< content-length: 13
< content-type: text/plain
< date: Sat, 05 Dec 2020 00:51:37 GMT
< server: envoy
< x-envoy-upstream-service-time: 0
< wasm: enabled
<
example body
```

☝️ The `wasm: enabled` header was successfully applied to the response from the custom WASM filter! 🎉

## Credits, Links, More Information, and Other SDK Options

Finding clear information about how this proxy WASM filter stuff works essentially requires you to navigate several SDK codebases,
and finding extra information from GitHub comments, YouTube, or Twitter. This is still the early days of proxies generally supporting WASM filters, and really Envoy is the target proxy for most of this work. But, even Envoy doesn't yet have official documentation for the ABI (Application Binary Interface) to enable this feature. This should come in the near-ish future. Still looking for more information around this stuff myself.

With the lack of documentation, I really wouldn't have been able to make this SDK without the previous work from the community including the C++, Rust, AssemblyScript, and TinyGo code bases. There were also several talks avaialble on YouTube to help provide context. I am really looking forward to the growth of this WASM extension ecosystem, and how other proxies might implement this feature in the future.

### 🔗 Links to Helpful Information

* [Specification](https://github.com/proxy-wasm/spec)
* [WebAssembly Extensions For Network Proxies - John Plevyak](https://www.youtube.com/watch?v=OIUPf8m7CGA)
* [Extending Envoy with WebAssembly - John Plevyak & Dhi Aurrahman](https://www.youtube.com/watch?v=XdWmm_mtVXI)
* [Building HTTP Request Filters for Consul with Web Assembly Hub - Nic Jackson, Christian Posta](https://www.youtube.com/watch?v=qyXqMKziFaE)
* [Building idiomatic Envoy SDKs for Rust and Go - Yaroslav Skopets, Takeshi Yoneda](https://www.youtube.com/watch?v=bqIaAp4EIkg)
* [Hands-on WASM filters and singletons - Emmanuel Mayssat](https://www.youtube.com/watch?v=BZsyqYiD520)
* [Tiny Go Website](https://tinygo.org/)
* [Tiny Go is an Official Google Sponsored Project](https://twitter.com/TinyGolang/status/1223887654158307328)
* [Tiny Go Important Build Options](https://tinygo.org/usage/important-options/)
* [Tiny Go - Reduce your WebAssembly binaries 72% - from 56KB to 26KB to 16KB](https://dev.to/sendilkumarn/reduce-your-webassembly-binaries-72-from-56kb-to-26kb-to-16kb-40mi)

### 🛠 Other SDKs

* [C++ SDK](https://github.com/proxy-wasm/proxy-wasm-cpp-sdk)
* [Rust SDK](https://github.com/proxy-wasm/proxy-wasm-rust-sdk)
* [Tiny Go SDK](https://github.com/tetratelabs/proxy-wasm-go-sdk)
* [AssemblyScript SDK](https://github.com/solo-io/proxy-runtime)

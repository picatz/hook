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
              name       = "security_headers"
              local_file = "/tmp/filters/security_headers.wasm"
            }
          ]
        }
      }
    }
  }
}
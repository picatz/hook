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
              name       = "header_auth"
              local_file = "/tmp/filters/header_auth.wasm"
            }
          ]
        }
      }
    }
  }
}
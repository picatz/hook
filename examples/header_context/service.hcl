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
              name       = "header"
              local_file = "/tmp/filters/header.wasm"
            }
          ]
        }
      }
    }
  }
}
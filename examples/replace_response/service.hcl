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
}
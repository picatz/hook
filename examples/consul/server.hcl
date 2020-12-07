data_dir           = "/tmp/"
log_level          = "DEBUG"
datacenter         = "dc1"
primary_datacenter = "dc1"
server             = true
bootstrap_expect   = 1
bind_addr          = "0.0.0.0"

ui_config {
  enabled = true
}

ports {
  grpc = 8502
}

connect {
  enabled = true
}
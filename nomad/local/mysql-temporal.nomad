job "mysql-temporal" {
  datacenters = ["dc1"]
  type        = "service"

  group "mysql-temporal" {
    count = 1

    restart {
      attempts = 10
      interval = "5m"
      delay    = "25s"
      mode     = "delay"
    }

    task "mysql-server" {
      driver = "docker"

      env = {
        "MYSQL_ROOT_PASSWORD" = "password"
      }

      config {
        image = "hashicorp/mysql-portworx-demo:latest"

        ports = ["db"]
      }

      resources {
        cpu    = 500
        memory = 1024
      }

      service {
        name = "mysql-server"
        tags = [
          "traefik.enable=true",
        ]
        port = "db"

        check {
          type     = "tcp"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }
    network {
      port "db" {
        static = 3306
      }
    }
  }
}


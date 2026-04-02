terraform {
  required_version = "~> 1.1, >= 1.1.2"

  required_providers {
    postgresql = {
      source  = "cyrilgdn/postgresql"
      version = "1.26.0"
    }
  }

}

provider "postgresql" {
  host            = "localhost"
  port            = 6001
  database        = "app_db"
  username        = "app_user"
  password        = "app_password"
  connect_timeout = 15
  sslmode         = "disable"
}

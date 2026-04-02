resource "postgresql_schema" "random_data" {
  name  = "random_data"
  owner = "app_user"
}

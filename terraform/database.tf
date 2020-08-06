resource "random_id" "db_name_suffix" {
  byte_length = 5
}

resource "google_sql_database_instance" "master" {
  provider         = google-beta
  name             = "master-instance-${random_id.db_name_suffix.hex}"
  database_version = "MYSQL_5_7"
  region           = var.region
  depends_on       = [google_service_networking_connection.private_vpc_connection]
  settings {
    tier = "db-n1-standard-1"
    ip_configuration {
      ipv4_enabled    = false
      private_network = google_compute_network.private_network.id
    }
  }
}

resource "google_sql_database" "database" {
  project  = var.project_id
  name     = var.db_name
  instance = google_sql_database_instance.master.name
  charset  = "utf8mb4"
}

resource "google_sql_user" "users" {
  project  = var.project_id
  name     = var.db_username
  instance = google_sql_database_instance.master.name
  host     = "%"
  password = var.db_password
}

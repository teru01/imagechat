resource "google_storage_bucket" "bucket" {
  project       = var.project_id
  name          = var.bucket_name
  location      = "ASIA"
  force_destroy = true

  bucket_policy_only = true

  cors {
    origin          = ["https://${var.my_domain}", "http://${var.my_domain}"]
    method          = ["GET", "HEAD", "PUT", "POST", "DELETE"]
    response_header = ["*"]
    max_age_seconds = 3600
  }
}

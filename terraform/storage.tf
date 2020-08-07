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

resource "google_storage_bucket_iam_member" "member" {
  bucket = google_storage_bucket.bucket.name
  role = "roles/storage.objectViewer"
  member = "allUsers"
}

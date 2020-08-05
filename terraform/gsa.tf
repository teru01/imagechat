locals {
  annotations = formatlist(
    "serviceAccount:%s.svc.id.goog[%s]", var.project_id, var.k8s_svc_accounts,
  )
}

resource "google_service_account" "service_account" {
  project      = var.project_id
  account_id   = var.gke_service_account
  display_name = var.gke_service_account
}

data "google_iam_policy" "sqladmin" {
  binding {
    role = "roles/cloudsql.client"

    members = [
      "serviceAccount:${google_service_account.service_account.email}",
    ]
  }
}

resource "google_project_iam_binding" "cloudsql-binding" {
  project = var.project_id
  role    = "roles/cloudsql.client"
  members = [
    "serviceAccount:${google_service_account.service_account.email}",
  ]
}

resource "google_service_account_iam_binding" "ksa-binding" {
  service_account_id = google_service_account.service_account.name
  role               = "roles/iam.workloadIdentityUser"
  members            = local.annotations
}

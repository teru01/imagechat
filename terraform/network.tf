resource "google_compute_network" "private_network" {
  provider = google-beta
  name     = "private-network"
}

resource "google_compute_subnetwork" "imagechat-subnet" {
  project       = var.project_id
  name          = "imagechat-subnet"
  ip_cidr_range = "10.2.0.0/16"
  region        = var.region
  network       = google_compute_network.private_network.id
}

resource "google_compute_global_address" "private_ip_address" {
  provider = google-beta

  name          = "private-ip-address"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.private_network.id
}

resource "google_service_networking_connection" "private_vpc_connection" {
  provider = google-beta

  network                 = google_compute_network.private_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_address.name]
}

resource "google_dns_managed_zone" "imagechat-zone" {
  project  = var.project_id
  name     = "imagechat-zone"
  dns_name = var.dns_name
}

resource "google_dns_record_set" "record" {
  project = var.project_id
  name    = google_dns_managed_zone.imagechat-zone.dns_name
  type    = "A"
  ttl     = 300

  managed_zone = google_dns_managed_zone.imagechat-zone.name

  rrdatas = [google_compute_global_address.default.address]
}

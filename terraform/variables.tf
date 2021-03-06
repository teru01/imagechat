# 環境変数から読み込ませる
variable "project_id" {}

variable "db_username" {}
variable "db_password" {}
variable "dns_name" {}
variable "cert_path" {}
variable "key_path" {}

variable "terraform_account_email" {}

variable "bucket_name" {
  default = "imagechet-storage-3344"
}

variable "my_domain" {
  default = "imagechat.ga"
}

variable db_name {
  default = "imagechat-db"
}

variable "zone" {
  default = "asia-northeast1-a"
}

variable "region" {
  default = "asia-northeast1"
}

variable "k8s_svc_accounts" {
  type    = list(string)
  default = ["default/myksa"]
}

variable "gke_service_account" {
  default = "gketosql"
}

variable "k8s_service_account" {
  default = "myksa"
}

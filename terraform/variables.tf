# 環境変数から読み込ませる
variable "project_id" {}

variable "db_username" {}
variable "db_password" {}

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

variable "dns_name" {}

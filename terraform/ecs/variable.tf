variable "app_name" {}

variable "vpc_id" {}

variable "db_user" {}

variable "db_password" {}

variable "db_host" {}

variable "db_name" {}

variable "app_names" {
  default = ["shimo0108-app"]
}

variable "http_listener_arn" {}

variable "https_listener_arn" {}

variable "alb_security_group_id" {}

variable "cluster_name" {}

variable "public_subnet_ids" {}

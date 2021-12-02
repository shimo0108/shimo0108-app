variable "app_name" {}
variable "vpc_id" {}
variable "ingress_ports" {
  description = "list of ingress ports"
  default     = [80, 443, 9999]
}
variable "public_subnet_ids" {}
variable "target_group_arn" {}
variable "sub_acm_id" {}
variable "domain" {}

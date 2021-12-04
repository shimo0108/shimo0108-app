variable "aws_region" {
  type    = string
  default = "ap-northeast-1"
}
variable "aws_profile" {
  type        = string
  default     = "shimo0108"
  description = "AWS CLI's profile"
}
variable "app_name" {
  type    = string
  default = "shimo0108-app"
}

variable "domain" {}

variable "db_name" {}

variable "db_user" {}

variable "db_password" {}

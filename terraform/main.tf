module "network" {
  source   = "./network"
  app_name = var.app_name
}

module "acm" {
  source = "./acm"
  domain = var.domain
}

module "spa" {
  source   = "./spa"
  app_name = var.app_name
  domain   = var.domain
  acm_id   = module.acm.acm_id
}

module "subdomain_acm" {
  source = "./subdomain_acm"
  domain = var.domain
}

module "ecs_cluster" {
  source = "./ecs_cluster"

  app_name = var.app_name
}

module "elb" {
  source = "./elb"
  app_name          = var.app_name
  vpc_id            = module.network.vpc_id
  public_subnet_ids = module.network.public_subnet_ids
  sub_acm_id        = module.subdomain_acm.sub_acm_id
  domain            = var.domain
  target_group_arn  = module.ecs.target_group_arn
}

module "rds" {
  source = "./rds"

  app_name    = var.app_name
  db_name     = var.db_name
  db_user     = var.db_user
  db_password = var.db_password

  vpc_id                = module.network.vpc_id
  alb_security_group_id = module.elb.alb_security_group_id
  private_subnet_ids    = module.network.private_subnet_ids
}

module "ecs" {
  source = "./ecs"
  app_name = var.app_name
  db_name        = var.db_name
  db_user        = var.db_user
  db_password    = var.db_password
  db_host        = module.rds.db_address

  vpc_id                = module.network.vpc_id
  http_listener_arn     = module.elb.http_listener_arn
  https_listener_arn    = module.elb.https_listener_arn
  alb_security_group_id = module.elb.alb_security_group_id
  cluster_name          = module.ecs_cluster.cluster_name
  public_subnet_ids     = module.network.public_subnet_ids
}

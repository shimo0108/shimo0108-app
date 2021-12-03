data "aws_caller_identity" "user" {}

data "aws_iam_policy" "ecs_task_execution_role_policy" {
  arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

data "aws_iam_policy_document" "ecs_task_execution" {
  source_json = data.aws_iam_policy.ecs_task_execution_role_policy.policy

  statement {
    effect    = "Allow"
    actions   = ["ssm:GetParameters", "kms:Decrypt"]
    resources = ["*"]
  }
}

data "template_file" "app_container_definitions" {
  template = file("./ecs/app_service.json")

  vars = {
    account_id  = local.account_id
    db_host     = var.db_host
    db_user     = var.db_user
    db_password = var.db_password
    db_name     = var.db_name
    ssm_agent_code = var.ssm_agent_code
    ssm_agent_id   = var.ssm_agent_id
  }
}

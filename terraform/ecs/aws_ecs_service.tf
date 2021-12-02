locals {
  account_id = data.aws_caller_identity.user.account_id
}

resource "aws_iam_policy" "ecs_task_execution" {
  name   = "ecs-task-execution"
  policy = data.aws_iam_policy_document.ecs_task_execution.json
}

module "ecs_task_execution_role" {
  source     = "../iam_role"
  name       = "ecs-task-execution"
  identifier = "ecs-tasks.amazonaws.com"
  policy_arn = [aws_iam_policy.ecs_task_execution.arn]
}

resource "aws_ecs_task_definition" "app" {
  family = "${var.app_name}-task"
  cpu                      = 256
  memory                   = 1024
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  container_definitions = data.template_file.app_container_definitions.rendered
  execution_role_arn    = module.ecs_task_execution_role.iam_role_arn
}


resource "aws_lb_target_group" "this" {
  name_prefix = "target"
  vpc_id      = var.vpc_id
  port        = 80
  protocol    = "HTTP"

  target_type = "ip"
  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_lb_listener_rule" "http_rule" {
  listener_arn = var.http_listener_arn
  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.this.id
  }
  condition {
    path_pattern {
      values = ["*"]
    }
  }
}
resource "aws_lb_listener_rule" "https_rule" {
  listener_arn = var.https_listener_arn
  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.this.id
  }
  condition {
    path_pattern {
      values = ["*"]
    }
  }
}

resource "aws_ecs_service" "shimo0108-app" {
  name            = "shimo0108-app-service"
  launch_type     = "FARGATE"
  desired_count   = 1
  cluster         = var.cluster_name
  task_definition = aws_ecs_task_definition.app.arn
  network_configuration {
    security_groups  = [var.alb_security_group_id]
    subnets          = var.public_subnet_ids
    assign_public_ip = true
  }
  load_balancer {
    target_group_arn = aws_lb_target_group.this.arn
    container_name   = "server"
    container_port   = 9999
  }
}

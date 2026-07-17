provider "aws" {
  region = var.aws_region

  default_tags {
    tags = {
      Project     = local.project_name
      Environment = local.environment
      ManagedBy   = "Terraform"
      Owner       = var.owner
    }
  }
}

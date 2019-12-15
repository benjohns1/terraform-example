provider "aws" {
  region = var.aws_region
}

# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY A SERVERLESS APP
# ---------------------------------------------------------------------------------------------------------------------

module "example_app" {
  source = "../../modules/aws-serverless-app"

  name = var.name

  deploy_zip = "${path.module}/../build/artifacts/app.zip"
  runtime    = "go1.x"
  handler    = "app"
  proxy_path = "app"

  environment_variables = {
    SampleEnvVar = "Environment variable set during deployment"
  }
}

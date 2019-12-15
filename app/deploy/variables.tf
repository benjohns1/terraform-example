# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------

variable "aws_region" {
  description = "The AWS region to deploy into"
  type        = string
  default     = "us-west-2"
}

variable "name" {
  description = "The name of the Lambda function and all other resources created by this module"
  type        = string
  default     = "terraform-example-app"
}

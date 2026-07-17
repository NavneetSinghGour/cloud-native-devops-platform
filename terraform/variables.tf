variable "aws_region" {
  description = "AWS Region"
  type        = string
}

variable "owner" {
  description = "Owner of the infrastructure"
  type        = string
}

variable "vpc_cidr" {
  description = "CIDR block for the VPC"
  type        = string
}

variable "public_subnet_1_cidr" {
  description = "CIDR block for Public Subnet A"
  type        = string
}

variable "public_subnet_2_cidr" {
  description = "CIDR block for Public Subnet B"
  type        = string
}

variable "key_name" {
  description = "AWS EC2 Key Pair Name"
  type        = string
}

variable "instance_type" {
  description = "EC2 Instance Type"
  type        = string
}

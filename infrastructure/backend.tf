terraform {
  backend "s3" {
    bucket         = "iac-tfstate-solo-exam-alison"
    key            = "ecs/api.tfstate"
    region         = "us-east-1"
    dynamodb_table = "terraform_solo"
  }
}
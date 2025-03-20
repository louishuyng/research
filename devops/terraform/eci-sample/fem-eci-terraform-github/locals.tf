locals {
  repos = {
    "fem-eci-terraform-tfe" = {
      description        = "Automation for Terraform Enterprise"
      gitignore_template = "Terraform"
      name               = "fem-eci-terraform-tfe"
      topics             = ["terraform"]
      visibility         = "private"
    }

    "fem-eci-terraform-github" = {
      description        = "Automation for GitHub"
      gitignore_template = "Terraform"
      name               = "fem-eci-terraform-github"
      topics             = ["terraform"]
      visibility         = "private"
    }

    "fem-eci-terraform-aws-network" = {
      description        = "Automation for AWS Network"
      gitignore_template = "Terraform"
      name               = "fem-eci-terraform-aws-network"
      topics             = ["terraform"]
      visibility         = "private"
    }

    "fem-eci-terraform-aws-cluster" = {
      description        = "Automation for AWS Cluster"
      gitignore_template = "Terraform"
      name               = "fem-eci-terraform-aws-cluster"
      topics             = ["terraform"]
      visibility         = "private"
    }

    "fem-eci-terraform-product-service" = {
      description        = "Automation for product service"
      gitignore_template = "Terraform"
      name               = "fem-eci-terraform-product-service"
      topics             = ["terraform"]
      visibility         = "private"
    }

    "fem-eci-service" = {
      description        = "Example product service"
      gitignore_template = "Go"
      name               = "fem-eci-service"
      topics             = ["backend"]
      visibility         = "private"
    }

    "my-spotify-tfe" = {
      description        = "Spotify Terraform"
      gitignore_template = "Terraform"
      name               = "my-spotify-tfe"
      topics             = ["terraform"]
      visibility         = "private"
    }
  }
}

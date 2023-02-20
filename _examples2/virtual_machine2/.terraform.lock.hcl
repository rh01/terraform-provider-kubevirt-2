# This file is maintained automatically by "terraform init".
# Manual edits may be lost in future updates.

provider "example.com/exampleprovider/kubevirt" {
  version     = "1.0.0"
  constraints = "~> 1.0.0"
  hashes = [
    "h1:Ln3rsLRINocpy+rNZ1rkk71Av8UmVpprM2fBAOZ8z+U=",
  ]
}

provider "registry.terraform.io/hashicorp/kubernetes" {
  version     = "2.17.0"
  constraints = ">= 1.10.0"
  hashes = [
    "h1:Dq/EHg8mKP9wDDTJx5CzZ+w44wutIZJGfQLrAIznAqY=",
  ]
}

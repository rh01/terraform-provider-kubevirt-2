# This file is maintained automatically by "terraform init".
# Manual edits may be lost in future updates.

provider "example.com/exampleprovider/kubevirt" {
  version     = "1.0.0"
  constraints = "~> 1.0.0"
  hashes = [
    "h1:NsFWpakCVzbrGpuSCt+N6IYL4Zw6eEKDbXBuOqEeyEk=",
  ]
}

provider "registry.terraform.io/gitlabhq/gitlab" {
  version     = "3.18.0"
  constraints = "~> 3.18.0"
  hashes = [
    "h1:JZsPjdsOqjG6l+s96d7Awp4XZ9Fwvgv7S1kCLZExuHA=",
  ]
}

provider "registry.terraform.io/hashicorp/kubernetes" {
  version     = "2.17.0"
  constraints = ">= 1.10.0"
  hashes = [
    "h1:Dq/EHg8mKP9wDDTJx5CzZ+w44wutIZJGfQLrAIznAqY=",
  ]
}

provider "registry.terraform.io/hashicorp/random" {
  version     = "3.4.3"
  constraints = ">= 2.0.0"
  hashes = [
    "h1:saZR+mhthL0OZl4SyHXZraxyaBNVMxiZzks78nWcZ2o=",
  ]
}

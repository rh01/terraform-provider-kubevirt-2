# This file is maintained automatically by "terraform init".
# Manual edits may be lost in future updates.

provider "example.com/exampleprovider/kubevirt" {
  version     = "1.0.0"
  constraints = "~> 1.0.0"
  hashes = [
    "h1:rvRgl7saoHYSjaNxHMvFsUX6uN1jVgYYFH9iReXVTF4=",
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
  version     = "2.18.0"
  constraints = ">= 1.10.0"
  hashes = [
    "h1:H81uBnhN5NDpFrZ39Q9mEIY2bcjZv9oyT4+BlQgeFzM=",
  ]
}

provider "registry.terraform.io/hashicorp/random" {
  version     = "3.4.3"
  constraints = ">= 2.0.0"
  hashes = [
    "h1:saZR+mhthL0OZl4SyHXZraxyaBNVMxiZzks78nWcZ2o=",
  ]
}

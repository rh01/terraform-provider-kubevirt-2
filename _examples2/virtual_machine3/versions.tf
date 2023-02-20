terraform {
  required_version = ">= 0.13"

  required_providers {

    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = ">= 1.10.0"
    }
    random = {
      source  = "hashicorp/random"
      version = ">= 2.0"
    }
    gitlab = {
      source  = "gitlabhq/gitlab"
      version = "~> 3.18.0"
    }
    kubevirt = {
      source  = "example.com/exampleprovider/kubevirt"
      version = "~> 1.0.0"
    }
  }
}

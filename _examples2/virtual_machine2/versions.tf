terraform {
  required_version = ">= 0.13"

  required_providers {

    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = ">= 1.10.0"
    }
  
    kubevirt = {
      source  = "example.com/exampleprovider/kubevirt"
      version = "~> 1.0.0"
    }
  }
}

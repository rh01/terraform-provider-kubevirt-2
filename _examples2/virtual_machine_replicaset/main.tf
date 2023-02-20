provider "kubevirt" {
  host = var.host

  client_certificate     = base64decode(var.client_certificate)
  client_key             = base64decode(var.client_key)
  cluster_ca_certificate = base64decode(var.cluster_ca_certificate)
}

locals {
  # 定义虚拟机的数量
  vm_count = 1
  # 定义使用image
  image = "harbor-registry.inner.youdao.com/kubevirt/win10-qcow2-v1"
  # 虚拟机名称，前缀
  vm_name_preffix = "testcmcc"
  # 命名空间
  namespace = "test"
}


resource "kubevirt_virtual_machine_instance_replica_set" "virtual_machine_instance_replica_set" {

  metadata {
    name      = local.vm_name_preffix
    namespace = local.namespace
     
  }
  spec {
    replicas = "2"

    selector {
      match_labels = {
        "kubevirt.io/vmReplicaSet" = "${local.vm_name_preffix}"
      }
    }


    template {
      metadata {
        labels = {
          "kubevirt.io/vmReplicaSet" = "${local.vm_name_preffix}"
        }
        namespace = local.namespace
      }
      spec {

        domain {
          machine {
            type = "q35"
          }
          resources {
            requests = {
              memory = "4G"
              cpu    = 4
            }
          }

          devices {
            disk {
              name = "${local.vm_name_preffix}-datavolumedisk1"
              disk_device {
                disk {
                  bus = "virtio"
                }
              }
            }

            interface {
              name                     = "default"
              interface_binding_method = "InterfaceMasquerade"
              model                    = "e1000"
            }

          }

        }
        network {
          name = "default"
          network_source {
            pod {
            }
          }
        }
        termination_grace_period_seconds = 0

        # readiness_probe {
        #   tcp_socket {
        #     port = "7333"
        #   }
        # }
        # liveness_probe {
        #   tcp_socket {
        #     port = "7333"
        #   }
        # }

        volume {
          name = "${local.vm_name_preffix}-datavolumedisk1"
          volume_source {
            container_disk {
              image             = local.image
              image_pull_secret = "regcred"
            }
          }
        }
      }
    }
  }
}

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
  vm_name_preffix = "testcmc"
  # 命名空间
  namespace = "test"
}




resource "kubevirt_virtual_machine" "virtual_machine" {
  count = local.vm_count

  metadata {
    name      = "${local.vm_name_preffix}-${count.index}"
    namespace = local.namespace
    labels = {
      "key1" = "value1"
      "x"    = "1"
    }
  }
  spec {
    run_strategy = "Always"

    template {
      metadata {
        labels = {
          "kubevirt.io/vm" = "${local.vm_name_preffix}-${count.index}"
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
              name = "${local.vm_name_preffix}-${count.index}-datavolumedisk1"
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

        # readiness_probe {
        #   tcp_socket {
        #     port = "7333"
        #   }
        # }
        liveness_probe {
          tcp_socket {
            port = "7333"
          }
        }

        volume {

          name = "${local.vm_name_preffix}-${count.index}-datavolumedisk1"
          volume_source {
            container_disk {
              image             = local.image
              image_pull_secret = "regcred"
            }
          }
        }

        # affinity {
        #   pod_anti_affinity {
        #     preferred_during_scheduling_ignored_during_execution {
        #       weight = 100
        #       pod_affinity_term {
        #         label_selector {
        #           match_labels = {
        #             anti-affinity-key = "anti-affinity-val"
        #           }
        #         }
        #         topology_key = "kubernetes.io/hostname"
        #       }
        #     }
        #   }
        # }
      }
    }
  }
}

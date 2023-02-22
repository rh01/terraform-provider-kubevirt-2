package virtualmachine

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	apiresource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
<<<<<<< HEAD
	kubevirtapiv1 "kubevirt.io/api/core/v1"
	cdiv1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
=======
	kubevirtapiv1 "kubevirt.io/client-go/api/v1"
	cdiv1 "kubevirt.io/containerized-data-importer/pkg/apis/core/v1alpha1"
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
)

func getExpectedVirtualMachine(name string, namespace string, source cdiv1.DataVolumeSource, labels map[string]string) *kubevirtapiv1.VirtualMachine {
	return &kubevirtapiv1.VirtualMachine{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    labels,
			Annotations: map[string]string{
<<<<<<< HEAD
				"kubevirt.io/storage-observed-api-version": "v1alpha3",
				"kubevirt.io/latest-observed-api-version":  "v1",
=======
				"kubevirt.io/latest-observed-api-version":  "v1",
				"kubevirt.io/storage-observed-api-version": "v1alpha3",
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
			},
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "VirtualMachine",
			APIVersion: kubevirtapiv1.GroupVersion.String(),
		},
		Spec: kubevirtapiv1.VirtualMachineSpec{
			RunStrategy: func(src kubevirtapiv1.VirtualMachineRunStrategy) *kubevirtapiv1.VirtualMachineRunStrategy {
				return &src
			}(kubevirtapiv1.RunStrategyAlways),
<<<<<<< HEAD
			DataVolumeTemplates: []kubevirtapiv1.DataVolumeTemplateSpec{
=======
			DataVolumeTemplates: []cdiv1.DataVolume{
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      fmt.Sprintf("%s-bootvolume", name),
						Namespace: namespace,
					},
					Spec: cdiv1.DataVolumeSpec{
<<<<<<< HEAD
						Source: &source,
=======
						Source: source,
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
						PVC: &corev1.PersistentVolumeClaimSpec{
							AccessModes: []corev1.PersistentVolumeAccessMode{
								"ReadWriteOnce",
							},
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									"storage": (func() resource.Quantity { res, _ := resource.ParseQuantity("10Gi"); return res })(),
								},
							},
						},
					},
				},
			},
			Template: &kubevirtapiv1.VirtualMachineInstanceTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: namespace,
					Labels:    map[string]string{"templateKey1": "templateVal1"},
				},
				Spec: kubevirtapiv1.VirtualMachineInstanceSpec{
					Volumes: []kubevirtapiv1.Volume{
						{
							Name: "datavolumedisk1",
							VolumeSource: kubevirtapiv1.VolumeSource{
								DataVolume: &kubevirtapiv1.DataVolumeSource{
									Name: fmt.Sprintf("%s-bootvolume", name),
								},
							},
						},
					},
					Networks: []kubevirtapiv1.Network{
						{
							Name: "default",
							NetworkSource: kubevirtapiv1.NetworkSource{
								Pod: &kubevirtapiv1.PodNetwork{},
							},
						},
					},
					Domain: kubevirtapiv1.DomainSpec{
						Resources: kubevirtapiv1.ResourceRequirements{
							Requests: corev1.ResourceList{
<<<<<<< HEAD
								corev1.ResourceCPU:    apiresource.MustParse(fmt.Sprint(1)),
								corev1.ResourceMemory: apiresource.MustParse("120Mi"),
							},
						},
						Machine: &kubevirtapiv1.Machine{
=======
								corev1.ResourceMemory: apiresource.MustParse("120Mi"),
								corev1.ResourceCPU:    apiresource.MustParse(fmt.Sprint(1)),
							},
						},
						Machine: kubevirtapiv1.Machine{
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
							Type: "q35",
						},
						Devices: kubevirtapiv1.Devices{
							Disks: []kubevirtapiv1.Disk{
								{
									Name: "datavolumedisk1",
									DiskDevice: kubevirtapiv1.DiskDevice{
										Disk: &kubevirtapiv1.DiskTarget{
											Bus: "virtio",
										},
									},
								},
							},
							Interfaces: []kubevirtapiv1.Interface{
								{
									Name: "default",
									InterfaceBindingMethod: kubevirtapiv1.InterfaceBindingMethod{
										Bridge: &kubevirtapiv1.InterfaceBridge{},
									},
								},
							},
						},
					},
					Affinity:                      &corev1.Affinity{},
					TerminationGracePeriodSeconds: func(src int64) *int64 { return &src }(0),
					DNSConfig:                     &corev1.PodDNSConfig{},
				},
			},
		},
		Status: kubevirtapiv1.VirtualMachineStatus{
<<<<<<< HEAD
			Created:         true,
			Ready:           true,
			PrintableStatus: "Running",
			VolumeSnapshotStatuses: []kubevirtapiv1.VolumeSnapshotStatus{
				{
					Name:    "datavolumedisk1",
					Enabled: false,
					Reason:  "No VolumeSnapshotClass: Volume snapshots are not configured for this StorageClass [local] [datavolumedisk1]",
				},
			},
=======
			Created: true,
			Ready:   true,
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
		},
	}
}

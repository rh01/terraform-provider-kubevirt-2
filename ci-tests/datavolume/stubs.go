package datavolume

import (
	k8sv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
<<<<<<< HEAD
	cdiv1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
=======
	cdiv1 "kubevirt.io/containerized-data-importer/pkg/apis/core/v1alpha1"
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
)

func getExpectedDataVolume(name string, namespace string, source cdiv1.DataVolumeSource, labels map[string]string) *cdiv1.DataVolume {
	return &cdiv1.DataVolume{
		TypeMeta: metav1.TypeMeta{
			Kind:       "DataVolume",
<<<<<<< HEAD
			APIVersion: "cdi.kubevirt.io/v1beta1",
=======
			APIVersion: "cdi.kubevirt.io/v1alpha1",
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:         name,
			GenerateName: "",
			Namespace:    namespace,
			Labels:       labels,
		},
		Spec: cdiv1.DataVolumeSpec{
<<<<<<< HEAD
			Source: &source,
=======
			Source: source,
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
			PVC: &k8sv1.PersistentVolumeClaimSpec{
				AccessModes: []k8sv1.PersistentVolumeAccessMode{
					"ReadWriteOnce",
				},
				Resources: k8sv1.ResourceRequirements{
					Requests: k8sv1.ResourceList{
						"storage": (func() resource.Quantity { res, _ := resource.ParseQuantity("10Gi"); return res })(),
					},
				},
			},
		},
		Status: cdiv1.DataVolumeStatus{
<<<<<<< HEAD
			ClaimName: name,
			Phase:     "Succeeded",
			Progress:  "100.0%",
=======
			Phase:    "Succeeded",
			Progress: "100.0%",
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
		},
	}
}

package main

import (
	"fmt"
	myv1alpha1 "github.com/Arnobkumarsaha/messi/pkg/apis/arnob.com/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// newDeployment creates a new Deployment for a messi resource. It also sets
// the appropriate OwnerReferences on the resource so handleObject can discover
// the messi resource that 'owns' it.
func newDeployment(lm10 *myv1alpha1.Messi) *appsv1.Deployment {
	fmt.Println("newDeployment is called")
	labels := map[string]string{
		"app":       lm10.Spec.DeploymentImage,
		"controller": lm10.Name,
	}
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      lm10.Spec.DeploymentName,
			Namespace: lm10.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(lm10, myv1alpha1.SchemeGroupVersion.WithKind("Messi")),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: lm10.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "messi-container",
							Image: lm10.Spec.DeploymentImage,
						},
					},
				},
			},
		},
	}
}

func newService() *corev1.Service {
	// ctx := context.TODO()

	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "mysvc",
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": "demo",
			},
			Type: corev1.ServiceTypeNodePort,
			Ports: []corev1.ServicePort{
				{
					NodePort: int32(30012),
					Port: 2345,
					TargetPort: intstr.IntOrString{
						IntVal: 8080,
					},
				},
			},
		},
	}
	/*_, err:= svcClient.Create(ctx,&svc, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
	}*/
}

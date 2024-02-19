package handlers

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesHandler struct {
	Namespace string
	Name      string
	Image     string
	Replicas  int32
}

func (k *KubernetesHandler) CreateDeployment() error {
	// Create a Kubernetes client
	config, err := clientcmd.BuildConfigFromFlags("", "/path/to/kubeconfig")
	if err != nil {
		return fmt.Errorf("failed to build config: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create clientset: %v", err)
	}

	// Create a Deployment object
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      k.Name,
			Namespace: k.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &k.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": k.Name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": k.Name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  k.Name,
							Image: k.Image,
						},
					},
				},
			},
		},
	}

	// Create the Deployment
	_, err = clientset.AppsV1().Deployments(k.Namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create Deployment: %v", err)
	}

	return nil

}

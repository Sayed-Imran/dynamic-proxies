package handlers

import (
	"context"
	"fmt"

	"github.com/sayed-imran/dynamic-proxies/config"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesHandler struct {
	Namespace string
	Name      string
	Image     string
	Replicas  int32
	Port      int32
}

func (k *KubernetesHandler) CreateDeployment() error {
	// Create a Kubernetes client
	var baseConfig  = config.BaseConfig{}
	err := config.LoadConfig(&baseConfig)
	fmt.Println(baseConfig.KubeConfigPath)
	fmt.Println(baseConfig.Namespace)
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}
	config, err := clientcmd.BuildConfigFromFlags("", "config.yml")
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
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: k.Port,
								},
							},
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

func (k *KubernetesHandler) CreateService() error {
	// Create a Kubernetes client
	config, err := clientcmd.BuildConfigFromFlags("", "/path/to/kubeconfig")
	if err != nil {
		return fmt.Errorf("failed to build config: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create clientset: %v", err)
	}

	// Create a Service object
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      k.Name,
			Namespace: k.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": k.Name,
			},
			Ports: []corev1.ServicePort{
				{
					Protocol:   corev1.ProtocolTCP,
					Port:       k.Port,
					TargetPort: intstr.FromInt(int(k.Port)),
				},
			},
		},
	}

	// Create the Service
	_, err = clientset.CoreV1().Services(k.Namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create Service: %v", err)
	}

	return nil
}

func (k *KubernetesHandler) DeleteDeployment() error {
	// Create a Kubernetes client
	config, err := clientcmd.BuildConfigFromFlags("", "config.yml")
	if err != nil {
		return fmt.Errorf("failed to build config: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create clientset: %v", err)
	}

	// Delete the Deployment
	err = clientset.AppsV1().Deployments(k.Namespace).Delete(context.TODO(), k.Name, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete Deployment: %v", err)
	}

	return nil
}

func (k *KubernetesHandler) DeleteService() error {
	// Create a Kubernetes client
	config, err := clientcmd.BuildConfigFromFlags("", "config.yml")
	if err != nil {
		return fmt.Errorf("failed to build config: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create clientset: %v", err)
	}

	// Delete the Service
	err = clientset.CoreV1().Services(k.Namespace).Delete(context.TODO(), k.Name, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete Service: %v", err)
	}

	return nil
}

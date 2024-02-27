package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io"

	customresource "github.com/sayed-imran/dynamic-proxies/custom_resource"
	errorHandler "github.com/sayed-imran/dynamic-proxies/handlers/error_handler"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

type KubernetesHandler struct {
	Clientset     *kubernetes.Clientset
	DynamicClient dynamic.Interface
	Namespace     string
	Name          string
	Image         string
	Replicas      int32
	Port          int32
}

func (k *KubernetesHandler) CreateDeployment() error {

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
	_, err := k.Clientset.AppsV1().Deployments(k.Namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create Deployment: %v", err)
	}

	return nil

}

func (k *KubernetesHandler) CreateService() error {

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
	_, err := k.Clientset.CoreV1().Services(k.Namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create Service: %v", err)
	}

	return nil
}

func (k *KubernetesHandler) CreateVirtualService() error {

	// Create a VirtualService object
	virtualService := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "networking.istio.io/v1alpha3",
			"kind":       "VirtualService",
			"metadata": map[string]string{
				"name":      k.Name,
				"namespace": k.Namespace,
			},
			"spec": map[string]interface{}{
				"hosts": []string{
					"*",
				},
				"gateways": []string{
					"istio-system/microservice-gateway",
				},
				"http": []map[string]interface{}{
					{
						"match": []map[string]interface{}{
							{
								"uri": map[string]interface{}{
									"prefix": "/",
								},
							},
						},
						"route": []map[string]interface{}{
							{
								"destination": map[string]interface{}{
									"host": fmt.Sprintf("%s.%s.svc.cluster.local", k.Name, k.Namespace),
									"port": map[string]interface{}{
										"number": k.Port,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Create the VirtualService
	_, err := k.DynamicClient.Resource(customresource.VirtualService).Namespace(k.Namespace).Create(context.TODO(), virtualService, metav1.CreateOptions{})
	errorHandler.ErrorHandler(err, "Error creating VirtualService")
	return nil
}

func (k *KubernetesHandler) DeleteDeployment() error {

	// Delete the Deployment
	err := k.Clientset.AppsV1().Deployments(k.Namespace).Delete(context.TODO(), k.Name, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete Deployment: %v", err)
	}

	return nil
}

func (k *KubernetesHandler) DeleteService() error {

	// Delete the Service
	err := k.Clientset.CoreV1().Services(k.Namespace).Delete(context.TODO(), k.Name, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete Service: %v", err)
	}

	return nil
}

func (k *KubernetesHandler) DeleteVirtualService() error {

	// Delete the VirtualService
	err := k.DynamicClient.Resource(customresource.VirtualService).Namespace(k.Namespace).Delete(context.TODO(), k.Name, metav1.DeleteOptions{})
	errorHandler.ErrorHandler(err, "Error deleting VirtualService")
	return nil
}

func (k *KubernetesHandler) GetDeployment() (*appsv1.Deployment, error) {

	// Get the Deployment
	deployment, err := k.Clientset.AppsV1().Deployments(k.Namespace).Get(context.TODO(), k.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get Deployment: %v", err)
	}

	return deployment, nil
}

func (k *KubernetesHandler) GetService() (*corev1.Service, error) {

	// Get the Service
	service, err := k.Clientset.CoreV1().Services(k.Namespace).Get(context.TODO(), k.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get Service: %v", err)
	}

	return service, nil
}

func (k *KubernetesHandler) GetVirtualService() (*unstructured.Unstructured, error) {

	// Get the VirtualService
	virtualService, err := k.DynamicClient.Resource(customresource.VirtualService).Namespace(k.Namespace).Get(context.TODO(), k.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get VirtualService: %v", err)
	}

	return virtualService, nil
}

func (k *KubernetesHandler) GetDeploymentLogs() (string, error) {

	// Get the logs of the first container in the Deployment
	podList, err := k.Clientset.CoreV1().Pods(k.Namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s", k.Name),
	})
	if err != nil {
		return "", fmt.Errorf("failed to get Pod list: %v", err)
	}

	if len(podList.Items) == 0 {
		return "", fmt.Errorf("no Pods found")
	}

	var logs string
	for _, pod := range podList.Items {
		podLogOpts := corev1.PodLogOptions{}
		req := k.Clientset.CoreV1().Pods(k.Namespace).GetLogs(pod.Name, &podLogOpts)
		podLogs, err := req.Stream(context.Background())
		if err != nil {
			return "", fmt.Errorf("failed to get Pod logs: %v", err)
		}
		defer podLogs.Close()
		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, podLogs)
		if err != nil {
			return "", fmt.Errorf("failed to copy Pod logs: %v", err)
		}
		logs += buf.String()
	}

	return logs, nil
}

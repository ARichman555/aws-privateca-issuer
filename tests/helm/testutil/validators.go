package testutil

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ResourceNames provides common resource name patterns
type ResourceNames struct {
	Release string
}

func (r ResourceNames) Deployment() string {
	return r.Release + "-aws-privateca-issuer"
}

func (r ResourceNames) Service() string {
	return r.Release + "-aws-privateca-issuer"
}

func (r ResourceNames) ServiceAccount() string {
	return r.Release + "-aws-privateca-issuer"
}

func (r ResourceNames) ClusterRole() string {
	return r.Release + "-aws-privateca-issuer"
}

func (r ResourceNames) ClusterRoleBinding() string {
	return r.Release + "-aws-privateca-issuer"
}

// Common validation helpers
func (h *TestHelper) GetDeployment(name string) *appsv1.Deployment {
	deployment, err := h.Clientset.AppsV1().Deployments(h.Namespace).Get(context.TODO(), name, metav1.GetOptions{})
	require.NoError(h.T, err)
	return deployment
}

func (h *TestHelper) GetService(name string) *corev1.Service {
	service, err := h.Clientset.CoreV1().Services(h.Namespace).Get(context.TODO(), name, metav1.GetOptions{})
	require.NoError(h.T, err)
	return service
}

func (h *TestHelper) GetServiceAccount(name string) *corev1.ServiceAccount {
	sa, err := h.Clientset.CoreV1().ServiceAccounts(h.Namespace).Get(context.TODO(), name, metav1.GetOptions{})
	require.NoError(h.T, err)
	return sa
}

// Validation helpers
func ValidateDeploymentReplicas(t *testing.T, deployment *appsv1.Deployment, expectedReplicas int32) {
	assert.Equal(t, expectedReplicas, *deployment.Spec.Replicas)
}

func ValidateDeploymentImage(t *testing.T, deployment *appsv1.Deployment, expectedImage string) {
	container := deployment.Spec.Template.Spec.Containers[0]
	assert.Contains(t, container.Image, expectedImage)
}

func ValidateDeploymentArgs(t *testing.T, deployment *appsv1.Deployment, expectedArg string) {
	container := deployment.Spec.Template.Spec.Containers[0]
	assert.Contains(t, container.Args, expectedArg)
}

func ValidateServiceType(t *testing.T, service *corev1.Service, expectedType corev1.ServiceType) {
	assert.Equal(t, expectedType, service.Spec.Type)
}

func ValidateServicePort(t *testing.T, service *corev1.Service, expectedPort int32) {
	assert.Equal(t, expectedPort, service.Spec.Ports[0].Port)
}

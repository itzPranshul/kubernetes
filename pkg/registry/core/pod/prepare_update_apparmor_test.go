package pod

import (
    "context"
    "testing"

    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/kubernetes/pkg/kubelet/apis/apparmor" // correct import for prefix
)

func TestPrepareForUpdate_PopulatesAppArmorFromAnnotation(t *testing.T) {
    ctx := context.TODO()

    // Old pod has only the AppArmor annotation (old format)
    oldPod := &corev1.Pod{
        ObjectMeta: metav1.ObjectMeta{
            Name:      "test-pod",
            Namespace: "default",
            Annotations: map[string]string{
                apparmor.ContainerAnnotationKeyPrefix + "test-container": "localhost/test-profile",
            },
        },
        Spec: corev1.PodSpec{
            Containers: []corev1.Container{
                {Name: "test-container"},
            },
        },
    }

    // New pod copy (simulating an update)
    newPod := oldPod.DeepCopy()

    // Run PrepareForUpdate
    podStrategy{}.PrepareForUpdate(ctx, newPod, oldPod)

    sc := newPod.Spec.Containers[0].SecurityContext
    if sc == nil || sc.AppArmorProfile == nil {
        t.Fatalf("expected AppArmorProfile to be set from old annotation, but got nil")
    }

    if sc.AppArmorProfile.Type != corev1.AppArmorProfileTypeLocalhost {
        t.Errorf("expected AppArmorProfile.Type=Localhost, got %q", sc.AppArmorProfile.Type)
    }
    if sc.AppArmorProfile.LocalhostProfile == nil || *sc.AppArmorProfile.LocalhostProfile != "test-profile" {
        t.Errorf("expected LocalhostProfile 'test-profile', got %v", sc.AppArmorProfile.LocalhostProfile)
    }
}

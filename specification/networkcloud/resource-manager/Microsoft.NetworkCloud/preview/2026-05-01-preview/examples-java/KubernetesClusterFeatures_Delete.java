
/**
 * Samples for KubernetesClusterFeatures Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/KubernetesClusterFeatures_Delete.json
     */
    /**
     * Sample code: Delete Kubernetes cluster feature.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        deleteKubernetesClusterFeature(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.kubernetesClusterFeatures().delete("resourceGroupName", "kubernetesClusterName", "featureName", null,
            null, com.azure.core.util.Context.NONE);
    }
}

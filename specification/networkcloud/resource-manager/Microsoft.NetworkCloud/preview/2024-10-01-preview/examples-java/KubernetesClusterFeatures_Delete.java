
/**
 * Samples for KubernetesClusterFeatures Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-10-01-preview/examples/
     * KubernetesClusterFeatures_Delete.json
     */
    /**
     * Sample code: Delete Kubernetes cluster feature.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        deleteKubernetesClusterFeature(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.kubernetesClusterFeatures().delete("resourceGroupName", "kubernetesClusterName", "featureName",
            com.azure.core.util.Context.NONE);
    }
}

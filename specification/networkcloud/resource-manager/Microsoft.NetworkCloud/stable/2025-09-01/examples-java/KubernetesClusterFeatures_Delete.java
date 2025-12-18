
/**
 * Samples for KubernetesClusterFeatures Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * KubernetesClusterFeatures_Delete.json
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

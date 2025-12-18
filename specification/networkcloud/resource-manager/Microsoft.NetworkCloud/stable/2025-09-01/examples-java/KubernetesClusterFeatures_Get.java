
/**
 * Samples for KubernetesClusterFeatures Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * KubernetesClusterFeatures_Get.json
     */
    /**
     * Sample code: Get Kubernetes cluster feature.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getKubernetesClusterFeature(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.kubernetesClusterFeatures().getWithResponse("resourceGroupName", "kubernetesClusterName", "featureName",
            com.azure.core.util.Context.NONE);
    }
}

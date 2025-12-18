
/**
 * Samples for KubernetesClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * KubernetesClusters_Get.json
     */
    /**
     * Sample code: Get Kubernetes cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getKubernetesCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.kubernetesClusters().getByResourceGroupWithResponse("resourceGroupName", "kubernetesClusterName",
            com.azure.core.util.Context.NONE);
    }
}

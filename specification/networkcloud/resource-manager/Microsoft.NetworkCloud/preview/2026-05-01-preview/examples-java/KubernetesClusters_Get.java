
/**
 * Samples for KubernetesClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/KubernetesClusters_Get.json
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

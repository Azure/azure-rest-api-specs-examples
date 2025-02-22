
/**
 * Samples for KubernetesClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2024-07-01/examples/
     * KubernetesClusters_Delete.json
     */
    /**
     * Sample code: Delete Kubernetes cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteKubernetesCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.kubernetesClusters().delete("resourceGroupName", "kubernetesClusterName",
            com.azure.core.util.Context.NONE);
    }
}

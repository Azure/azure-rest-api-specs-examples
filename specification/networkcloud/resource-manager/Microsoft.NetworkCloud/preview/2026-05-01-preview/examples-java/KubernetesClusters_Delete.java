
/**
 * Samples for KubernetesClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/KubernetesClusters_Delete.json
     */
    /**
     * Sample code: Delete Kubernetes cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteKubernetesCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.kubernetesClusters().delete("resourceGroupName", "kubernetesClusterName", null, null,
            com.azure.core.util.Context.NONE);
    }
}

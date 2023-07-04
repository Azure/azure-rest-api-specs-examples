/** Samples for KubernetesClusters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/KubernetesClusters_Delete.json
     */
    /**
     * Sample code: Delete Kubernetes cluster.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteKubernetesCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .kubernetesClusters()
            .delete("resourceGroupName", "kubernetesClusterName", com.azure.core.util.Context.NONE);
    }
}

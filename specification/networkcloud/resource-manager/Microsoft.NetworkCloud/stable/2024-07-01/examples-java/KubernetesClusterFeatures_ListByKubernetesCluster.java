
/**
 * Samples for KubernetesClusterFeatures ListByKubernetesCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2024-07-01/examples/
     * KubernetesClusterFeatures_ListByKubernetesCluster.json
     */
    /**
     * Sample code: List features for the Kubernetes cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listFeaturesForTheKubernetesCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.kubernetesClusterFeatures().listByKubernetesCluster("resourceGroupName", "kubernetesClusterName",
            com.azure.core.util.Context.NONE);
    }
}

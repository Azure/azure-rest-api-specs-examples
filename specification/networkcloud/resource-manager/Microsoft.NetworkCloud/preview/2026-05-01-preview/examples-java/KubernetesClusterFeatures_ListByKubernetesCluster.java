
/**
 * Samples for KubernetesClusterFeatures ListByKubernetesCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/KubernetesClusterFeatures_ListByKubernetesCluster.json
     */
    /**
     * Sample code: List features for the Kubernetes cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listFeaturesForTheKubernetesCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.kubernetesClusterFeatures().listByKubernetesCluster("resourceGroupName", "kubernetesClusterName", null,
            null, com.azure.core.util.Context.NONE);
    }
}

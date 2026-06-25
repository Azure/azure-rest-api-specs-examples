
/**
 * Samples for KubernetesClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/KubernetesClusters_ListBySubscription.json
     */
    /**
     * Sample code: List Kubernetes clusters for subscription.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listKubernetesClustersForSubscription(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.kubernetesClusters().list(null, null, com.azure.core.util.Context.NONE);
    }
}

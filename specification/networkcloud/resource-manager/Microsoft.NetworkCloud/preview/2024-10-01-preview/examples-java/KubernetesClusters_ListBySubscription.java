
/**
 * Samples for KubernetesClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-10-01-preview/examples/
     * KubernetesClusters_ListBySubscription.json
     */
    /**
     * Sample code: List Kubernetes clusters for subscription.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listKubernetesClustersForSubscription(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.kubernetesClusters().list(com.azure.core.util.Context.NONE);
    }
}

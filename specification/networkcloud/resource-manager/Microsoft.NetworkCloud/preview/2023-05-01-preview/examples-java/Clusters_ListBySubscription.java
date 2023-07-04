/** Samples for Clusters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/Clusters_ListBySubscription.json
     */
    /**
     * Sample code: List clusters for subscription.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listClustersForSubscription(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.clusters().list(com.azure.core.util.Context.NONE);
    }
}

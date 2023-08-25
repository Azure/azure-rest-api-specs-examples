/** Samples for ClusterManagers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2023-07-01/examples/ClusterManagers_ListBySubscription.json
     */
    /**
     * Sample code: List cluster managers for subscription.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listClusterManagersForSubscription(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.clusterManagers().list(com.azure.core.util.Context.NONE);
    }
}

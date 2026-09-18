
/**
 * Samples for ClusterManagers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ClusterManagers_ListBySubscription.json
     */
    /**
     * Sample code: List cluster managers for subscription.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listClusterManagersForSubscription(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.clusterManagers().list(null, null, com.azure.core.util.Context.NONE);
    }
}

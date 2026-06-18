
/**
 * Samples for NetworkMonitors List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkMonitors_ListBySubscription.json
     */
    /**
     * Sample code: NetworkMonitors_ListBySubscription.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkMonitorsListBySubscription(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkMonitors().list(com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for NetworkMonitors ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkMonitors_ListByResourceGroup.json
     */
    /**
     * Sample code: NetworkMonitors_ListByResourceGroup.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkMonitorsListByResourceGroup(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkMonitors().listByResourceGroup("example-rg", com.azure.core.util.Context.NONE);
    }
}

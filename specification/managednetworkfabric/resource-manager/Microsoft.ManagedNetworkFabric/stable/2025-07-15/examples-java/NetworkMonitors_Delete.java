
/**
 * Samples for NetworkMonitors Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkMonitors_Delete.json
     */
    /**
     * Sample code: NetworkMonitors_Delete_MaximumSet.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkMonitorsDeleteMaximumSet(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkMonitors().delete("rgmanagednetworkfabric", "example-monitor", com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for NetworkMonitors GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkMonitors_Get.json
     */
    /**
     * Sample code: NetworkMonitors_Get_MaximumSet.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkMonitorsGetMaximumSet(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkMonitors().getByResourceGroupWithResponse("example-rg", "example-monitor",
            com.azure.core.util.Context.NONE);
    }
}

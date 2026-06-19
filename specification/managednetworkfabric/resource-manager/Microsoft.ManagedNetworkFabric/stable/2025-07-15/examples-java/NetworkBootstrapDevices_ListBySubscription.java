
/**
 * Samples for NetworkBootstrapDevices List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkBootstrapDevices_ListBySubscription.json
     */
    /**
     * Sample code: NetworkBootstrapDevices_ListBySubscription.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkBootstrapDevicesListBySubscription(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkBootstrapDevices().list(com.azure.core.util.Context.NONE);
    }
}

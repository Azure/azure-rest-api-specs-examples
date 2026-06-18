
/**
 * Samples for NetworkDevices List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkDevices_ListBySubscription.json
     */
    /**
     * Sample code: NetworkDevices_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkDevicesListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkDevices().list(com.azure.core.util.Context.NONE);
    }
}

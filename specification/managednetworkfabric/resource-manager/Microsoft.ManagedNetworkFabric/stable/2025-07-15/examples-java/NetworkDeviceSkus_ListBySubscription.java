
/**
 * Samples for NetworkDeviceSkus List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkDeviceSkus_ListBySubscription.json
     */
    /**
     * Sample code: NetworkDeviceSkus_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkDeviceSkusListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkDeviceSkus().list(com.azure.core.util.Context.NONE);
    }
}

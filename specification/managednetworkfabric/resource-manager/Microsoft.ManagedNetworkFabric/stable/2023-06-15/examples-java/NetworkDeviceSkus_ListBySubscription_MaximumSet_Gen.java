
/**
 * Samples for NetworkDeviceSkus List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/
     * NetworkDeviceSkus_ListBySubscription_MaximumSet_Gen.json
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

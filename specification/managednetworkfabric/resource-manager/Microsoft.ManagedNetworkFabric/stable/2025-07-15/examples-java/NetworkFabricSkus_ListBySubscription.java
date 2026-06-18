
/**
 * Samples for NetworkFabricSkus List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabricSkus_ListBySubscription.json
     */
    /**
     * Sample code: NetworkFabricSkus_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricSkusListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabricSkus().list(com.azure.core.util.Context.NONE);
    }
}

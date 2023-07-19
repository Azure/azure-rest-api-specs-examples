/** Samples for NetworkFabricSkus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabricSkus_ListBySubscription_MaximumSet_Gen.json
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

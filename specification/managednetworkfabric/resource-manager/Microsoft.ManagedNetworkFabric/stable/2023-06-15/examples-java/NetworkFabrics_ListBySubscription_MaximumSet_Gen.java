
/**
 * Samples for NetworkFabrics List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/
     * NetworkFabrics_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkFabrics_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().list(com.azure.core.util.Context.NONE);
    }
}

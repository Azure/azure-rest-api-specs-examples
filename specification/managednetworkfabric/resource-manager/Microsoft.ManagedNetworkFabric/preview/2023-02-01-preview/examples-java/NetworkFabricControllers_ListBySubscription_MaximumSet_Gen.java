/** Samples for NetworkFabricControllers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkFabricControllers_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkFabricControllers_ListBySubscription_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricControllersListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabricControllers().list(com.azure.core.util.Context.NONE);
    }
}

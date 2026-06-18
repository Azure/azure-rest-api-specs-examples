
/**
 * Samples for NetworkFabricControllers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabricControllers_ListBySubscription.json
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

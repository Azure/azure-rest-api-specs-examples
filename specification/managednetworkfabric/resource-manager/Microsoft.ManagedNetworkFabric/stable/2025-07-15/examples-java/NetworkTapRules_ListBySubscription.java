
/**
 * Samples for NetworkTapRules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkTapRules_ListBySubscription.json
     */
    /**
     * Sample code: NetworkTapRules_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkTapRulesListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkTapRules().list(com.azure.core.util.Context.NONE);
    }
}

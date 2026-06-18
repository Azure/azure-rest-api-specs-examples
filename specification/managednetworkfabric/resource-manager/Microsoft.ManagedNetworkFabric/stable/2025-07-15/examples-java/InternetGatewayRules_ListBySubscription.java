
/**
 * Samples for InternetGatewayRules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/InternetGatewayRules_ListBySubscription.json
     */
    /**
     * Sample code: InternetGatewayRules_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internetGatewayRulesListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.internetGatewayRules().list(com.azure.core.util.Context.NONE);
    }
}

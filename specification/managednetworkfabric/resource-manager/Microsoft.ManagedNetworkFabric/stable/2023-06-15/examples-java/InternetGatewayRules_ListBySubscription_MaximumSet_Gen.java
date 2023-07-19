/** Samples for InternetGatewayRules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/InternetGatewayRules_ListBySubscription_MaximumSet_Gen.json
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

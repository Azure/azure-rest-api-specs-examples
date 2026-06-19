
/**
 * Samples for InternetGatewayRules ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/InternetGatewayRules_ListByResourceGroup.json
     */
    /**
     * Sample code: InternetGatewayRules_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internetGatewayRulesListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.internetGatewayRules().listByResourceGroup("example-internetGatewayRule",
            com.azure.core.util.Context.NONE);
    }
}

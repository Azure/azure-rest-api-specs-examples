/** Samples for InternetGatewayRules ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/InternetGatewayRules_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: InternetGatewayRules_ListByResourceGroup_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internetGatewayRulesListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .internetGatewayRules()
            .listByResourceGroup("example-internetGatewayRule", com.azure.core.util.Context.NONE);
    }
}

/** Samples for InternetGatewayRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/InternetGatewayRules_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: InternetGatewayRules_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internetGatewayRulesDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .internetGatewayRules()
            .delete("example-rg", "example-internetGatewayRule", com.azure.core.util.Context.NONE);
    }
}

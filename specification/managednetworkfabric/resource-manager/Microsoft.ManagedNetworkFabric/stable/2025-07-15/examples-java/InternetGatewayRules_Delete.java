
/**
 * Samples for InternetGatewayRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/InternetGatewayRules_Delete.json
     */
    /**
     * Sample code: InternetGatewayRules_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internetGatewayRulesDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.internetGatewayRules().delete("example-rg", "example-internetGatewayRule",
            com.azure.core.util.Context.NONE);
    }
}

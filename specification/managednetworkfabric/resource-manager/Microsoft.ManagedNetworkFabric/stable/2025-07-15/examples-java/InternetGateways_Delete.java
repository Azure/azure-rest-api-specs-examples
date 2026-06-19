
/**
 * Samples for InternetGateways Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/InternetGateways_Delete.json
     */
    /**
     * Sample code: InternetGateways_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internetGatewaysDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.internetGateways().delete("example-rg", "example-internetGateway", com.azure.core.util.Context.NONE);
    }
}

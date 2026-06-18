
/**
 * Samples for InternetGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/InternetGateways_Get.json
     */
    /**
     * Sample code: InternetGateways_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internetGatewaysGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.internetGateways().getByResourceGroupWithResponse("example-rg", "example-internetGateway",
            com.azure.core.util.Context.NONE);
    }
}

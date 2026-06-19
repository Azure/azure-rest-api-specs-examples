
/**
 * Samples for InternetGateways ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/InternetGateways_ListByResourceGroup.json
     */
    /**
     * Sample code: InternetGateways_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internetGatewaysListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.internetGateways().listByResourceGroup("example-rg", com.azure.core.util.Context.NONE);
    }
}

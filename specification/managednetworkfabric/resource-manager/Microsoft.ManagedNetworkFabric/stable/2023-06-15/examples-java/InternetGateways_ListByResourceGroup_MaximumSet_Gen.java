
/**
 * Samples for InternetGateways ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/
     * InternetGateways_ListByResourceGroup_MaximumSet_Gen.json
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

/** Samples for InternetGateways GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/InternetGateways_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: InternetGateways_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internetGatewaysGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .internetGateways()
            .getByResourceGroupWithResponse("example-rg", "example-internetGateway", com.azure.core.util.Context.NONE);
    }
}

/** Samples for NetworkFabrics GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabrics_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkFabrics_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkFabrics()
            .getByResourceGroupWithResponse("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}

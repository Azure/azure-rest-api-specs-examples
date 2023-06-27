/** Samples for NetworkFabrics GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkFabrics_Get_MaximumSet_Gen.json
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
            .getByResourceGroupWithResponse("resourceGroupName", "FabricName", com.azure.core.util.Context.NONE);
    }
}

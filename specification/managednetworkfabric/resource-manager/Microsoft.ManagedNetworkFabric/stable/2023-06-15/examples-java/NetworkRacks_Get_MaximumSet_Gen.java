/** Samples for NetworkRacks GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkRacks_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkRacks_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkRacksGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkRacks()
            .getByResourceGroupWithResponse("example-rg", "example-rack", com.azure.core.util.Context.NONE);
    }
}

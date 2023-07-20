/** Samples for NetworkTaps GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkTaps_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkTaps_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkTapsGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkTaps()
            .getByResourceGroupWithResponse("example-rg", "example-networkTap", com.azure.core.util.Context.NONE);
    }
}

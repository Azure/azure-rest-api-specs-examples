/** Samples for NetworkToNetworkInterconnects ListByNetworkFabric. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkToNetworkInterconnects_ListByNetworkFabric_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkToNetworkInterconnects_ListByNetworkFabric_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkToNetworkInterconnectsListByNetworkFabricMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkToNetworkInterconnects()
            .listByNetworkFabric("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}

/** Samples for NetworkToNetworkInterconnects Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkToNetworkInterconnects_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkToNetworkInterconnects_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkToNetworkInterconnectsDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkToNetworkInterconnects()
            .delete("example-rg", "example-fabric", "example-nni", com.azure.core.util.Context.NONE);
    }
}

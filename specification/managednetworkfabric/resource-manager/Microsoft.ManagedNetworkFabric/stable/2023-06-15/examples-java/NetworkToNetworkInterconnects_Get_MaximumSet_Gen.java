/** Samples for NetworkToNetworkInterconnects Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkToNetworkInterconnects_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkToNetworkInterconnects_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkToNetworkInterconnectsGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkToNetworkInterconnects()
            .getWithResponse("example-rg", "example-fabric", "example-nni", com.azure.core.util.Context.NONE);
    }
}

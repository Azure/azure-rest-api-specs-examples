/** Samples for NetworkToNetworkInterconnects List. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkToNetworkInterconnects_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkToNetworkInterconnects_List_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkToNetworkInterconnectsListMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkToNetworkInterconnects()
            .list("resourceGroupName", "FabricName", com.azure.core.util.Context.NONE);
    }
}

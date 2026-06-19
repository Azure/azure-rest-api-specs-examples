
/**
 * Samples for NetworkToNetworkInterconnects ListByNetworkFabric.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkToNetworkInterconnects_ListByNetworkFabric.json
     */
    /**
     * Sample code: NetworkToNetworkInterconnects_ListByNetworkFabric_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkToNetworkInterconnectsListByNetworkFabricMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkToNetworkInterconnects().listByNetworkFabric("example-rg", "example-nf",
            com.azure.core.util.Context.NONE);
    }
}

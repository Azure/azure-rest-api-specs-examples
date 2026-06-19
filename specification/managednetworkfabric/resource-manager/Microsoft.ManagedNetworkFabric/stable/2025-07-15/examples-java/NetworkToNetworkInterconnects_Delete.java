
/**
 * Samples for NetworkToNetworkInterconnects Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkToNetworkInterconnects_Delete.json
     */
    /**
     * Sample code: NetworkToNetworkInterconnects_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkToNetworkInterconnectsDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkToNetworkInterconnects().delete("example-rg", "example-nf", "example-nni",
            com.azure.core.util.Context.NONE);
    }
}

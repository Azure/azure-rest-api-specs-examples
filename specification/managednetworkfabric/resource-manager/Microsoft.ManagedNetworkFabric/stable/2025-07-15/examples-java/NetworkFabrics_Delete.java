
/**
 * Samples for NetworkFabrics Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_Delete.json
     */
    /**
     * Sample code: NetworkFabrics_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().delete("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}

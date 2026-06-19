
/**
 * Samples for NetworkFabrics CommitConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_CommitConfiguration.json
     */
    /**
     * Sample code: NetworkFabrics_CommitConfiguration_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsCommitConfigurationMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().commitConfiguration("example-rg", "example-networkFabric", null,
            com.azure.core.util.Context.NONE);
    }
}

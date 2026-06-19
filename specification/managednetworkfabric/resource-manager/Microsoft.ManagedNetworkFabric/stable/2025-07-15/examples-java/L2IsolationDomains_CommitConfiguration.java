
/**
 * Samples for L2IsolationDomains CommitConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/L2IsolationDomains_CommitConfiguration.json
     */
    /**
     * Sample code: L2IsolationDomains_CommitConfiguration_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void l2IsolationDomainsCommitConfigurationMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.l2IsolationDomains().commitConfiguration("example-rg", "example-l2domain",
            com.azure.core.util.Context.NONE);
    }
}

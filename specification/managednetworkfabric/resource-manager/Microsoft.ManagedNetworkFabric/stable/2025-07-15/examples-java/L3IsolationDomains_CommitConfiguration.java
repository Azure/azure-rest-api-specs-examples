
/**
 * Samples for L3IsolationDomains CommitConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/L3IsolationDomains_CommitConfiguration.json
     */
    /**
     * Sample code: L3IsolationDomains_CommitConfiguration_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void l3IsolationDomainsCommitConfigurationMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.l3IsolationDomains().commitConfiguration("example-rg", "example-l3domain",
            com.azure.core.util.Context.NONE);
    }
}

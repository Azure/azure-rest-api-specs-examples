
/**
 * Samples for L3IsolationDomains ValidateConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/L3IsolationDomains_ValidateConfiguration.json
     */
    /**
     * Sample code: L3IsolationDomains_ValidateConfiguration_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void l3IsolationDomainsValidateConfigurationMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.l3IsolationDomains().validateConfiguration("example-rg", "example-l3domain",
            com.azure.core.util.Context.NONE);
    }
}

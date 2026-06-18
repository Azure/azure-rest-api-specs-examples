
/**
 * Samples for L3IsolationDomains Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/L3IsolationDomains_Delete.json
     */
    /**
     * Sample code: L3IsolationDomains_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void l3IsolationDomainsDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.l3IsolationDomains().delete("example-rg", "example-l3domain", com.azure.core.util.Context.NONE);
    }
}

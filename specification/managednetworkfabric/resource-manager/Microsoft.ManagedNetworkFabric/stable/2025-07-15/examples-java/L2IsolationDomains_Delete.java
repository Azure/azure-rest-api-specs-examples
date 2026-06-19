
/**
 * Samples for L2IsolationDomains Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/L2IsolationDomains_Delete.json
     */
    /**
     * Sample code: L2IsolationDomains_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void l2IsolationDomainsDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.l2IsolationDomains().delete("example-rg", "example-l2domain", com.azure.core.util.Context.NONE);
    }
}

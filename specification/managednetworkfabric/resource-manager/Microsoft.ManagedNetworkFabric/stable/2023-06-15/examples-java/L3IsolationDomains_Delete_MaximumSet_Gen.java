
/**
 * Samples for L3IsolationDomains Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/
     * L3IsolationDomains_Delete_MaximumSet_Gen.json
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

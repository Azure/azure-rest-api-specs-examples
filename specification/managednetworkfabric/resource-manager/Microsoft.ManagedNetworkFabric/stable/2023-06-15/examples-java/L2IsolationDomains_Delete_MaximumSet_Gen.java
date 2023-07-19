/** Samples for L2IsolationDomains Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/L2IsolationDomains_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: L2IsolationDomains_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void l2IsolationDomainsDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.l2IsolationDomains().delete("example-rg", "example-l2Domain", com.azure.core.util.Context.NONE);
    }
}

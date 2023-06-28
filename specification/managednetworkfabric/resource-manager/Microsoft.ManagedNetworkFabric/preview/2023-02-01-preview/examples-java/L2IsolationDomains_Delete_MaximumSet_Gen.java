/** Samples for L2IsolationDomains Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/L2IsolationDomains_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: L2IsolationDomains_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void l2IsolationDomainsDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.l2IsolationDomains().delete("resourceGroupName", "example-l2domain", com.azure.core.util.Context.NONE);
    }
}

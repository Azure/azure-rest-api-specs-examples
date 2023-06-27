/** Samples for L3IsolationDomains GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/L3IsolationDomains_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: L3IsolationDomains_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void l3IsolationDomainsGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .l3IsolationDomains()
            .getByResourceGroupWithResponse("resourceGroupName", "example-l3domain", com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for L3IsolationDomains GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/L3IsolationDomains_Get.json
     */
    /**
     * Sample code: L3IsolationDomains_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void l3IsolationDomainsGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.l3IsolationDomains().getByResourceGroupWithResponse("example-rg", "example-l3domain",
            com.azure.core.util.Context.NONE);
    }
}

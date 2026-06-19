
/**
 * Samples for L2IsolationDomains GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/L2IsolationDomains_Get.json
     */
    /**
     * Sample code: L2IsolationDomains_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void l2IsolationDomainsGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.l2IsolationDomains().getByResourceGroupWithResponse("example-rg", "example-l2domain",
            com.azure.core.util.Context.NONE);
    }
}

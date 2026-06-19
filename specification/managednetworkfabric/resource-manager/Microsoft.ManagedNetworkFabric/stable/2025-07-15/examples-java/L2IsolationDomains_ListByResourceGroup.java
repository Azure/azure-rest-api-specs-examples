
/**
 * Samples for L2IsolationDomains ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/L2IsolationDomains_ListByResourceGroup.json
     */
    /**
     * Sample code: L2IsolationDomains_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void l2IsolationDomainsListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.l2IsolationDomains().listByResourceGroup("example-rg", com.azure.core.util.Context.NONE);
    }
}

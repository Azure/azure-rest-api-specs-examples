
/**
 * Samples for L3IsolationDomains ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/L3IsolationDomains_ListByResourceGroup.json
     */
    /**
     * Sample code: L3IsolationDomains_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void l3IsolationDomainsListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.l3IsolationDomains().listByResourceGroup("example-rg", com.azure.core.util.Context.NONE);
    }
}

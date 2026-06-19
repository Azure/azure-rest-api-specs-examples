
/**
 * Samples for L2IsolationDomains List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/L2IsolationDomains_ListBySubscription.json
     */
    /**
     * Sample code: L2IsolationDomains_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void l2IsolationDomainsListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.l2IsolationDomains().list(com.azure.core.util.Context.NONE);
    }
}

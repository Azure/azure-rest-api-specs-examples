
/**
 * Samples for L3IsolationDomains List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/L3IsolationDomains_ListBySubscription.json
     */
    /**
     * Sample code: L3IsolationDomains_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void l3IsolationDomainsListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.l3IsolationDomains().list(com.azure.core.util.Context.NONE);
    }
}

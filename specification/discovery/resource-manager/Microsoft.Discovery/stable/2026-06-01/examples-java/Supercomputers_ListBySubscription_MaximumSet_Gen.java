
/**
 * Samples for Supercomputers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Supercomputers_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Supercomputers_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void
        supercomputersListBySubscriptionMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.supercomputers().list(com.azure.core.util.Context.NONE);
    }
}

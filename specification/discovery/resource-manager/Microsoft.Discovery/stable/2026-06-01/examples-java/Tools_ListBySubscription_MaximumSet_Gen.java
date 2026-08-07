
/**
 * Samples for Tools List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Tools_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Tools_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void toolsListBySubscriptionMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.tools().list(com.azure.core.util.Context.NONE);
    }
}

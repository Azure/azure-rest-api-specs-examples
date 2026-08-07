
/**
 * Samples for Workspaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Workspaces_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void
        workspacesListBySubscriptionMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.workspaces().list(com.azure.core.util.Context.NONE);
    }
}

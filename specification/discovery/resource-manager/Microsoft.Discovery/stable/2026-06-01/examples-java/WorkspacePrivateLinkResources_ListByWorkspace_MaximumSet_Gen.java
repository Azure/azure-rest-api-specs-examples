
/**
 * Samples for WorkspacePrivateLinkResources ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/WorkspacePrivateLinkResources_ListByWorkspace_MaximumSet_Gen.json
     */
    /**
     * Sample code: WorkspacePrivateLinkResources_ListByWorkspace_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void workspacePrivateLinkResourcesListByWorkspaceMaximumSet(
        com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.workspacePrivateLinkResources().listByWorkspace("rgdiscovery", "9aa6a22ca481a4fa4e",
            com.azure.core.util.Context.NONE);
    }
}

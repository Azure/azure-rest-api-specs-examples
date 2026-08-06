
/**
 * Samples for WorkspacePrivateEndpointConnections ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/WorkspacePrivateEndpointConnections_ListByWorkspace_MaximumSet_Gen.json
     */
    /**
     * Sample code: WorkspacePrivateEndpointConnections_ListByWorkspace_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void workspacePrivateEndpointConnectionsListByWorkspaceMaximumSet(
        com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.workspacePrivateEndpointConnections().listByWorkspace("rgdiscovery", "704cee821b47e58afe",
            com.azure.core.util.Context.NONE);
    }
}

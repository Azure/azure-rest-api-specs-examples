
/**
 * Samples for WorkspacePrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/WorkspacePrivateEndpointConnections_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: WorkspacePrivateEndpointConnections_Delete_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void workspacePrivateEndpointConnectionsDeleteMaximumSet(
        com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.workspacePrivateEndpointConnections().delete("rgdiscovery", "2602de8dc5723c9502", "connection",
            com.azure.core.util.Context.NONE);
    }
}

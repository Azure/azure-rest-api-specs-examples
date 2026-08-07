
/**
 * Samples for WorkspacePrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/WorkspacePrivateEndpointConnections_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: WorkspacePrivateEndpointConnections_Get_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void
        workspacePrivateEndpointConnectionsGetMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.workspacePrivateEndpointConnections().getWithResponse("rgdiscovery", "de936b8038cf3dc9ad", "connection",
            com.azure.core.util.Context.NONE);
    }
}

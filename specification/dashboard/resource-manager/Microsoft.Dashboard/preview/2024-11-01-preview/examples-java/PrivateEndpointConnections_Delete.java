
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/PrivateEndpointConnections_Delete.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Delete.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void privateEndpointConnectionsDelete(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.privateEndpointConnections().delete("myResourceGroup", "myWorkspace", "myConnection",
            com.azure.core.util.Context.NONE);
    }
}

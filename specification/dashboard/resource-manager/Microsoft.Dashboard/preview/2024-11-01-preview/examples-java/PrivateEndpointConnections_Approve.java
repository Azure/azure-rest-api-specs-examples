
/**
 * Samples for PrivateEndpointConnections Approve.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/PrivateEndpointConnections_Approve.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Approve.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void privateEndpointConnectionsApprove(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.privateEndpointConnections().define("myConnection")
            .withExistingGrafana("myResourceGroup", "myWorkspace").create();
    }
}

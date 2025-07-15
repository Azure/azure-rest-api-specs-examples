
/**
 * Samples for PrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/PrivateEndpointConnections_List.json
     */
    /**
     * Sample code: PrivateEndpointConnections_List.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void privateEndpointConnectionsList(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.privateEndpointConnections().list("myResourceGroup", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}

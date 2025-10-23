
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/PrivateEndpointConnections_Get.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Get.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void privateEndpointConnectionsGet(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.privateEndpointConnections().getWithResponse("myResourceGroup", "myWorkspace", "myConnection",
            com.azure.core.util.Context.NONE);
    }
}

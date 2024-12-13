
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2023-09-01/examples/
     * PrivateEndpointConnections_Get.json
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

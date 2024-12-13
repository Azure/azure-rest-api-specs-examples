
/**
 * Samples for ManagedPrivateEndpoints Refresh.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2023-09-01/examples/
     * ManagedPrivateEndpoints_Refresh.json
     */
    /**
     * Sample code: ManagedPrivateEndpoint_Refresh.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void managedPrivateEndpointRefresh(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.managedPrivateEndpoints().refresh("myResourceGroup", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}

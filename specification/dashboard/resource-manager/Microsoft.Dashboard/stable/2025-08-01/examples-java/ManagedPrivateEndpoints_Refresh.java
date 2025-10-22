
/**
 * Samples for ManagedPrivateEndpoints Refresh.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/ManagedPrivateEndpoints_Refresh.json
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

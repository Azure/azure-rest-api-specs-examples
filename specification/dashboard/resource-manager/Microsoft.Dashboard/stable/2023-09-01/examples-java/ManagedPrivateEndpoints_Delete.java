
/**
 * Samples for ManagedPrivateEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2023-09-01/examples/
     * ManagedPrivateEndpoints_Delete.json
     */
    /**
     * Sample code: ManagedPrivateEndpoint_Delete.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void managedPrivateEndpointDelete(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.managedPrivateEndpoints().delete("myResourceGroup", "myWorkspace", "myMPEName",
            com.azure.core.util.Context.NONE);
    }
}

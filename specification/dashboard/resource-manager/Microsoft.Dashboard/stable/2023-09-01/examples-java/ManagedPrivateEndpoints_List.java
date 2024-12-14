
/**
 * Samples for ManagedPrivateEndpoints List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2023-09-01/examples/
     * ManagedPrivateEndpoints_List.json
     */
    /**
     * Sample code: ManagedPrivateEndpoint_List.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void managedPrivateEndpointList(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.managedPrivateEndpoints().list("myResourceGroup", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}

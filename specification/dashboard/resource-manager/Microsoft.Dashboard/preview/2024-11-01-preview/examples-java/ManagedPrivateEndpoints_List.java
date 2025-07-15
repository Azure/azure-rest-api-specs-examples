
/**
 * Samples for ManagedPrivateEndpoints List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/ManagedPrivateEndpoints_List.json
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


/**
 * Samples for ManagedPrivateEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/ManagedPrivateEndpoints_Get.json
     */
    /**
     * Sample code: ManagedPrivateEndpoint_Get.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void managedPrivateEndpointGet(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.managedPrivateEndpoints().getWithResponse("myResourceGroup", "myWorkspace", "myMPEName",
            com.azure.core.util.Context.NONE);
    }
}

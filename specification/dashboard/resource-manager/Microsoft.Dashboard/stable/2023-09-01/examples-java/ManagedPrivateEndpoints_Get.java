
/**
 * Samples for ManagedPrivateEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2023-09-01/examples/
     * ManagedPrivateEndpoints_Get.json
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


/**
 * Samples for PrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2023-09-01/examples/PrivateLinkResources_Get.
     * json
     */
    /**
     * Sample code: PrivateLinkResources_Get.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void privateLinkResourcesGet(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.privateLinkResources().getWithResponse("myResourceGroup", "myWorkspace", "grafana",
            com.azure.core.util.Context.NONE);
    }
}

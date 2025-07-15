
/**
 * Samples for PrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/PrivateLinkResources_Get.json
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

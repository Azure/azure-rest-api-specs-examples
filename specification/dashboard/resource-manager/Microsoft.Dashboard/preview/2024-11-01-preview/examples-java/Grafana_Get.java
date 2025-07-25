
/**
 * Samples for Grafana GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/Grafana_Get.json
     */
    /**
     * Sample code: Grafana_Get.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void grafanaGet(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.grafanas().getByResourceGroupWithResponse("myResourceGroup", "myWorkspace",
            com.azure.core.util.Context.NONE);
    }
}

import com.azure.core.util.Context;

/** Samples for Grafana GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/preview/2021-09-01-preview/examples/Grafana_Get.json
     */
    /**
     * Sample code: Grafana_Get.
     *
     * @param manager Entry point to DashboardManager.
     */
    public static void grafanaGet(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.grafanas().getByResourceGroupWithResponse("myResourceGroup", "myWorkspace", Context.NONE);
    }
}

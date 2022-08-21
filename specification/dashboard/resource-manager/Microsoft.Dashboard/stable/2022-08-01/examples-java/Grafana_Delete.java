import com.azure.core.util.Context;

/** Samples for Grafana Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2022-08-01/examples/Grafana_Delete.json
     */
    /**
     * Sample code: Grafana_Delete.
     *
     * @param manager Entry point to DashboardManager.
     */
    public static void grafanaDelete(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.grafanas().delete("myResourceGroup", "myWorkspace", Context.NONE);
    }
}

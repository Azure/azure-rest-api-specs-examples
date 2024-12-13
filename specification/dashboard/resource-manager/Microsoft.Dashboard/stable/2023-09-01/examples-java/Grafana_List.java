
/**
 * Samples for Grafana List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2023-09-01/examples/Grafana_List.json
     */
    /**
     * Sample code: Grafana_List.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void grafanaList(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.grafanas().list(com.azure.core.util.Context.NONE);
    }
}

import com.azure.core.util.Context;

/** Samples for Grafana List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/preview/2021-09-01-preview/examples/Grafana_List.json
     */
    /**
     * Sample code: Grafana_List.
     *
     * @param manager Entry point to DashboardManager.
     */
    public static void grafanaList(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.grafanas().list(Context.NONE);
    }
}

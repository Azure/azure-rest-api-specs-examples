
/**
 * Samples for Grafana List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/Grafana_List.json
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


/**
 * Samples for Grafana Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/Grafana_Delete.json
     */
    /**
     * Sample code: Grafana_Delete.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void grafanaDelete(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.grafanas().delete("myResourceGroup", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}

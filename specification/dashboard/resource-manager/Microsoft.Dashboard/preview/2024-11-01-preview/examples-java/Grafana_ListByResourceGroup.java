
/**
 * Samples for Grafana ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/Grafana_ListByResourceGroup.json
     */
    /**
     * Sample code: Grafana_ListByResourceGroup.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void grafanaListByResourceGroup(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.grafanas().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}

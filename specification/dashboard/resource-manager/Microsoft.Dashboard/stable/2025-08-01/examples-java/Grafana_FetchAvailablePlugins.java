
/**
 * Samples for Grafana FetchAvailablePlugins.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/Grafana_FetchAvailablePlugins.json
     */
    /**
     * Sample code: Grafana_FetchAvailablePlugins.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void grafanaFetchAvailablePlugins(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.grafanas().fetchAvailablePluginsWithResponse("myResourceGroup", "myWorkspace",
            com.azure.core.util.Context.NONE);
    }
}

/** Samples for Grafana FetchAvailablePlugins. */
public final class Main {
    /*
     * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2023-09-01/examples/
     * Grafana_FetchAvailablePlugins.json
     */
    /**
     * Sample code: Grafana_FetchAvailablePlugins.
     *
     * @param manager Entry point to DashboardManager.
     */
    public static void grafanaFetchAvailablePlugins(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager
            .grafanas()
            .fetchAvailablePluginsWithResponse("myResourceGroup", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}

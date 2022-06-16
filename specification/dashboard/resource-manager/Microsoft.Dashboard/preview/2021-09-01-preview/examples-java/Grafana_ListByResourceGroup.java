import com.azure.core.util.Context;

/** Samples for Grafana ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/preview/2021-09-01-preview/examples/Grafana_ListByResourceGroup.json
     */
    /**
     * Sample code: Grafana_ListByResourceGroup.
     *
     * @param manager Entry point to DashboardManager.
     */
    public static void grafanaListByResourceGroup(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.grafanas().listByResourceGroup("myResourceGroup", Context.NONE);
    }
}

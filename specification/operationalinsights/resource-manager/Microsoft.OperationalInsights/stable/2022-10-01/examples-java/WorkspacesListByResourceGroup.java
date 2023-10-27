/** Samples for Workspaces ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2022-10-01/examples/WorkspacesListByResourceGroup.json
     */
    /**
     * Sample code: WorkspacesGet.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void workspacesGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.workspaces().listByResourceGroup("oiautorest6685", com.azure.core.util.Context.NONE);
    }
}

import com.azure.core.util.Context;

/** Samples for Workspaces ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/WorkspacesListByResourceGroup.json
     */
    /**
     * Sample code: WorkspacesGet.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void workspacesGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.workspaces().listByResourceGroup("oiautorest6685", Context.NONE);
    }
}

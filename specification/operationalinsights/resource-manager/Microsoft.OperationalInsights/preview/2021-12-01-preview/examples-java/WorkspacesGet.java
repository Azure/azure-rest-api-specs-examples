import com.azure.core.util.Context;

/** Samples for Workspaces GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/WorkspacesGet.json
     */
    /**
     * Sample code: WorkspaceGet.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void workspaceGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.workspaces().getByResourceGroupWithResponse("oiautorest6685", "oiautorest6685", Context.NONE);
    }
}

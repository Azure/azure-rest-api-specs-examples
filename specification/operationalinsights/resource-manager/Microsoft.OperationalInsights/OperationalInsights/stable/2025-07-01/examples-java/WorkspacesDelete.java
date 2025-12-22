
/**
 * Samples for Workspaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/WorkspacesDelete.json
     */
    /**
     * Sample code: WorkspacesDelete.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void workspacesDelete(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.workspaces().delete("oiautorest6685", "oiautorest6685", null, com.azure.core.util.Context.NONE);
    }
}

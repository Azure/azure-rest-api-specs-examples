
/**
 * Samples for Workspaces Failover.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/WorkspacesFailover.json
     */
    /**
     * Sample code: WorkspacesFailover.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void workspacesFailover(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.workspaces().failover("oiautorest6685", "eastus", "oiautorest6685", com.azure.core.util.Context.NONE);
    }
}

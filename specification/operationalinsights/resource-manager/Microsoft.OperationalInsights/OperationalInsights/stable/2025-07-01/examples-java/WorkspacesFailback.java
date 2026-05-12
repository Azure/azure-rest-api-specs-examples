
/**
 * Samples for Workspaces Failback.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/WorkspacesFailback.json
     */
    /**
     * Sample code: WorkspacesFailover.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void workspacesFailover(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.workspaces().failback("oiautorest6685", "oiautorest6685", com.azure.core.util.Context.NONE);
    }
}

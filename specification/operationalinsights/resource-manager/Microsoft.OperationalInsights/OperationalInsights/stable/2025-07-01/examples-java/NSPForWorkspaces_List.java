
/**
 * Samples for Workspaces ListNsp.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/NSPForWorkspaces_List.json
     */
    /**
     * Sample code: List NSP configs by Scheduled Query Rule.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void
        listNSPConfigsByScheduledQueryRule(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.workspaces().listNsp("exampleRG", "someWorkspace", com.azure.core.util.Context.NONE);
    }
}

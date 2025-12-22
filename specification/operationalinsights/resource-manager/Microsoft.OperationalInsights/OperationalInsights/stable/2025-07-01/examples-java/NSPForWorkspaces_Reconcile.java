
/**
 * Samples for Workspaces ReconcileNsp.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/NSPForWorkspaces_Reconcile.json
     */
    /**
     * Sample code: Reconcile NSP config for Scheduled Query Rule.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void
        reconcileNSPConfigForScheduledQueryRule(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.workspaces().reconcileNsp("exampleRG", "someWorkspace", "somePerimeterConfiguration",
            com.azure.core.util.Context.NONE);
    }
}

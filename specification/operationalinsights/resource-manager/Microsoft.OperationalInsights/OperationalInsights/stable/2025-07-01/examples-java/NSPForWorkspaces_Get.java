
/**
 * Samples for Workspaces GetNsp.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/NSPForWorkspaces_Get.json
     */
    /**
     * Sample code: Get NSP config by name for Scheduled Query Rule.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void
        getNSPConfigByNameForScheduledQueryRule(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.workspaces().getNspWithResponse("exampleRG", "someWorkspace", "somePerimeterConfiguration",
            com.azure.core.util.Context.NONE);
    }
}

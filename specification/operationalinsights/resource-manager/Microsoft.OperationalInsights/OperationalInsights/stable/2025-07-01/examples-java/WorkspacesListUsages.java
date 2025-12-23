
/**
 * Samples for Usages List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/WorkspacesListUsages.json
     */
    /**
     * Sample code: UsagesList.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void usagesList(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.usages().list("rg1", "TestLinkWS", com.azure.core.util.Context.NONE);
    }
}

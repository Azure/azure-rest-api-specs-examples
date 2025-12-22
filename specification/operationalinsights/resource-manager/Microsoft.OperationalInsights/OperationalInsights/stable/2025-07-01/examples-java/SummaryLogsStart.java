
/**
 * Samples for SummaryLogsOperation Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/SummaryLogsStart.json
     */
    /**
     * Sample code: SummaryLogsStart.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void summaryLogsStart(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.summaryLogsOperations().start("exampleresourcegroup", "exampleworkspace", "summarylogs3",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for SummaryLogsOperation Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/SummaryLogsDelete.json
     */
    /**
     * Sample code: SummaryLogsDelete.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void summaryLogsDelete(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.summaryLogsOperations().delete("oiautorest6685", "oiautorest6685", "summarylogs1",
            com.azure.core.util.Context.NONE);
    }
}

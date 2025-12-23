
/**
 * Samples for SummaryLogsOperation Stop.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/SummaryLogsStop.json
     */
    /**
     * Sample code: SummaryLogsStop.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void summaryLogsStop(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.summaryLogsOperations().stopWithResponse("oiautorest6685", "oiautorest6685", "summarylogs1",
            com.azure.core.util.Context.NONE);
    }
}

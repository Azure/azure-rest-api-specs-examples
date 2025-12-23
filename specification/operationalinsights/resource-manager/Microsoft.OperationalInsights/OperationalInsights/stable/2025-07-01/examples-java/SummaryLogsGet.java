
/**
 * Samples for SummaryLogsOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/SummaryLogsGet.json
     */
    /**
     * Sample code: SummaryLogsGet.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void summaryLogsGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.summaryLogsOperations().getWithResponse("oiautorest6685", "oiautorest6685", "summarylogs1",
            com.azure.core.util.Context.NONE);
    }
}

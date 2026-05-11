
/**
 * Samples for SummaryLogsOperation ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/SummaryLogsList.json
     */
    /**
     * Sample code: SummaryLogsListByWorkspace.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void summaryLogsListByWorkspace(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.summaryLogsOperations().listByWorkspace("oiautorest6685", "oiautorest6685",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for SummaryLogsOperation Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/SummaryLogsStart.json
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

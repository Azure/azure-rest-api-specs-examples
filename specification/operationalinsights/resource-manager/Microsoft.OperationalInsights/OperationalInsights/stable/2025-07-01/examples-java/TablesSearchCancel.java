
/**
 * Samples for Tables CancelSearch.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/TablesSearchCancel.json
     */
    /**
     * Sample code: TablesSearchCancel.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void tablesSearchCancel(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.tables().cancelSearchWithResponse("oiautorest6685", "oiautorest6685", "table1_SRCH",
            com.azure.core.util.Context.NONE);
    }
}

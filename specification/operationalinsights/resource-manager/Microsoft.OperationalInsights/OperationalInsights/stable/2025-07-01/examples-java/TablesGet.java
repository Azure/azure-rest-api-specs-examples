
/**
 * Samples for Tables Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/TablesGet.json
     */
    /**
     * Sample code: TablesGet.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void tablesGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.tables().getWithResponse("oiautorest6685", "oiautorest6685", "table1_SRCH",
            com.azure.core.util.Context.NONE);
    }
}

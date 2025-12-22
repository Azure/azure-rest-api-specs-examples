
/**
 * Samples for Tables Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/TablesDelete.json
     */
    /**
     * Sample code: TablesDelete.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void tablesDelete(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.tables().delete("oiautorest6685", "oiautorest6685", "table1_CL", com.azure.core.util.Context.NONE);
    }
}

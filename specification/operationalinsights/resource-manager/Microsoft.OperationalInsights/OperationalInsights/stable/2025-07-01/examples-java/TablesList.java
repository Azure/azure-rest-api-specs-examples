
/**
 * Samples for Tables ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/TablesList.json
     */
    /**
     * Sample code: TablesListByWorkspace.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void tablesListByWorkspace(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.tables().listByWorkspace("oiautorest6685", "oiautorest6685", com.azure.core.util.Context.NONE);
    }
}

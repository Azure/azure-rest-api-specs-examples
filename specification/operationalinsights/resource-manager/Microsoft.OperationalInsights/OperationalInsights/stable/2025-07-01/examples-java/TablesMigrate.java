
/**
 * Samples for Tables Migrate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/TablesMigrate.json
     */
    /**
     * Sample code: TablesMigrate.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void tablesMigrate(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.tables().migrateWithResponse("oiautorest6685", "oiautorest6685", "table1_CL",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for DataExports Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/DataExportDelete.json
     */
    /**
     * Sample code: DataExportDelete.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void dataExportDelete(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.dataExports().deleteWithResponse("RgTest1", "DeWnTest1234", "export1",
            com.azure.core.util.Context.NONE);
    }
}

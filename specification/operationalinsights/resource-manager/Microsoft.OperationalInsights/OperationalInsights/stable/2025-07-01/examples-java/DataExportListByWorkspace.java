
/**
 * Samples for DataExports ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/DataExportListByWorkspace.json
     */
    /**
     * Sample code: DataExportGet.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void dataExportGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.dataExports().listByWorkspace("RgTest1", "DeWnTest1234", com.azure.core.util.Context.NONE);
    }
}

import com.azure.core.util.Context;

/** Samples for DataExports ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/DataExportListByWorkspace.json
     */
    /**
     * Sample code: DataExportGet.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void dataExportGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.dataExports().listByWorkspace("RgTest1", "DeWnTest1234", Context.NONE);
    }
}

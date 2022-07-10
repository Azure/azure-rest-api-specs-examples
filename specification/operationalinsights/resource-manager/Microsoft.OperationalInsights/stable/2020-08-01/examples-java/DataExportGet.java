import com.azure.core.util.Context;

/** Samples for DataExports Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/DataExportGet.json
     */
    /**
     * Sample code: DataExportGet.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void dataExportGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.dataExports().getWithResponse("RgTest1", "DeWnTest1234", "export1", Context.NONE);
    }
}

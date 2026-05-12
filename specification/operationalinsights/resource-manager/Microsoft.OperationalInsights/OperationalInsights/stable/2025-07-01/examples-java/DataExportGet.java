
/**
 * Samples for DataExports Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/DataExportGet.json
     */
    /**
     * Sample code: DataExportGet.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void dataExportGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.dataExports().getWithResponse("RgTest1", "DeWnTest1234", "export1", com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for StorageInsightConfigs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/
     * StorageInsightsGet.json
     */
    /**
     * Sample code: StorageInsightsGet.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void storageInsightsGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.storageInsightConfigs().getWithResponse("OIAutoRest5123", "aztest5048", "AzTestSI1110",
            com.azure.core.util.Context.NONE);
    }
}

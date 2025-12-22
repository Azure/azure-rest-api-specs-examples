
/**
 * Samples for StorageInsightConfigs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/StorageInsightsDelete.json
     */
    /**
     * Sample code: StorageInsightsDelete.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void storageInsightsDelete(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.storageInsightConfigs().deleteWithResponse("OIAutoRest5123", "aztest5048", "AzTestSI1110",
            com.azure.core.util.Context.NONE);
    }
}

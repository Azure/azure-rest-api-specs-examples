
/**
 * Samples for StorageInsightConfigs ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/StorageInsightsListByWorkspace.json
     */
    /**
     * Sample code: StorageInsightsList.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void storageInsightsList(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.storageInsightConfigs().listByWorkspace("OIAutoRest5123", "aztest5048",
            com.azure.core.util.Context.NONE);
    }
}

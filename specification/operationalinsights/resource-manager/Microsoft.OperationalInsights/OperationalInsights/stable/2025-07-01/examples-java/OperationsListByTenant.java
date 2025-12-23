
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/OperationsListByTenant.json
     */
    /**
     * Sample code: Get specific operation status.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void getSpecificOperationStatus(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}

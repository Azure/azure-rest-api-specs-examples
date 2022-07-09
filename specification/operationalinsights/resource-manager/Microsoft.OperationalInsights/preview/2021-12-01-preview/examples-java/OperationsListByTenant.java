import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/OperationsListByTenant.json
     */
    /**
     * Sample code: Get specific operation status.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void getSpecificOperationStatus(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.operations().list(Context.NONE);
    }
}

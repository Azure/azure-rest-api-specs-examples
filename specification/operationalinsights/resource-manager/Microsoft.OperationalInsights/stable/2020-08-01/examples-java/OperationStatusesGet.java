import com.azure.core.util.Context;

/** Samples for OperationStatuses Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/OperationStatusesGet.json
     */
    /**
     * Sample code: Get specific operation status.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void getSpecificOperationStatus(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.operationStatuses().getWithResponse("West US", "713192d7-503f-477a-9cfe-4efc3ee2bd11", Context.NONE);
    }
}

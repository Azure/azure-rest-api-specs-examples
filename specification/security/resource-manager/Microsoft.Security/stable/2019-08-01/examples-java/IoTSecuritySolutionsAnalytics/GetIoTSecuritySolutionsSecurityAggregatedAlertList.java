
/**
 * Samples for IotSecuritySolutionsAnalyticsAggregatedAlert List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/
     * IoTSecuritySolutionsAnalytics/GetIoTSecuritySolutionsSecurityAggregatedAlertList.json
     */
    /**
     * Sample code: Get the aggregated alert list of yours IoT Security solution.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getTheAggregatedAlertListOfYoursIoTSecuritySolution(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.iotSecuritySolutionsAnalyticsAggregatedAlerts().list("MyGroup", "default", null,
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for IotSecuritySolutionsAnalyticsAggregatedAlert Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/
     * IoTSecuritySolutionsAnalytics/GetIoTSecuritySolutionsSecurityAggregatedAlert.json
     */
    /**
     * Sample code: Get the aggregated security analytics alert of yours IoT Security solution. This aggregation is
     * performed by alert name.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        getTheAggregatedSecurityAnalyticsAlertOfYoursIoTSecuritySolutionThisAggregationIsPerformedByAlertName(
            com.azure.resourcemanager.security.SecurityManager manager) {
        manager.iotSecuritySolutionsAnalyticsAggregatedAlerts().getWithResponse("MyGroup", "default",
            "IoT_Bruteforce_Fail/2019-02-02", com.azure.core.util.Context.NONE);
    }
}

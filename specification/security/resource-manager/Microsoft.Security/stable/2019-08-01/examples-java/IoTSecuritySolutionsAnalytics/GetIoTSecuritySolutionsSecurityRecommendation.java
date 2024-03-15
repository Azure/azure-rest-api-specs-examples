
/**
 * Samples for IotSecuritySolutionsAnalyticsRecommendation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/
     * IoTSecuritySolutionsAnalytics/GetIoTSecuritySolutionsSecurityRecommendation.json
     */
    /**
     * Sample code: Get the aggregated security analytics recommendation of yours IoT Security solution.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getTheAggregatedSecurityAnalyticsRecommendationOfYoursIoTSecuritySolution(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.iotSecuritySolutionsAnalyticsRecommendations().getWithResponse("IoTEdgeResources", "default",
            "OpenPortsOnDevice", com.azure.core.util.Context.NONE);
    }
}

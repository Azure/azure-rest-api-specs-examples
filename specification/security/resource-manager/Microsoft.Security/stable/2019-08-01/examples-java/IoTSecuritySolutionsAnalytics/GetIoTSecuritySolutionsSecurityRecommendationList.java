
/**
 * Samples for IotSecuritySolutionsAnalyticsRecommendation List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/
     * IoTSecuritySolutionsAnalytics/GetIoTSecuritySolutionsSecurityRecommendationList.json
     */
    /**
     * Sample code: Get the list of aggregated security analytics recommendations of yours IoT Security solution.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getTheListOfAggregatedSecurityAnalyticsRecommendationsOfYoursIoTSecuritySolution(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.iotSecuritySolutionsAnalyticsRecommendations().list("IoTEdgeResources", "default", null,
            com.azure.core.util.Context.NONE);
    }
}

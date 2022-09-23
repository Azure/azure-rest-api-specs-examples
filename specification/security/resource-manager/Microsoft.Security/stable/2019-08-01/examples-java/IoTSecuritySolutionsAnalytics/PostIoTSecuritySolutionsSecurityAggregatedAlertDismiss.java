import com.azure.core.util.Context;

/** Samples for IotSecuritySolutionsAnalyticsAggregatedAlert Dismiss. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutionsAnalytics/PostIoTSecuritySolutionsSecurityAggregatedAlertDismiss.json
     */
    /**
     * Sample code: Dismiss an aggregated IoT Security Solution Alert.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void dismissAnAggregatedIoTSecuritySolutionAlert(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .iotSecuritySolutionsAnalyticsAggregatedAlerts()
            .dismissWithResponse("IoTEdgeResources", "default", "IoT_Bruteforce_Fail/2019-02-02/dismiss", Context.NONE);
    }
}

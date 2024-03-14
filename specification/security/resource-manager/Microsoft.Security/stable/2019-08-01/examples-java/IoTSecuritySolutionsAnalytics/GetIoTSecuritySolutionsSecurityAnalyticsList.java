
/**
 * Samples for IotSecuritySolutionAnalytics List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/
     * IoTSecuritySolutionsAnalytics/GetIoTSecuritySolutionsSecurityAnalyticsList.json
     */
    /**
     * Sample code: Get Security Solution Analytics.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecuritySolutionAnalytics(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.iotSecuritySolutionAnalytics().listWithResponse("MyGroup", "default", com.azure.core.util.Context.NONE);
    }
}

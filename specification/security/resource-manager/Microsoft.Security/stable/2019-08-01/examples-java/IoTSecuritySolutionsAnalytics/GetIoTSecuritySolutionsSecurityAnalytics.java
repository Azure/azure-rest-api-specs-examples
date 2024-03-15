
/**
 * Samples for IotSecuritySolutionAnalytics Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/
     * IoTSecuritySolutionsAnalytics/GetIoTSecuritySolutionsSecurityAnalytics.json
     */
    /**
     * Sample code: Get Security Solution Analytics.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecuritySolutionAnalytics(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.iotSecuritySolutionAnalytics().getWithResponse("MyGroup", "default", com.azure.core.util.Context.NONE);
    }
}

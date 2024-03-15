
/**
 * Samples for Tasks ListByHomeRegion.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2015-06-01-preview/examples/Tasks/
     * GetTasksSubscriptionLocation_example.json
     */
    /**
     * Sample code: Get security recommendations tasks from security data location.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecurityRecommendationsTasksFromSecurityDataLocation(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.tasks().listByHomeRegion("westeurope", null, com.azure.core.util.Context.NONE);
    }
}

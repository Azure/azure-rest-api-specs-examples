
/**
 * Samples for Tasks GetSubscriptionLevelTask.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2015-06-01-preview/examples/Tasks/
     * GetTaskSubscriptionLocation_example.json
     */
    /**
     * Sample code: Get security recommendation task from security data location.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecurityRecommendationTaskFromSecurityDataLocation(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.tasks().getSubscriptionLevelTaskWithResponse("westeurope", "62609ee7-d0a5-8616-9fe4-1df5cca7758d",
            com.azure.core.util.Context.NONE);
    }
}

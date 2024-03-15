
/**
 * Samples for Tasks List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2015-06-01-preview/examples/Tasks/
     * GetTasksSubscription_example.json
     */
    /**
     * Sample code: Get security recommendations tasks.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecurityRecommendationsTasks(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.tasks().list(null, com.azure.core.util.Context.NONE);
    }
}

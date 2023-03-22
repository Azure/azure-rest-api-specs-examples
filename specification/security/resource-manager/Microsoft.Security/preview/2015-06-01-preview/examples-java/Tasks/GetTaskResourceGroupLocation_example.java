/** Samples for Tasks GetResourceGroupLevelTask. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2015-06-01-preview/examples/Tasks/GetTaskResourceGroupLocation_example.json
     */
    /**
     * Sample code: Get security recommendation task in a resource group.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecurityRecommendationTaskInAResourceGroup(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .tasks()
            .getResourceGroupLevelTaskWithResponse(
                "myRg", "westeurope", "d55b4dc0-779c-c66c-33e5-d7bce24c4222", com.azure.core.util.Context.NONE);
    }
}

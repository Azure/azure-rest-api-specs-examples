
/**
 * Samples for WorkflowRunActionScopeRepetitions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * WorkflowRunActionScopeRepetitions_Get.json
     */
    /**
     * Sample code: Get a scoped repetition.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAScopedRepetition(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWorkflowRunActionScopeRepetitions().getWithResponse(
            "testResourceGroup", "test-name", "testFlow", "08586776228332053161046300351", "for_each", "000000",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for WorkflowRunActionRepetitions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * WorkflowRunActionRepetitions_Get.json
     */
    /**
     * Sample code: Get a repetition.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getARepetition(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWorkflowRunActionRepetitions().getWithResponse("testResourceGroup",
            "test-name", "testFlow", "08586776228332053161046300351", "testAction", "000001",
            com.azure.core.util.Context.NONE);
    }
}

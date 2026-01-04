
/**
 * Samples for WorkflowRunActionRepetitionsRequestHistories Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * WorkflowRunActionRepetitionsRequestHistories_Get.json
     */
    /**
     * Sample code: Get a repetition request history.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getARepetitionRequestHistory(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWorkflowRunActionRepetitionsRequestHistories().getWithResponse(
            "test-resource-group", "test-name", "test-workflow", "08586776228332053161046300351", "HTTP_Webhook",
            "000001", "08586611142732800686", com.azure.core.util.Context.NONE);
    }
}

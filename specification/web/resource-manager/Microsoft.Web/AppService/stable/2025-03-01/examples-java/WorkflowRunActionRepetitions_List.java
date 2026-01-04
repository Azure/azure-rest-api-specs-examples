
/**
 * Samples for WorkflowRunActionRepetitions List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * WorkflowRunActionRepetitions_List.json
     */
    /**
     * Sample code: List repetitions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRepetitions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWorkflowRunActionRepetitions().list("testResourceGroup",
            "test-name", "testFlow", "08586776228332053161046300351", "testAction", com.azure.core.util.Context.NONE);
    }
}

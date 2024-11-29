
/**
 * Samples for WorkflowRunActionScopeRepetitions List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/
     * WorkflowRunActionScopeRepetitions_List.json
     */
    /**
     * Sample code: List the scoped repetitions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listTheScopedRepetitions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWorkflowRunActionScopeRepetitions().list("testResourceGroup",
            "test-name", "testFlow", "08586776228332053161046300351", "for_each", com.azure.core.util.Context.NONE);
    }
}

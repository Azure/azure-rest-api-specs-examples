
/**
 * Samples for WorkflowRuns List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/WorkflowRuns_List.json
     */
    /**
     * Sample code: List workflow runs.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listWorkflowRuns(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWorkflowRuns().list("test-resource-group", "test-name",
            "test-workflow", null, null, com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for WorkflowRuns Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/WorkflowRuns_Cancel.json
     */
    /**
     * Sample code: Cancel a workflow run.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cancelAWorkflowRun(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWorkflowRuns().cancelWithResponse("test-resource-group",
            "test-name", "test-workflow", "08586676746934337772206998657CU22", com.azure.core.util.Context.NONE);
    }
}

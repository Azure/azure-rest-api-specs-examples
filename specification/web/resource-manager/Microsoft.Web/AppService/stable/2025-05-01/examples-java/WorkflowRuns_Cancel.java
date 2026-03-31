
/**
 * Samples for WorkflowRuns Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowRuns_Cancel.json
     */
    /**
     * Sample code: Cancel a workflow run.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void cancelAWorkflowRun(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowRuns().cancelWithResponse("test-resource-group", "test-name",
            "test-workflow", "08586676746934337772206998657CU22", com.azure.core.util.Context.NONE);
    }
}

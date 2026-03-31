
/**
 * Samples for WorkflowRuns List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowRuns_List.json
     */
    /**
     * Sample code: List workflow runs.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listWorkflowRuns(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowRuns().list("test-resource-group", "test-name", "test-workflow", null, null,
            com.azure.core.util.Context.NONE);
    }
}

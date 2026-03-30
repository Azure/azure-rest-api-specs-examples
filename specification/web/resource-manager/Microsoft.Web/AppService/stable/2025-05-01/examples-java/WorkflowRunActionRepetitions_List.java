
/**
 * Samples for WorkflowRunActionRepetitions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowRunActionRepetitions_List.json
     */
    /**
     * Sample code: List repetitions.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listRepetitions(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowRunActionRepetitions().list("testResourceGroup", "test-name", "testFlow",
            "08586776228332053161046300351", "testAction", com.azure.core.util.Context.NONE);
    }
}

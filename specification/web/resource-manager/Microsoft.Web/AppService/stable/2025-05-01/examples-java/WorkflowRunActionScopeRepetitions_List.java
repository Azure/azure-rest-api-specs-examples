
/**
 * Samples for WorkflowRunActionScopeRepetitions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowRunActionScopeRepetitions_List.json
     */
    /**
     * Sample code: List the scoped repetitions.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listTheScopedRepetitions(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowRunActionScopeRepetitions().list("testResourceGroup", "test-name",
            "testFlow", "08586776228332053161046300351", "for_each", com.azure.core.util.Context.NONE);
    }
}

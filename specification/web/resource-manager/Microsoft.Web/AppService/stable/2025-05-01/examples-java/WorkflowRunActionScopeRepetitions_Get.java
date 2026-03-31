
/**
 * Samples for WorkflowRunActionScopeRepetitions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowRunActionScopeRepetitions_Get.json
     */
    /**
     * Sample code: Get a scoped repetition.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAScopedRepetition(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowRunActionScopeRepetitions().getWithResponse("testResourceGroup", "test-name",
            "testFlow", "08586776228332053161046300351", "for_each", "000000", com.azure.core.util.Context.NONE);
    }
}

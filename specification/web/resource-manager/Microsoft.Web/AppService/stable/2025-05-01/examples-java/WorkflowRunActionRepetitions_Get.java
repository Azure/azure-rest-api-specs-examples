
/**
 * Samples for WorkflowRunActionRepetitions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowRunActionRepetitions_Get.json
     */
    /**
     * Sample code: Get a repetition.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getARepetition(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowRunActionRepetitions().getWithResponse("testResourceGroup", "test-name",
            "testFlow", "08586776228332053161046300351", "testAction", "000001", com.azure.core.util.Context.NONE);
    }
}

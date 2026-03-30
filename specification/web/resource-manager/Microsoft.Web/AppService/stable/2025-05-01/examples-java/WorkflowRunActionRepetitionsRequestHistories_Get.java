
/**
 * Samples for WorkflowRunActionRepetitionsRequestHistories Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowRunActionRepetitionsRequestHistories_Get.json
     */
    /**
     * Sample code: Get a repetition request history.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getARepetitionRequestHistory(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowRunActionRepetitionsRequestHistories().getWithResponse("test-resource-group",
            "test-name", "test-workflow", "08586776228332053161046300351", "HTTP_Webhook", "000001",
            "08586611142732800686", com.azure.core.util.Context.NONE);
    }
}

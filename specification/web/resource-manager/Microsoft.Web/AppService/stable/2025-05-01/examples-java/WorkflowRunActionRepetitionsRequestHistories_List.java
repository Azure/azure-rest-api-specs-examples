
/**
 * Samples for WorkflowRunActionRepetitionsRequestHistories List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowRunActionRepetitionsRequestHistories_List.json
     */
    /**
     * Sample code: List repetition request history.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listRepetitionRequestHistory(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowRunActionRepetitionsRequestHistories().list("test-resource-group",
            "test-name", "test-workflow", "08586776228332053161046300351", "HTTP_Webhook", "000001",
            com.azure.core.util.Context.NONE);
    }
}

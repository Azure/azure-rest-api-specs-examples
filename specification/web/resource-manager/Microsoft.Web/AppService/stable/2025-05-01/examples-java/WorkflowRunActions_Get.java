
/**
 * Samples for WorkflowRunActions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowRunActions_Get.json
     */
    /**
     * Sample code: Get a workflow run action.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAWorkflowRunAction(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowRunActions().getWithResponse("test-resource-group", "test-name",
            "test-workflow", "08586676746934337772206998657CU22", "HTTP", com.azure.core.util.Context.NONE);
    }
}

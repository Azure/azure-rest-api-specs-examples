
/**
 * Samples for WorkflowVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowVersions_Get.json
     */
    /**
     * Sample code: Get a workflow version.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAWorkflowVersion(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowVersions().getWithResponse("test-resource-group", "test-name",
            "test-workflow", "08586676824806722526", com.azure.core.util.Context.NONE);
    }
}

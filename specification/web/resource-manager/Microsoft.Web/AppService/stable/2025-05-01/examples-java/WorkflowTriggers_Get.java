
/**
 * Samples for WorkflowTriggers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowTriggers_Get.json
     */
    /**
     * Sample code: Get a workflow trigger.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAWorkflowTrigger(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowTriggers().getWithResponse("test-resource-group", "test-name",
            "test-workflow", "manual", com.azure.core.util.Context.NONE);
    }
}

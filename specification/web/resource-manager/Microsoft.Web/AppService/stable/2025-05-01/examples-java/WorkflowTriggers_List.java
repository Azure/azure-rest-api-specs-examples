
/**
 * Samples for WorkflowTriggers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowTriggers_List.json
     */
    /**
     * Sample code: List workflow triggers.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listWorkflowTriggers(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowTriggers().list("test-resource-group", "test-name", "test-workflow", null,
            null, com.azure.core.util.Context.NONE);
    }
}

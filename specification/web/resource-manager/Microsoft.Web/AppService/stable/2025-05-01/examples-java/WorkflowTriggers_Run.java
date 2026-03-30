
/**
 * Samples for WorkflowTriggers Run.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowTriggers_Run.json
     */
    /**
     * Sample code: Run a workflow trigger.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void runAWorkflowTrigger(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowTriggers().run("test-resource-group", "test-name", "test-workflow",
            "recurrence", com.azure.core.util.Context.NONE);
    }
}

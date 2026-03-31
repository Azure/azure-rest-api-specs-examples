
/**
 * Samples for WorkflowTriggers ListCallbackUrl.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowTriggers_ListCallbackUrl.json
     */
    /**
     * Sample code: Get the callback URL for a trigger.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getTheCallbackURLForATrigger(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowTriggers().listCallbackUrlWithResponse("test-resource-group", "test-name",
            "test-workflow", "manual", com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for WorkflowTriggerHistories List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowTriggerHistories_List.json
     */
    /**
     * Sample code: List a workflow trigger history.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listAWorkflowTriggerHistory(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowTriggerHistories().list("testResourceGroup", "test-name", "testWorkflowName",
            "testTriggerName", null, null, com.azure.core.util.Context.NONE);
    }
}

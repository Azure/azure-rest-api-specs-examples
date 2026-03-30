
/**
 * Samples for WorkflowTriggerHistories Resubmit.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowTriggerHistories_Resubmit.json
     */
    /**
     * Sample code: Resubmit a workflow run based on the trigger history.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        resubmitAWorkflowRunBasedOnTheTriggerHistory(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowTriggerHistories().resubmit("testResourceGroup", "test-name",
            "testWorkflowName", "testTriggerName", "testHistoryName", com.azure.core.util.Context.NONE);
    }
}

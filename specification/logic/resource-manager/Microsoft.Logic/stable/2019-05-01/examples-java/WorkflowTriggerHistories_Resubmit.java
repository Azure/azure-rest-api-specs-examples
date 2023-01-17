/** Samples for WorkflowTriggerHistories Resubmit. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowTriggerHistories_Resubmit.json
     */
    /**
     * Sample code: Resubmit a workflow run based on the trigger history.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void resubmitAWorkflowRunBasedOnTheTriggerHistory(
        com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowTriggerHistories()
            .resubmitWithResponse(
                "testResourceGroup",
                "testWorkflowName",
                "testTriggerName",
                "testHistoryName",
                com.azure.core.util.Context.NONE);
    }
}

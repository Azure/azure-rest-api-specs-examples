
/**
 * Samples for WorkflowRunActionRepetitions ListExpressionTraces.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/
     * WorkflowRunActionRepetitions_ListExpressionTraces.json
     */
    /**
     * Sample code: List expression traces for a repetition.
     * 
     * @param manager Entry point to LogicManager.
     */
    public static void listExpressionTracesForARepetition(com.azure.resourcemanager.logic.LogicManager manager) {
        manager.workflowRunActionRepetitions().listExpressionTraces("testResourceGroup", "testFlow",
            "08586776228332053161046300351", "testAction", "000001", com.azure.core.util.Context.NONE);
    }
}

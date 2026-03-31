
/**
 * Samples for WorkflowRunActionRepetitions ListExpressionTraces.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowRunActionRepetitions_ListExpressionTraces.json
     */
    /**
     * Sample code: List expression traces for a repetition.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        listExpressionTracesForARepetition(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowRunActionRepetitions().listExpressionTraces("testResourceGroup", "test-name",
            "testFlow", "08586776228332053161046300351", "testAction", "000001", com.azure.core.util.Context.NONE);
    }
}

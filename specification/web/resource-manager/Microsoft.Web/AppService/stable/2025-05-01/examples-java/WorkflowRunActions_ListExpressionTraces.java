
/**
 * Samples for WorkflowRunActions ListExpressionTraces.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowRunActions_ListExpressionTraces.json
     */
    /**
     * Sample code: List expression traces.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listExpressionTraces(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowRunActions().listExpressionTraces("testResourceGroup", "test-name",
            "testFlow", "08586776228332053161046300351", "testAction", com.azure.core.util.Context.NONE);
    }
}

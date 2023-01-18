/** Samples for WorkflowRunActions ListExpressionTraces. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRunActions_ListExpressionTraces.json
     */
    /**
     * Sample code: List expression traces.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void listExpressionTraces(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowRunActions()
            .listExpressionTraces(
                "testResourceGroup",
                "testFlow",
                "08586776228332053161046300351",
                "testAction",
                com.azure.core.util.Context.NONE);
    }
}

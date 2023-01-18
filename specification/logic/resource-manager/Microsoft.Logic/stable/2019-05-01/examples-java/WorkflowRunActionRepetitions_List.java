/** Samples for WorkflowRunActionRepetitions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRunActionRepetitions_List.json
     */
    /**
     * Sample code: List repetitions.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void listRepetitions(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowRunActionRepetitions()
            .list(
                "testResourceGroup",
                "testFlow",
                "08586776228332053161046300351",
                "testAction",
                com.azure.core.util.Context.NONE);
    }
}

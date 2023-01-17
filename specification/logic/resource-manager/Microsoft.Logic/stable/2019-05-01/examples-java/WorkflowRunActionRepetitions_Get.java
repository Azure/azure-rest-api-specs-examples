/** Samples for WorkflowRunActionRepetitions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRunActionRepetitions_Get.json
     */
    /**
     * Sample code: Get a repetition.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getARepetition(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowRunActionRepetitions()
            .getWithResponse(
                "testResourceGroup",
                "testFlow",
                "08586776228332053161046300351",
                "testAction",
                "000001",
                com.azure.core.util.Context.NONE);
    }
}

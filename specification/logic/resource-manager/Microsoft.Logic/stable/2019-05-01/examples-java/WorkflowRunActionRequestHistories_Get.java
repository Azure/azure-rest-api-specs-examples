/** Samples for WorkflowRunActionRequestHistories Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRunActionRequestHistories_Get.json
     */
    /**
     * Sample code: Get a request history.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getARequestHistory(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowRunActionRequestHistories()
            .getWithResponse(
                "test-resource-group",
                "test-workflow",
                "08586776228332053161046300351",
                "HTTP_Webhook",
                "08586611142732800686",
                com.azure.core.util.Context.NONE);
    }
}

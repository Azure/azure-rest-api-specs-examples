
/**
 * Samples for WorkflowRunOperations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRunOperations_Get.json
     */
    /**
     * Sample code: Get a run operation.
     * 
     * @param manager Entry point to LogicManager.
     */
    public static void getARunOperation(com.azure.resourcemanager.logic.LogicManager manager) {
        manager.workflowRunOperations().getWithResponse("testResourceGroup", "testFlow",
            "08586774142730039209110422528", "ebdcbbde-c4db-43ec-987c-fd0f7726f43b", com.azure.core.util.Context.NONE);
    }
}

/** Samples for Workflows GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_Get.json
     */
    /**
     * Sample code: Get a workflow.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getAWorkflow(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflows()
            .getByResourceGroupWithResponse("test-resource-group", "test-workflow", com.azure.core.util.Context.NONE);
    }
}

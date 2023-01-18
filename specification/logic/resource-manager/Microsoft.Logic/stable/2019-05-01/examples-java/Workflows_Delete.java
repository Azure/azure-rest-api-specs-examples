/** Samples for Workflows Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_Delete.json
     */
    /**
     * Sample code: Delete a workflow.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void deleteAWorkflow(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflows()
            .deleteByResourceGroupWithResponse(
                "test-resource-group", "test-workflow", com.azure.core.util.Context.NONE);
    }
}

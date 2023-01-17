/** Samples for Workflows Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_Update.json
     */
    /**
     * Sample code: Patch a workflow.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void patchAWorkflow(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflows()
            .updateWithResponse("test-resource-group", "test-workflow", com.azure.core.util.Context.NONE);
    }
}

/** Samples for Workflows Disable. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_Disable.json
     */
    /**
     * Sample code: Disable a workflow.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void disableAWorkflow(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflows()
            .disableWithResponse("test-resource-group", "test-workflow", com.azure.core.util.Context.NONE);
    }
}

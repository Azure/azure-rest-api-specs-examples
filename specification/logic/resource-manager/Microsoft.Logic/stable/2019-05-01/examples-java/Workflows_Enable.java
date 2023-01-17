/** Samples for Workflows Enable. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_Enable.json
     */
    /**
     * Sample code: Enable a workflow.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void enableAWorkflow(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflows()
            .enableWithResponse("test-resource-group", "test-workflow", com.azure.core.util.Context.NONE);
    }
}

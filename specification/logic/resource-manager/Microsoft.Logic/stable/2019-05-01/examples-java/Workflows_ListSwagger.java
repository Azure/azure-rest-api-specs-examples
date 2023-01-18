/** Samples for Workflows ListSwagger. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_ListSwagger.json
     */
    /**
     * Sample code: Get the swagger for a workflow.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getTheSwaggerForAWorkflow(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflows()
            .listSwaggerWithResponse("testResourceGroup", "testWorkflowName", com.azure.core.util.Context.NONE);
    }
}

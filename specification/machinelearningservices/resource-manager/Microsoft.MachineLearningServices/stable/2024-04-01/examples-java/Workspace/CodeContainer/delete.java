
/**
 * Samples for CodeContainers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/CodeContainer/delete.json
     */
    /**
     * Sample code: Delete Workspace Code Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteWorkspaceCodeContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.codeContainers().deleteWithResponse("testrg123", "testworkspace", "testContainer",
            com.azure.core.util.Context.NONE);
    }
}

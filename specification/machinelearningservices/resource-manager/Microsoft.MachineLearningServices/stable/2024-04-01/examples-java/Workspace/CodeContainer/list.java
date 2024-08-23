
/**
 * Samples for CodeContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/CodeContainer/list.json
     */
    /**
     * Sample code: List Workspace Code Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listWorkspaceCodeContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.codeContainers().list("testrg123", "testworkspace", null, com.azure.core.util.Context.NONE);
    }
}

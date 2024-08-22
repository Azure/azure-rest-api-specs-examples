
/**
 * Samples for EnvironmentContainers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/EnvironmentContainer/delete.json
     */
    /**
     * Sample code: Delete Workspace Environment Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteWorkspaceEnvironmentContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.environmentContainers().deleteWithResponse("testrg123", "testworkspace", "testContainer",
            com.azure.core.util.Context.NONE);
    }
}

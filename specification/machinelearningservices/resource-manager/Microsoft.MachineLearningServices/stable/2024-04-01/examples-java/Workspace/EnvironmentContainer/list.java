
/**
 * Samples for EnvironmentContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/EnvironmentContainer/list.json
     */
    /**
     * Sample code: List Workspace Environment Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listWorkspaceEnvironmentContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.environmentContainers().list("testrg123", "testworkspace", null, null,
            com.azure.core.util.Context.NONE);
    }
}

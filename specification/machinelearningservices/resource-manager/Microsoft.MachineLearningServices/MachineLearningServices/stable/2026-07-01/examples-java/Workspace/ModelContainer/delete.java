
/**
 * Samples for ModelContainers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/ModelContainer/delete.json
     */
    /**
     * Sample code: Delete Workspace Model Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteWorkspaceModelContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.modelContainers().deleteWithResponse("testrg123", "workspace123", "testContainer",
            com.azure.core.util.Context.NONE);
    }
}

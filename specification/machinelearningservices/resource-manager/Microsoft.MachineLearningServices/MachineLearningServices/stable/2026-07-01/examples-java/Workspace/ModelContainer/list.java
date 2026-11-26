
/**
 * Samples for ModelContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/ModelContainer/list.json
     */
    /**
     * Sample code: List Workspace Model Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listWorkspaceModelContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.modelContainers().list("testrg123", "workspace123", null, null, null, com.azure.core.util.Context.NONE);
    }
}

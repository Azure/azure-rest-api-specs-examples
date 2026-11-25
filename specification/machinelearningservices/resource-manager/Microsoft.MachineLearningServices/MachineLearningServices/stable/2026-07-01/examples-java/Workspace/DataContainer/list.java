
/**
 * Samples for DataContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/DataContainer/list.json
     */
    /**
     * Sample code: List Workspace Data Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listWorkspaceDataContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.dataContainers().list("testrg123", "workspace123", null, null, com.azure.core.util.Context.NONE);
    }
}

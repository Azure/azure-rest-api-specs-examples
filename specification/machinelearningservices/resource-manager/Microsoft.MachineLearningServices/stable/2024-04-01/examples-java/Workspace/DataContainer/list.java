
/**
 * Samples for DataContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/DataContainer/list.json
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

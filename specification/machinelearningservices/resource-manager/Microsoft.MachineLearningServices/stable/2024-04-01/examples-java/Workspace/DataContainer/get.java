
/**
 * Samples for DataContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/DataContainer/get.json
     */
    /**
     * Sample code: Get Workspace Data Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspaceDataContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.dataContainers().getWithResponse("testrg123", "workspace123", "datacontainer123",
            com.azure.core.util.Context.NONE);
    }
}

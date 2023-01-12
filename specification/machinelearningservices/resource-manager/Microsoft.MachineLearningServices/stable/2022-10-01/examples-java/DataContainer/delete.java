/** Samples for DataContainers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/DataContainer/delete.json
     */
    /**
     * Sample code: Delete Data Container.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteDataContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .dataContainers()
            .deleteWithResponse("testrg123", "workspace123", "datacontainer123", com.azure.core.util.Context.NONE);
    }
}

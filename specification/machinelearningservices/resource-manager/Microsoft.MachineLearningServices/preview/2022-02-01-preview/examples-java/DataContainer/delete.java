import com.azure.core.util.Context;

/** Samples for DataContainers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/DataContainer/delete.json
     */
    /**
     * Sample code: Delete Data Container.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteDataContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.dataContainers().deleteWithResponse("testrg123", "workspace123", "datacontainer123", Context.NONE);
    }
}

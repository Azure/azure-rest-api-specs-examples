import com.azure.core.util.Context;

/** Samples for DataContainers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/DataContainer/get.json
     */
    /**
     * Sample code: Get Data Container.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getDataContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.dataContainers().getWithResponse("testrg123", "workspace123", "datacontainer123", Context.NONE);
    }
}

import com.azure.core.util.Context;

/** Samples for DataContainers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/DataContainer/list.json
     */
    /**
     * Sample code: List Data Container.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listDataContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.dataContainers().list("testrg123", "workspace123", null, null, Context.NONE);
    }
}

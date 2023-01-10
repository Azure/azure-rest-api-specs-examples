import com.azure.core.util.Context;

/** Samples for DataContainers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/DataContainer/list.json
     */
    /**
     * Sample code: List Data Container.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listDataContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.dataContainers().list("testrg123", "workspace123", null, null, Context.NONE);
    }
}

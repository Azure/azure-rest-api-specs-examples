import com.azure.core.util.Context;

/** Samples for ModelContainers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/ModelContainer/list.json
     */
    /**
     * Sample code: List Model Container.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listModelContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.modelContainers().list("testrg123", "workspace123", null, null, null, Context.NONE);
    }
}

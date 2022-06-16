import com.azure.core.util.Context;

/** Samples for ModelContainers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/ModelContainer/get.json
     */
    /**
     * Sample code: Get Model Container.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getModelContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.modelContainers().getWithResponse("testrg123", "workspace123", "testContainer", Context.NONE);
    }
}

import com.azure.core.util.Context;

/** Samples for ModelContainers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/ModelContainer/get.json
     */
    /**
     * Sample code: Get Model Container.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getModelContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.modelContainers().getWithResponse("testrg123", "workspace123", "testContainer", Context.NONE);
    }
}

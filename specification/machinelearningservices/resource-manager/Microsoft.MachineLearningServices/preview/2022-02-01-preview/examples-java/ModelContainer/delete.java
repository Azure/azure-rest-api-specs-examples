import com.azure.core.util.Context;

/** Samples for ModelContainers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/ModelContainer/delete.json
     */
    /**
     * Sample code: Delete Model Container.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteModelContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.modelContainers().deleteWithResponse("testrg123", "workspace123", "testContainer", Context.NONE);
    }
}

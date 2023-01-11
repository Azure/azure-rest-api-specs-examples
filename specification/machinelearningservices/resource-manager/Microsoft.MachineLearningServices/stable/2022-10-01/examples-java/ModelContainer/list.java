/** Samples for ModelContainers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/ModelContainer/list.json
     */
    /**
     * Sample code: List Model Container.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listModelContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.modelContainers().list("testrg123", "workspace123", null, null, null, com.azure.core.util.Context.NONE);
    }
}

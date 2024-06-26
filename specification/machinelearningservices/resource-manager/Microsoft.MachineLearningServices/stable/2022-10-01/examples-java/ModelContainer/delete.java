/** Samples for ModelContainers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/ModelContainer/delete.json
     */
    /**
     * Sample code: Delete Model Container.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteModelContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .modelContainers()
            .deleteWithResponse("testrg123", "workspace123", "testContainer", com.azure.core.util.Context.NONE);
    }
}

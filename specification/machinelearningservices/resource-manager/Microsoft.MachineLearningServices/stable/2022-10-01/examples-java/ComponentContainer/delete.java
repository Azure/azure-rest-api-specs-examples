/** Samples for ComponentContainers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/ComponentContainer/delete.json
     */
    /**
     * Sample code: Delete Component Container.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteComponentContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .componentContainers()
            .deleteWithResponse("test-rg", "my-aml-workspace", "string", com.azure.core.util.Context.NONE);
    }
}

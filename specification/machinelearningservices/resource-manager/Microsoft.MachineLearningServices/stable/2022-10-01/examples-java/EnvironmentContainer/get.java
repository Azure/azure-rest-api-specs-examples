/** Samples for EnvironmentContainers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/EnvironmentContainer/get.json
     */
    /**
     * Sample code: Get Environment Container.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getEnvironmentContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .environmentContainers()
            .getWithResponse("testrg123", "testworkspace", "testEnvironment", com.azure.core.util.Context.NONE);
    }
}

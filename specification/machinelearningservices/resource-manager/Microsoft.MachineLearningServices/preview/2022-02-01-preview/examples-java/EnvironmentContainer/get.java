import com.azure.core.util.Context;

/** Samples for EnvironmentContainers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/EnvironmentContainer/get.json
     */
    /**
     * Sample code: Get Environment Container.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getEnvironmentContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.environmentContainers().getWithResponse("testrg123", "testworkspace", "testEnvironment", Context.NONE);
    }
}

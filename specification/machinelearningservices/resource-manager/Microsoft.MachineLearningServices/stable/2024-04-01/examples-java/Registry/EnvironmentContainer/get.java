
/**
 * Samples for RegistryEnvironmentContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/EnvironmentContainer/get.json
     */
    /**
     * Sample code: Get Registry Environment Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getRegistryEnvironmentContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryEnvironmentContainers().getWithResponse("testrg123", "testregistry", "testEnvironment",
            com.azure.core.util.Context.NONE);
    }
}

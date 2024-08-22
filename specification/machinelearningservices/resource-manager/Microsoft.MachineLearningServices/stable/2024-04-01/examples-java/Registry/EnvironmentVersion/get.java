
/**
 * Samples for RegistryEnvironmentVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/EnvironmentVersion/get.json
     */
    /**
     * Sample code: Get Registry Environment Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getRegistryEnvironmentVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryEnvironmentVersions().getWithResponse("test-rg", "my-aml-registry", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}

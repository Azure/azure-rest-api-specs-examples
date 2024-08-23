
/**
 * Samples for RegistryComponentVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/ComponentVersion/get.json
     */
    /**
     * Sample code: Get Registry Component Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getRegistryComponentVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryComponentVersions().getWithResponse("test-rg", "my-aml-registry", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}

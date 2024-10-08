
/**
 * Samples for RegistryModelVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/ModelVersion/get.json
     */
    /**
     * Sample code: Get Registry Model Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getRegistryModelVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryModelVersions().getWithResponse("test-rg", "my-aml-registry", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}

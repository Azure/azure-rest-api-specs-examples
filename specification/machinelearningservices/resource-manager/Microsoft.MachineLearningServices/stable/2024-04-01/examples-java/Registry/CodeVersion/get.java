
/**
 * Samples for RegistryCodeVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/CodeVersion/get.json
     */
    /**
     * Sample code: Get Registry Code Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getRegistryCodeVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryCodeVersions().getWithResponse("test-rg", "my-aml-registry", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}

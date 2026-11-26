
/**
 * Samples for RegistryCodeVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registry/CodeVersion/get.json
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


/**
 * Samples for RegistryDataVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registry/DataVersionBase/get.json
     */
    /**
     * Sample code: Get Registry Data Version Base.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getRegistryDataVersionBase(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryDataVersions().getWithResponse("test-rg", "registryName", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for RegistryDataContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registry/DataContainer/get.json
     */
    /**
     * Sample code: Get Registry Data Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getRegistryDataContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryDataContainers().getWithResponse("test-rg", "registryName", "string",
            com.azure.core.util.Context.NONE);
    }
}

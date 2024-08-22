
/**
 * Samples for RegistryDataContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/DataContainer/get.json
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

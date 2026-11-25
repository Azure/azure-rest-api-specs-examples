
/**
 * Samples for RegistryComponentContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registry/ComponentContainer/get.json
     */
    /**
     * Sample code: Get Registry Component Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getRegistryComponentContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryComponentContainers().getWithResponse("test-rg", "my-aml-registry", "string",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for RegistryComponentContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/ComponentContainer/get.json
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


/**
 * Samples for RegistryComponentContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/ComponentContainer/list.json
     */
    /**
     * Sample code: List Registry Component Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listRegistryComponentContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryComponentContainers().list("test-rg", "my-aml-registry", null,
            com.azure.core.util.Context.NONE);
    }
}

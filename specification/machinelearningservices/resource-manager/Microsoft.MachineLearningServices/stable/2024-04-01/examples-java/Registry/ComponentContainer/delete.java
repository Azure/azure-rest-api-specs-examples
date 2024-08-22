
/**
 * Samples for RegistryComponentContainers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/ComponentContainer/delete.json
     */
    /**
     * Sample code: Delete Registry Component Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteRegistryComponentContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryComponentContainers().delete("test-rg", "my-aml-registry", "string",
            com.azure.core.util.Context.NONE);
    }
}

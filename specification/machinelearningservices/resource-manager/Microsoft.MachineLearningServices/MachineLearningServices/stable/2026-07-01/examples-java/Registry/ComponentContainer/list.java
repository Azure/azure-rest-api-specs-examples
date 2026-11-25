
/**
 * Samples for RegistryComponentContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registry/ComponentContainer/list.json
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

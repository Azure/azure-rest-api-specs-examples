
/**
 * Samples for RegistryDataContainers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registry/DataContainer/delete.json
     */
    /**
     * Sample code: Delete Registry Data Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteRegistryDataContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryDataContainers().delete("test-rg", "registryName", "string", com.azure.core.util.Context.NONE);
    }
}

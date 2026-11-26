
/**
 * Samples for RegistryDataVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registry/DataVersionBase/delete.json
     */
    /**
     * Sample code: Delete Registry Data Version Base.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteRegistryDataVersionBase(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryDataVersions().delete("test-rg", "registryName", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}

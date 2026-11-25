
/**
 * Samples for RegistryCodeVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registry/CodeVersion/delete.json
     */
    /**
     * Sample code: Delete Registry Code Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteRegistryCodeVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryCodeVersions().delete("test-rg", "my-aml-registry", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}

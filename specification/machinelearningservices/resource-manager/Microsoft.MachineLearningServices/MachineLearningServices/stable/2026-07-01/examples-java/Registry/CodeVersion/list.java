
/**
 * Samples for RegistryCodeVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registry/CodeVersion/list.json
     */
    /**
     * Sample code: List Registry Code Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listRegistryCodeVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryCodeVersions().list("test-rg", "my-aml-registry", "string", "string", 1, null,
            com.azure.core.util.Context.NONE);
    }
}

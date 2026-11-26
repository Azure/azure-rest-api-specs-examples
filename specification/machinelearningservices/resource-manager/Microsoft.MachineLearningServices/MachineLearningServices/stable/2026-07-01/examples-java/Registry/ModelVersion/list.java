
/**
 * Samples for RegistryModelVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registry/ModelVersion/list.json
     */
    /**
     * Sample code: List Registry Model Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listRegistryModelVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryModelVersions().list("test-rg", "my-aml-registry", "string", null, "string", 1, "string",
            "string", "string", "string", null, com.azure.core.util.Context.NONE);
    }
}

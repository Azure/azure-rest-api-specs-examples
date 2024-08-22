
/**
 * Samples for RegistryComponentVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/ComponentVersion/list.json
     */
    /**
     * Sample code: List Registry Component Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listRegistryComponentVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryComponentVersions().list("test-rg", "my-aml-registry", "string", "string", 1, null,
            com.azure.core.util.Context.NONE);
    }
}

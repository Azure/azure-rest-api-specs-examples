
/**
 * Samples for RegistryComponentVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/ComponentVersion/delete.json
     */
    /**
     * Sample code: Delete Registry Component Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteRegistryComponentVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryComponentVersions().delete("test-rg", "my-aml-registry", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}

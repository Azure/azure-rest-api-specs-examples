
/**
 * Samples for RegistryModelVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/ModelVersion/delete.json
     */
    /**
     * Sample code: Delete Registry Model Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteRegistryModelVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryModelVersions().delete("test-rg", "my-aml-registry", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}

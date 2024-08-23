
/**
 * Samples for RegistryEnvironmentVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/EnvironmentVersion/delete.json
     */
    /**
     * Sample code: Delete Registry Environment Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteRegistryEnvironmentVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryEnvironmentVersions().delete("test-rg", "my-aml-registry", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}

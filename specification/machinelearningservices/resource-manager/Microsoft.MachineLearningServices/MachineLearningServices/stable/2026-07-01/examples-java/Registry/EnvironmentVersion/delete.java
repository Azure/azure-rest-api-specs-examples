
/**
 * Samples for RegistryEnvironmentVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registry/EnvironmentVersion/delete.json
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

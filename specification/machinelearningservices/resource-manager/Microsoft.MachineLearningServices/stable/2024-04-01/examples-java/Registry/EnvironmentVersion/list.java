
/**
 * Samples for RegistryEnvironmentVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/EnvironmentVersion/list.json
     */
    /**
     * Sample code: List Registry Environment Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listRegistryEnvironmentVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryEnvironmentVersions().list("test-rg", "my-aml-regsitry", "string", "string", 1, null, null,
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for RegistryEnvironmentContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/EnvironmentContainer/list.json
     */
    /**
     * Sample code: List Registry Environment Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listRegistryEnvironmentContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryEnvironmentContainers().list("testrg123", "testregistry", null, null,
            com.azure.core.util.Context.NONE);
    }
}

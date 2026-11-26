
/**
 * Samples for RegistryEnvironmentContainers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registry/EnvironmentContainer/delete.json
     */
    /**
     * Sample code: Delete Registry Environment Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteRegistryEnvironmentContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryEnvironmentContainers().delete("testrg123", "testregistry", "testContainer",
            com.azure.core.util.Context.NONE);
    }
}

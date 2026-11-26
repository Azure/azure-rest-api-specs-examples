
/**
 * Samples for RegistryModelContainers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registry/ModelContainer/delete.json
     */
    /**
     * Sample code: Delete Registry Model Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteRegistryModelContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryModelContainers().delete("testrg123", "registry123", "testContainer",
            com.azure.core.util.Context.NONE);
    }
}

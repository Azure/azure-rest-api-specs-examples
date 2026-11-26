
/**
 * Samples for RegistryModelContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registry/ModelContainer/list.json
     */
    /**
     * Sample code: List Registry Model Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listRegistryModelContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryModelContainers().list("testrg123", "registry123", null, null,
            com.azure.core.util.Context.NONE);
    }
}

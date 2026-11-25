
/**
 * Samples for RegistryCodeContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registry/CodeContainer/list.json
     */
    /**
     * Sample code: List Registry Code Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listRegistryCodeContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryCodeContainers().list("testrg123", "testregistry", null, com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for RegistryCodeContainers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/CodeContainer/delete.json
     */
    /**
     * Sample code: Delete Registry Code Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteRegistryCodeContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryCodeContainers().delete("testrg123", "testregistry", "testContainer",
            com.azure.core.util.Context.NONE);
    }
}
